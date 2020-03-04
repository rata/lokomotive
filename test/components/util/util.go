// Copyright 2020 The Lokomotive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func KubeconfigPath(t *testing.T) string {
	kubeconfig := os.ExpandEnv(os.Getenv("KUBECONFIG"))

	if kubeconfig == "" {
		t.Fatalf("env var KUBECONFIG was not set")
	}

	return kubeconfig
}

func buildKubeConfig(t *testing.T) (*restclient.Config, error) {
	kubeconfig := KubeconfigPath(t)

	t.Logf("using KUBECONFIG=%s", kubeconfig)

	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func CreateKubeClient(t *testing.T) (*kubernetes.Clientset, error) {
	config, err := buildKubeConfig(t)
	if err != nil {
		t.Fatalf("could not build config from KUBECONFIG: %v", err)
	}

	return kubernetes.NewForConfig(config)
}

func WaitForStatefulSet(t *testing.T, client kubernetes.Interface, ns, name string, replicas int, retryInterval, timeout time.Duration) {
	if err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		ds, err := client.AppsV1().StatefulSets(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				t.Logf("waiting for statefulset %s to be available", name)
				return false, nil
			}
			return false, err
		}

		t.Logf("statefulset: %s, replicas: %d/%d", name, int(ds.Status.ReadyReplicas), replicas)

		if int(ds.Status.ReadyReplicas) == replicas {
			t.Logf("found required replicas")
			return true, nil
		}

		return false, nil
	}); err != nil {
		t.Errorf("error while waiting for the statefulset: %v", err)
	}
}

func WaitForDaemonSet(t *testing.T, client kubernetes.Interface, ns, name string, retryInterval, timeout time.Duration) {
	if err := wait.Poll(retryInterval, timeout, func() (done bool, err error) {
		ds, err := client.AppsV1().DaemonSets(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				t.Logf("waiting for daemonset %s to be available", name)
				return false, nil
			}
			return false, err
		}
		replicas := ds.Status.DesiredNumberScheduled

		if replicas == 0 {
			t.Logf("no replicas scheduled for daemonset %s", name)

			return false, nil
		}

		t.Logf("daemonset: %s, replicas: %d/%d", name, ds.Status.DesiredNumberScheduled, replicas)
		if ds.Status.NumberReady == replicas {
			t.Logf("found required replicas")
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("error while waiting for the daemonset: %v", err)
	}
}

func WaitForDeployment(t *testing.T, client kubernetes.Interface, ns, name string, retryInterval, timeout time.Duration) {
	var err error
	var deploy *appsv1.Deployment

	// Check the readiness of the Deployment
	if err = wait.PollImmediate(retryInterval, timeout, func() (done bool, err error) {
		deploy, err = client.AppsV1().Deployments(ns).Get(name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				t.Logf("waiting for deployment %s to be available", name)
				return false, nil
			}
			return false, err
		}

		replicas := int(deploy.Status.Replicas)

		if replicas == 0 {
			t.Logf("no replicas scheduled for deployment %s", name)

			return false, nil
		}

		t.Logf("deployment: %s, replicas: %d/%d", name, int(deploy.Status.AvailableReplicas), replicas)

		if int(deploy.Status.AvailableReplicas) == replicas {
			t.Logf("found required replicas")
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("error while waiting for the deployment: %v", err)
		return
	}

	// Check the readiness of the pods
	labelSet := labels.Set(deploy.Labels)
	if err := wait.PollImmediate(retryInterval, timeout, func() (done bool, err error) {
		pods, err := client.CoreV1().Pods(ns).List(metav1.ListOptions{LabelSelector: labelSet.String()})
		if err != nil {
			return false, err
		}
		pods = filterNonControllerPods(pods)
		// go through each pod in the returned list and check the readiness status of it
		for _, pod := range pods.Items {
			for _, cs := range pod.Status.ContainerStatuses {
				if cs.RestartCount > 10 {
					return false, fmt.Errorf("pod: %s, container %s; pod in CrashLoopBackOff", pod.Name, cs.Name)
				}
				if !cs.Ready {
					t.Logf("pod: %s, container %s; container not ready", pod.Name, cs.Name)
					return false, nil
				}
			}
			t.Logf("pod %s, has all containers in ready state", pod.Name)
		}
		t.Logf("all pods for deployment %s, are in ready state", deploy.Name)
		return true, nil
	}); err != nil {
		t.Errorf("error while waiting for the pods: %v", err)
	}
}

func filterNonControllerPods(pods *corev1.PodList) *corev1.PodList {
	var filteredPods []corev1.Pod

	for _, pod := range pods.Items {
		// The pod that has a controller, has this label
		if _, ok := pod.Labels["pod-template-hash"]; !ok {
			continue
		}
		filteredPods = append(filteredPods, pod)
	}
	pods.Items = filteredPods
	return pods
}

type PortForwardInfo struct {
	readyChan     chan struct{}
	stopChan      chan struct{}
	portForwarder *portforward.PortForwarder

	// Port forwarding only works with a pod. Even if we ask user to provide the service name, we
	// still have to figure out the pod name and select one of the pod the service is pointing at.
	// Instead we can rely on the fact that prometheus pods are started by statefulset and will
	// always have consistent naming.
	PodName   string
	Namespace string
	PodPort   int
	LocalPort int
}

func (p *PortForwardInfo) CloseChan() {
	// This to guard against the closed channel, if you close the closed channel it panics this
	// piece of code guards against that
	select {
	case <-p.stopChan:
		return
	default:
	}
	close(p.stopChan)
}

func (p *PortForwardInfo) PortForward(t *testing.T) {
	config, err := buildKubeConfig(t)
	if err != nil {
		t.Fatalf("could not build config from KUBECONFIG: %v", err)
	}

	roundTripper, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		t.Fatalf("could not create round tripper: %v", err)
	}

	serverURL, err := url.Parse(config.Host)
	if err != nil {
		t.Fatalf("could not parse the URL from kubeconfig: %v", err)
	}

	serverURL.Path = fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", p.Namespace, p.PodName)
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: roundTripper}, http.MethodPost, serverURL)

	ports := []string{fmt.Sprintf("0:%d", p.PodPort)}

	out, errOut := new(bytes.Buffer), new(bytes.Buffer)
	p.stopChan, p.readyChan = make(chan struct{}, 1), make(chan struct{}, 1)

	forwarder, err := portforward.New(dialer, ports, p.stopChan, p.readyChan, out, errOut)
	if err != nil {
		p.CloseChan()
		t.Fatalf("could not create forwarder: %v", err)
	}

	p.portForwarder = forwarder

	// This function will print any error or output to stdout
	go func() {
		for range p.readyChan {
		}

		t.Logf("output of port forwarder:\n%s\n", out.String())

		if len(errOut.String()) != 0 {
			p.CloseChan()
			t.Errorf(errOut.String())
		}
	}()

	go func() {
		if err := forwarder.ForwardPorts(); err != nil { // Locks until stopChan is closed.
			p.CloseChan()
			t.Errorf("could not establish port forwarding: %v", err)
		}
	}()
}

func (p *PortForwardInfo) findLocalPort(t *testing.T) {
	forwardedPorts, err := p.portForwarder.GetPorts()
	if err != nil {
		t.Fatalf("could not get information about ports: %v", err)
	}

	if len(forwardedPorts) != 1 {
		t.Fatalf("number of forwarded ports not 1, currently forwarding on %d ports.", len(forwardedPorts))
	}

	p.LocalPort = int(forwardedPorts[0].Local)
}

func (p *PortForwardInfo) WaitUntilForwardingAvailable(t *testing.T) {
	const portForwardTimeout = 2

	// Wait until port forwarding is available
	select {
	case <-p.readyChan:
	case <-time.After(portForwardTimeout * time.Minute):
		t.Fatal("timed out waiting for port forwarding")
	}
	p.findLocalPort(t)
}

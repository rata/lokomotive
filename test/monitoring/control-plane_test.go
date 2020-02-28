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

// TODO: Add AWS here once prometheus-operator is fully supported there
// +build packet
// +build prometheus
package monitoring

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	testutil "github.com/kinvolk/lokomotive/test/components/util"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func TestPrometheusMetrics(t *testing.T) {
	testCases := []struct {
		componentName string
		query         string
	}{
		{
			componentName: "kube-apiserver",
			query:         "apiserver_request_total",
		},
		{
			componentName: "coredns",
			query:         "coredns_build_info",
		},
		{
			componentName: "kube-scheduler",
			query:         "scheduler_schedule_attempts_total",
		},
		{
			componentName: "kube-controller-manager",
			query:         "workqueue_work_duration_seconds_bucket",
		},
		{
			componentName: "kube-proxy",
			query:         "kubeproxy_sync_proxy_rules_duration_seconds_bucket",
		},
		{
			componentName: "kubelet",
			query:         "kubelet_running_pod_count",
		},
		{
			componentName: "calico-felix",
			query:         "felix_active_local_endpoints",
		},
	}

	stopChan, readyChan := make(chan struct{}, 1), make(chan struct{}, 1)

	p := &testutil.PortForwardInfo{
		PodName:   "prometheus-prometheus-operator-prometheus-0",
		Namespace: "monitoring",
		StopChan:  stopChan,
		ReadyChan: readyChan,
		PodPort:   9090,
	}

	go p.PortForward(t)
	defer func() {
		close(stopChan)
	}()

	// Wait until port forwarding is available
	for range readyChan {
	}

	p.FindLocalPort(t)

	promClient, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://127.0.0.1:%d", p.LocalPort),
	})
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("prometheus-%s", tc.componentName), func(t *testing.T) {
			t.Logf("querying %q", tc.query)

			v1api := v1.NewAPI(promClient)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			results, warnings, err := v1api.Query(ctx, tc.query, time.Now())
			if err != nil {
				t.Fatalf("error querying Prometheus: %v", err)
			}

			if len(warnings) > 0 {
				t.Logf("warnings: %v", warnings)
			}
			t.Logf("found %d results for %s", len(strings.Split(results.String(), "\n")), tc.query)
		})
	}
}

# Lokomotive AWS installation guide

## Contents

* [Introduction](#introduction)
* [Requirements](#requirements)
* [Step 1: Install Lokoctl](#step-1-install-lokoctl)
* [Step 2: Set up a working directory](#step-2-set-up-a-working-directory)
* [Step 3: Set up credentials from environment variables](#step-3-set-up-credentials-from-environment-variables)
* [Step 4: Define cluster configuration](#step-4-define-cluster-configuration)
* [Step 5: Create Lokomotive cluster](#step-5-create-lokomotive-cluster)
* [Verification](#verification)
* [Cleanup](#cleanup)
* [Troubleshooting](#troubleshooting)
* [Conclusion](#conclusion)
* [Next steps](#next-steps)

## Introduction

This quickstart guide walks through the steps needed to create a Lokomotive cluster on AWS with
Flatcar Container Linux using Route53 as the DNS provider.

By the end of this guide, you'll have a production-ready Kubernetes cluster running on AWS.

## Requirements

* Basic understanding of Kubernetes concepts.
* AWS account and IAM credentials.
* AWS Route53 DNS Zone (registered Domain Name or delegated subdomain).
* Terraform v0.12.x and [terraform-provider-ct](https://github.com/poseidon/terraform-provider-ct)
  installed locally
* An SSH key pair for management access.
* `kubectl` installed locally to access the Kubernetes cluster.

## Steps

### Step 1: Install lokoctl

Lokoctl is a command-line interface for Lokomotive.

To install `lokoctl`, follow the instructions in the [lokoctl installation](../installer/lokoctl.md)
guide.

### Step 2: Set up a working directory


It's better to start fresh in a new working directory, as the state of the cluster is stored in this
directory.

This also makes the cleanup task easier.


```console
mkdir -p lokomotive-infra/myawscluster
cd lokomotive-infra/myawscluster
```

### Step 3: Set up credentials from environment variables

The AWS credentials file can be found at `~/.aws/credentials` if you have set up and configured AWS
CLI before.  If you want to use that account, you don't need to specify any credentials for lokoctl.

You can also take any other credentials mechanism used by the AWS CLI but environment variables may
be the safest option. Either prepend them when starting lokoctl or export each of them once in the
current terminal session:

```console
AWS_ACCESS_KEY_ID=EXAMPLEID AWS_SECRET_ACCESS_KEY=EXAMPLEKEY lokoctl ...
```
or

```console
export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
export AWS_DEFAULT_REGION=us-east-1
```

### Step 4: Define cluster configuration

To create a Lokomotive cluster, we need to define a configuration.

A [production-ready configuration](../../examples/aws-production) is already provided for ease of
use.  Copy the example configuration to the working directory and modify accordingly.

The provided configuration installs the Lokomotive cluster and the following components:

* [metrics-server](../configuration-guides/components/metrics-server.md)
* [openebs-operator](../configuration-guides/components/openebs-operator.md)
* [cert-manager](../configuration-guides/components/cert-manager.md)
* [flatcar-linux-update-operator](../configuration-guides/components/flatcar-linux-upate-operator.md)
* [openebs-storage-class](../configuration-guides/components/openebs-storage-class.md)
* [prometheus-operator](../configuration-guides/components/prometheus-operator.md)

You can configure the components as per the requirements.

Create a variables file `lokocfg.vars` to define the variables present in the configuration file.

```console
#lokocfg.vars
ssh_public_keys = ["public-ssh-key-1", "public-ssh-key-2", ...]

state_s3_bucket = "name-of-the-s3-bucket-to-store-the-cluster-state"
lock_dynamodb_table = "name-of-the-dynamodb-table-for-state-locking"

dns_zone = "dns-zone-name"
route53_zone_id = "zone-id-of-the-dns-zone"

cert_manager_email = "email-address-used-for-cert-manager-component"
grafana_admin_password = "password-for-grafana"
```

For advanced cluster configurations and more information refer to the [AWS configuration
guide.](../configuration-guides/platforms/aws.md)

### Step 5: Create Lokomotive cluster

Run the following command to create the cluster:

```console
lokoctl cluster install
```
Once the command finishes, your Lokomotive cluster details are stored in the path you've specified
under `asset_dir`.

## Verification

A successful installation results in the output:

```console
module.aws-myawscluster.null_resource.bootkube-start: Still creating... [1m50s elapsed]
module.aws-myawscluster.null_resource.bootkube-start: Still creating... [2m0s elapsed]
module.aws-myawscluster.null_resource.bootkube-start: Creation complete after 2m6s [id=5156996152315868880]

Apply complete! Resources: 118 added, 0 changed, 0 destroyed.

Your configurations are stored in /home/imran/lokoctl-assets/myawscluster

Now checking health and readiness of the cluster nodes ...

Node              Ready    Reason          Message

ip-10-0-39-75     True     KubeletReady    kubelet is posting ready status
ip-10-0-39-78     True     KubeletReady    kubelet is posting ready status
ip-10-0-39-29     True     KubeletReady    kubelet is posting ready status
ip-10-0-12-241    True     KubeletReady    kubelet is posting ready status
ip-10-0-12-244    True     KubeletReady    kubelet is posting ready status
ip-10-0-12-249    True     KubeletReady    kubelet is posting ready status

Success - cluster is healthy and nodes are ready!
```

Use the generated `kubeconfig` file to access the Kubernetes cluster and list nodes.

```console
export KUBECONFIG=./lokomotive-assets/auth/kubeconfig
kubectl get nodes
```

## Cleanup

To destroy the Lokomotive cluster, execute the following commmand:

```console
lokoctl cluster destroy --confirm
```

You can safely delete the working directory created for this quickstart guide if you no longer
require it.

## Troubleshooting

### Stuck At Copy Controller Secrets

If there is an execution error or no progress beyond the output provided below:

```console
...
module.aws-myawscluster.null_resource.copy-controller-secrets: Still creating... (8m30s elapsed)
module.aws-myawscluster.null_resource.copy-controller-secrets: Still creating... (8m40s elapsed)
...
```

The error is because the `ssh_pubkeys` provided in the configuration is missing in the `ssh-agent`

To rectify the error, you need to:

1. Execute [Cleanup](#cleanup) step.
2. Follow the steps [to add the SSH key to the
   ssh-agent.](https://help.github.com/en/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent#adding-your-ssh-key-to-the-ssh-agent)
3. Retry [Step 5.](#step-5-create-lokomotive-cluster)

### IAM Permission Issues

  * If the failure is due to insufficient permissions, check the IAM policy.

## Conclusion

After walking through this guide, you've learned how to set up a Lokomotive cluster on AWS.

## Next steps

You now have an understanding of how to create a Lokomotive cluster on AWS.

This quickstart guide did not cover configuring all the available options in the cluster
configuration.  For more information, you can visit the [AWS configuration
guide](docs/configuration-guides/platforms/aws.md)


# ExternalDNS configuration reference for Lokomotive

## Contents

* [Introduction](#introduction)
* [Prerequisites](#prerequisites)
* [Configuration](#configuration)
* [Argument reference](#argument-reference)
* [Installation](#installation)
* [Uninstalling](#uninstalling)

## Introduction

[ExternalDNS](https://github.com/kubernetes-incubator/external-dns) is a Kubernetes addon that
synchronizes exposed Kubernetes Services and Ingresses with DNS providers to make them discoverable.

## Prerequisites

* A Lokomotive cluster accessible via `kubectl`.

* An ingress controller such as [Contour](contour.md) for HTTP ingress.

## Configuration

ExternalDNS component with Contour supports managing DNS records for Services of type `LoadBalancer`
only. More information on this limitation is explained in this
[issue](https://github.com/projectcontour/contour/issues/403).

ExternalDNS component currently supports AWS Route53 DNS provider.

ExternalDNS component configuration example:

```tf
component "external-dns" {
  # Required arguments.
  aws {
    # Required arguments
    zone_type = "public"
    zone_id = "ZQXH02G1EPZ6R"
    # Optional arguments.
    aws_access_key_id = ""
    aws_secret_access_key = ""
  }

  # Optional arguments.
  sources = ["service"]
  namespace = "external-dns"
  policy = "upsert-only"
  metrics = false
}
```

ExternalDNS manages DNS entries for the values in the field `ingress_hosts` of the [Contour
component](contour.md#argument-reference).

## Argument reference

Table of all the arguments accepted by the component.

Example:

| Argument                    | Description                                                                                                       | Default        | Required |
|-----------------------------|-------------------------------------------------------------------------------------------------------------------|:--------------:|:--------:|
| `sources`                   | Kubernetes resources type to be observed for new DNS entries by ExternalDNS.                                      | ["service"]    | false    |
| `namespace`                 | Namespace to install ExternalDNS.                                                                                 | "external-dns" | false    |
| `policy`                    | Modify how DNS records are sychronized between sources and providers (options: sync, upsert-only).                | "upsert-only"  | false    |
| `metrics`                   | Enable metrics collection by Prometheus. Needs [Prometheus Operator component](prometheus-operator.md) installed. | false          | false    |
| `aws`                       | Configuration block for AWS Route53 DNS provider.                                                                 | -              | true     |
| `aws.zone_type`             | Filter for zones of this type (options: public, private).                                                         | "public"       | false    |
| `aws.zone_id`               | ID of the DNS zone.                                                                                               | -              | true     |
| `aws.aws_access_key_id`     | AWS access key ID for AWS credentials. Use environment variable AWS_ACCESS_KEY_ID instead.                        | -              | false    |
| `aws.aws_secret_access_key` | AWS secret access key for AWS credentials. Use environment variable AWS_SECRET_ACCESS_KEY instead.                | -              | false    |

## Installation

To install the ExternalDNS component:

```bash
lokoctl component install external-dns
```
## Uninstalling

To uninstall the component:

```bash
lokoctl component render-manifest external-dns | kubectl delete -f -
```

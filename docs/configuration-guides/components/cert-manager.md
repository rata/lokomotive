# Cert-Manager configuration reference for Lokomotive

## Contents

* [Introduction](#introduction)
* [Prerequisites](#prerequisites)
* [Configuration](#configuration)
* [Argument reference](#argument-reference)
* [Installation](#installation)
* [Uninstalling](#uninstalling)

## Introduction

[cert-manager](https://cert-manager.io/docs/) is a Kubernetes service that provisions TLS
certificates from Let’s Encrypt and other certificate authorities and manages their lifecycles.

## Prerequisites

* A Lokomotive cluster accessible via `kubectl`.

## Configuration

If you run a cluster `enable_aggregation` set to `false`, make sure you disable the webhooks
feature, which will not work without aggregation enabled.

cert-manager component configuration example:

```tf
component "cert-manager" {
  email = "example@example.com"
  namespace = "cert-manager"
  webhooks = false
}
```

## Argument reference

Table of all the arguments accepted by the component.

Example:

| Argument    | Description                                                  | Default      | Required |
|-------------|--------------------------------------------------------------|:------------:|:--------:|
| `email`     | Email used for certificates to receive expiry notifications. | -            | true     |
| `namespace` | Namespace to deploy the cert-manager into.                   | cert-manager | false    |
| `webhooks`  | Controls if webhooks should be deployed.                     | true         | false    |

## Installation

To install the cert-manager component:

```bash
lokoctl component install cert-manager
```
## Uninstalling

To uninstall the component:

```bash
lokoctl component render-manifest cert-manager | kubectl delete -f -
```

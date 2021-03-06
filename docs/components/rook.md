# Rook

[Rook](https://rook.io) is a storage orchestrator for Kubernetes.

## Installation

#### `.lokocfg` file

Declare the component's config:

```
component "rook" {
  namespace = "rook-test"

  node_selector {
    key      = "node-role.kubernetes.io/node"
    operator = "Exists"
  }

  node_selector {
    key      = "storage.lokomotive.io"
    operator = "In"

    # If the `operator` is set to `"In"`, `values` should be specified.
    values = [
      "foo",
    ]
  }

  toleration {
    key      = "storage.lokomotive.io"
    operator = "Equal"
    value    = "rook-ceph"
    effect   = "NoSchedule"
  }

  agent_toleration_key    = "storage.lokomotive.io"
  agent_toleration_effect = "NoSchedule"

  discover_toleration_key    = "storage.lokomotive.io"
  discover_toleration_effect = "NoSchedule"
}
```

## Argument Reference

| Argument | Explanation | Default | Required |
|----------|-------------|---------|----------|
| `namespace` | Namespace to deploy the rook operator into. | rook | false |
| `node_selector` | Node selectors for deploying the operator pod | - | false |
| `toleration` | Tolerations that the operator's pods will tolerate | - | false |
| `agent_toleration_key` | Toleration key for the rook agent pods | - | false |
| `agent_toleration_effect` | Toleration effect for the rook agent pods. Needs to be specified if `agent_toleration_key` is set. | - | false |
| `discover_toleration_key` | Toleration key for the rook discover pods | - | false |
| `discover_toleration_effect` | Toleration effect for the rook discover pods. Needs to be specified if `discover_toleration_key` is set. | - | false |

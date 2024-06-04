---
subcategory: "Lynx snapshot"
layout: "lynx"
page_title: "Provider: Lynx"
page_title: "Lynx Resource Manager: lynx_snapshot"
description: |-
  Manage local snapshots in Lynx.
---

# lynx_snapshot

Manage local snapshots in Lynx.

## Example Usage

```hcl
resource "lynx_snapshot" "my_snapshot" {
  title       = "Grafana Project Snapshot"
  description = "Grafana Project Snapshot"
  record_type = "project"
  record_id   = lynx_project.grafana.id

  team = {
    id = lynx_team.monitoring.id
  }
}
```

## Arguments Reference

The following arguments are supported:

* `title` - (Required) Set Team's `name` for your Lynx Snapshot resource.
* `description` - (Required) Description of your Lynx Snapshot resource.
* `record_type`  - (Required) 
* `team`  - (Required) Team ID linked with your Lynx Snapshot resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the lynx_Snapshot.
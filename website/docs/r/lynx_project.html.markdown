---
subcategory: "Lynx project"
layout: "lynx"
page_title: "Provider: Lynx"
page_title: "Lynx Resource Manager: lynx_project"
description: |-
  Manage Projects in Lynx.
---

# lynx_project

Manage local Projects in Lynx.

## Example Usage

```hcl
resource "lynx_project" "grafana" {
  name        = "Grafana"
  slug        = "grafana"
  description = "Grafana Project"

  team = {
    id = lynx_team.monitoring.id
  }
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The display Project `name` for your Lynx project resource.
* `slug` - (Required) Set `email` for your Lynx Project resource.
* `description` - (Required) Description of your Lynx Project resource.
* `team`  - (Required) Team ID linked with your Lynx Project resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the lynx__project.
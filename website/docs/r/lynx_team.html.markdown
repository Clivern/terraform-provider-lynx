---
subcategory: "Lynx Team"
layout: "lynx"
page_title: "Provider: Lynx"
page_title: "Lynx Resource Manager: lynx_team"
description: |-
  Manage local Teams in Lynx.
---

# lynx_team

Manage local Teams in Lynx.

## Example Usage

```hcl
resource "lynx_team" "monitoring" {
  name        = "Monitoring"
  slug        = "monitoring"
  description = "System Monitoring Team"

  members = [
    lynx_user.stella.id,
    lynx_user.skylar.id,
    lynx_user.erika.id,
    lynx_user.adriana.id
  ]
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) Set Team's `name` for your Lynx Team resource.
* `slug` - (Required) Set `email` for your Lynx Team resource.
* `description` - (Required) Description of your Lynx Team resource.
* `members`  - (Required) List of User ID's to include in your Lynx Team resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the lynx_Team.
---
subcategory: "Lynx Environment"
layout: "lynx"
page_title: "Provider: Lynx"
page_title: "Lynx Resource Manager: lynx_environment"
description: |-
  Manage local Environments in Lynx.
---

# lynx_environment

Manage local Environments in Lynx.

## Example Usage

```hcl
resource "lynx_environment" "prod" {
  name     = "Development"
  slug     = "dev"
  username = "~username-here~"
  secret   = "~secret-here~"

  project = {
    id = lynx_project.grafana.id
  }
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The display Environment `name` for your Lynx Environment resource.
* `slug` - (Required) Set `slug` for your Lynx Environment resource.
* `username` - (Required) Set `username` for Lynx Environment resource.
* `secret`  - (Required) Set `secret` for Lynx Environment resource.
* `project`  - (Required) Project ID linked with Lynx Environment resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the lynx_environment.
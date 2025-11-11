---
subcategory: "Lynx User"
layout: "lynx"
page_title: "Provider: Lynx"
page_title: "Lynx Resource Manager: lynx_user"
description: |-
  Manage local users in Lynx.
---

# lynx_user

Manage local users in Lynx.

## Example Usage

```hcl
resource "lynx_user" "stella" {
  name     = "Stella"
  email    = "stella@example.com"
  role     = "regular"
  password = "~password-here~"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The display `name` for your Lynx User resource.
* `email` - (Required) Set `email` for your Lynx User resource. 
* `role` - (Required) Set Lynx `Role` for your Lynx User resource. 
* `password` - (Required) Set `password` for your Lynx User resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the lynx_user.

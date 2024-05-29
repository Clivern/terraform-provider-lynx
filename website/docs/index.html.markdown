---
layout: "lynx"
page_title: "Provider: Lynx"
description: |-
  The Lynx Provider is used to interact with the many resources supported by Lynx (A Fast, Secure and Reliable Terraform Backend) through its APIs.

---

# Lynx Provider

The Lynx Provider can be used to configure infrastructure in [Lynx Terraform backend](https://github.com/Clivern/Lynx/blob/main/README.md) using the Lynx API's. Documentation regarding the [Data Sources](/docs/configuration/data-sources.html) and [Resources](/docs/configuration/resources.html) supported by the Lynx Provider can be found in the navigation to the left.

To learn the basics of Terraform using this provider, follow the
hands-on [get started tutorials](https://github.com/Clivern/Lynx/blob/main/README.md).

## Authenticating to Azure

## Example Usage


## Bugs and Feature Requests

The provider maintainers will often use the assignee field on an issue to mark
who is working on it.

* An issue assigned to an individual maintainer indicates that the maintainer is working
on the issue

* If you're interested in working on an issue please leave a comment on that issue

---

If you have configuration questions, or general questions about using the provider, try checking out:

* [Terraform's community resources](https://www.terraform.io/docs/extend/community/index.html)
* [HashiCorp support](https://support.hashicorp.com) for Terraform Enterprise customers

## Argument Reference

The following arguments are supported:

* `client_id` - (Optional).

-> **Note:** When Terraform is configured to use credentials with limited permissions you *must* set `skip_provider_registration` to true (or the environment variable `ARM_SKIP_PROVIDER_REGISTRATION=true`) in order to account for this - otherwise Terraform will, as described above, try to register any Resource Providers.


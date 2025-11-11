---
layout: "lynx"
page_title: "Provider: Lynx"
description: |-
  The Lynx Provider is used to interact with the many resources supported by Lynx (A Fast, Secure and Reliable Terraform Backend) through its APIs.

---

# Lynx Provider

The Lynx Provider can be used to configure infrastructure in [Lynx Terraform backend](https://github.com/Clivern/Lynx/blob/main/README.md) using the Lynx API's. Documentation regarding the Data Sources and Resources supported by the Lynx Provider can be found in the navigation to the left.

To learn the basics of Terraform using this provider, follow the hands-on [Get started tutorial](https://github.com/Clivern/Lynx/blob/main/README.md).

## Authenticating to Lynx

### Example Usage

```terraform
# Configure the Lynx Provider 
provider "lynx" {
  api_url = "http://localhost:4000/api/v1"
  api_key = "~api key here~"
}
# Create the resources
```

~> **IMPORTANT** Hard-coding yourLynx programmatic API key pair into a Terraform configuration is not recommended.
Consider the risks, especially the inadvertent submission of a configuration file containing secrets to a public repository.

### Environment Variables

You can also provide your credentials via the environment variables, 
Lynx programmatic API key pair respectively:

Usage (prefix the export commands with a space to avoid the keys being recorded in OS history):

```shell
$ export LYNX_API_URL="xxxx"
$ export LYNX_API_KEY="xxxx"
$ terraform plan
```
Please find more details in [`Lynx documentation`](https://lynx.clivern.com/documentation/api-and-tf-provider/).

## Argument Reference
In order to set up authentication with the Lynx provider, you must generate a programmatic API key for Lynx.

The following arguments are supported:

* `api_url` - (Required) A `api_url` block to be used to set the Lynx API endpoint.
* `api_key` - (Required) A `api_key` block to be used for the authentication in the Lynx API. 

## Helpful Links/Information

* Getting started with [Lynx](https://lynx.clivern.com/documentation/getting-started/).
* Changelogs for each release version of `terraform-provider-lynx` [here](https://github.com/clivern/terraform-provider-lynx/releases).

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
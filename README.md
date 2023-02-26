<p align="center">
    <img alt="Lynx Logo" src="https://www.vectorlogo.zone/logos/hashicorp/hashicorp-icon.svg" width="180" />
    <h3 align="center"></h3>
    <p align="center">
        <a href="https://github.com/Clivern/terraform-provider-lynx/actions/workflows/test.yml">
            <img src="https://github.com/Clivern/terraform-provider-lynx/actions/workflows/test.yml/badge.svg"/>
        </a>
        <a href="https://github.com/clivern/terraform-provider-lynx/releases">
            <img src="https://img.shields.io/badge/Version-0.2.0-1abc9c.svg">
        </a>
        <a href="https://github.com/clivern/terraform-provider-lynx/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>


### Usage

Here is an example to setup team, team users, project, project environment and a snapshot.

```hcl
terraform {
  required_providers {
    lynx = {
      source = "Clivern/lynx"
      version = "0.2.0"
    }
  }
}

provider "lynx" {
  api_url = "http://localhost:4000/api/v1"
  api_key = "~api key here~"
}

resource "lynx_user" "stella" {
  name     = "Stella"
  email    = "stella@example.com"
  role     = "regular"
  password = "~password-here~"
}

resource "lynx_user" "skylar" {
  name     = "Skylar"
  email    = "skylar@example.com"
  role     = "regular"
  password = "~password-here~"
}

resource "lynx_user" "erika" {
  name     = "Erika"
  email    = "erika@example.com"
  role     = "regular"
  password = "~password-here~"
}

resource "lynx_user" "adriana" {
  name     = "Adriana"
  email    = "adriana@example.com"
  role     = "regular"
  password = "~password-here~"
}

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

resource "lynx_project" "grafana" {
  name        = "Grafana"
  slug        = "grafana"
  description = "Grafana Project"

  team = {
    id = lynx_team.monitoring.id
  }
}

resource "lynx_environment" "prod" {
  name     = "Development"
  slug     = "dev"
  username = "~username-here~"
  secret   = "~secret-here~"

  project = {
    id = lynx_project.grafana.id
  }
}

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


### Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, `terraform-provider-lynx` is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/terraform-provider-lynx/releases) for changelogs for each release version of `terraform-provider-lynx`. It contains summaries of the most noteworthy changes made in each release. Also see the [Milestones section](https://github.com/clivern/terraform-provider-lynx/milestones) for the future roadmap.


### Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/terraform-provider-lynx/issues


### Security Issues

If you discover a security vulnerability within `terraform-provider-lynx`, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


### Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


### License

© 2023, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**terraform-provider-lynx** is authored and maintained by [@clivern](http://github.com/clivern).

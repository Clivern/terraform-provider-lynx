terraform {
  required_providers {
    lynx = {
      source = "registry.terraform.io/clivern/lynx"
    }
  }
}

provider "lynx" {
  api_url = "http://localhost:4000/api/v1"
  api_key = "bd11a454-a694-49c8-b3da-0fe6cf48a27d"
}

resource "lynx_user" "stella" {
  name     = "Stella Doe"
  email    = "stella@example.com"
  role     = "regular"
  password = "$87272663625"
}

output "user_id" {
  value = lynx_user.stella.id
}

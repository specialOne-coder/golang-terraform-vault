terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {
}

variable "docker_image" {
  type    = string
  default = "hashicorp/vault:latest"
}

resource "docker_container" "vault_container" {
  name  = "vault_container"
  image = var.docker_image
  ports {
    internal = 8200
    external = 8200
  }
}

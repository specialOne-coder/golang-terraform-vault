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
  default = "nginx:latest"
}


resource "docker_container" "nginx_container" {
  name  = "nginx_container"
  image = var.docker_image
  ports {
    internal = 80
    external = 8080
  }
}

resource "null_resource" "curl_nginx" {
  depends_on = [docker_container.nginx_container]

  provisioner "local-exec" {
    command = "curl -s http://localhost:8080 > index.html"
  }
}

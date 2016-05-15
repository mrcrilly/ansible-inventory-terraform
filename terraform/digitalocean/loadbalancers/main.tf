
resource "digitalocean_droplet" "haproxy-primary" {
  image = "centos-7-0-x64"
  name = "haproxy-primary"
  region = "sgp1"
  size = "512mb"
  ssh_keys = ["1207665"]
  private_networking = true
}

resource "digitalocean_floating_ip" "primary" {
  droplet_id = "${digitalocean_droplet.haproxy-primary.id}"
  region = "sgp1"
}

resource "digitalocean_droplet" "haproxy-secondary" {
  image = "centos-7-0-x64"
  name = "haproxy-secondary"
  region = "sgp1"
  size = "512mb"
  ssh_keys = ["1207665"]
  private_networking = true
}

resource "digitalocean_floating_ip" "secondary" {
  droplet_id = "${digitalocean_droplet.haproxy-secondary.id}"
  region = "sgp1"
}

resource "digitalocean_droplet" "haproxy-tertiary" {
  image = "centos-7-0-x64"
  name = "haproxy-tertiary"
  region = "sgp1"
  size = "512mb"
  ssh_keys = ["1207665"]
  private_networking = true
}

resource "digitalocean_floating_ip" "tertiary" {
  droplet_id = "${digitalocean_droplet.haproxy-tertiary.id}"
  region = "sgp1"
}

resource "digitalocean_droplet" "web" {
  count = 2
  image = "centos-7-0-x64" # CentOS 7
  name = "web-${count.index}"
  region = "sgp1" # Singapore
  size = "512mb"
  ssh_keys = ["1207665"] # [MBP Home]
  private_networking = true
}

module "db" {
  source = "./database"
}
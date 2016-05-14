
resource "digitalocean_droplet" "database" {
  count = 1
  image = "centos-7-0-x64" # CentOS 7
  name = "database-${count.index}"
  region = "sgp1" # Singapore
  size = "512mb"
  ssh_keys = ["1207665"] # [MBP Home]
  private_networking = true
}
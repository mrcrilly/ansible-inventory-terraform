
resource "digitalocean_droplet" "haproxy" {
  count = 1
  image = "centos-7-0-x64" # CentOS 7
  name = "haproxy-${count.index}"
  region = "sgp1" # Singapore
  size = "512mb"
  ssh_keys = ["1207665"] # [MBP Home]
  private_networking = true
}
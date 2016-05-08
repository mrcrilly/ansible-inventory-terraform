
provider "digitalocean" {}

module "webservers" {
  source = "./webserver"
}

module "loadbalancers" {
  source = "./loadbalancer"
}

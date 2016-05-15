
provider "digitalocean" {}

module "webservers" {
  source = "./webserver"
}

module "webservers_dev" {
  source = "./webserver"
}

module "loadbalancers" {
  source = "./loadbalancers"
}

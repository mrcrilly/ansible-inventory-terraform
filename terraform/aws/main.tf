
provider "aws" {}

resource "aws_vpc" "primary" {
  cidr_block = "10.1.0.0/16"  
}

resource "aws_internet_gateway" "primary" {
  vpc_id = "${aws_vpc.primary.id}"
}

resource "aws_subnet" "primary" {
  vpc_id = "${aws_vpc.primary.id}"
  cidr_block = "10.1.1.0/24"
}

# module "webservers" {
#   source = "./webserver"
#   subnet = "${aws_subnet.primary.id}"
# }

module "loadbalancers" {
  source = "./loadbalancer"
  subnet = "${aws_subnet.primary.id}"
}

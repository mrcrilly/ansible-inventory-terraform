
variable "subnet" {}

resource "aws_instance" "web" {
  ami = "ami-fedafc9d"
  count = 1
  instance_type = "t2.micro"
  subnet_id = "${var.subnet}"
}

resource "aws_instance" "web_dev" {
  ami = "ami-fedafc9d"
  count = 1
  instance_type = "t2.micro"
  subnet_id = "${var.subnet}"
  tags = {
    Name = "web_dev"
  }
}

resource "aws_eip" "web_dev" {
  instance = "${aws_instance.web_dev.id}"
  vpc      = true
}

module "db" {
  source = "./database"
  subnet = "${var.subnet}"
}
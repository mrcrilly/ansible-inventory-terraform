variable "subnet" {}

resource "aws_instance" "db" {
  ami = "ami-fedafc9d"
  count = 1
  instance_type = "t2.micro"
  subnet_id = "${var.subnet}"
}
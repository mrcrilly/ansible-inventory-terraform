
variable "subnet" {}

resource "aws_instance" "haproxy" {
  count = 1
  ami = "ami-fedafc9d"
  instance_type = "t2.micro"
  subnet_id = "${var.subnet}"
}

resource "aws_eip" "haproxy" {
  instance = "${aws_instance.haproxy.id}"
  vpc      = true
}
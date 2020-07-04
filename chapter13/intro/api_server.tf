provider "aws" {
  profile = "blocksys"
  region = "eu-central-1"
}

resource "aws_instance" "api_server" {
  ami = "ami-03818140b4ac9ae2b"
  instance_type = "t2.micro"
  key_name = aws_key_pair.api_server_key.key_name
  vpc_security_group_ids = [
    "sg-0fd3aee7c6e724bec"]
  subnet_id = "subnet-0d63449d9243f9bff"
}

resource "aws_key_pair" "api_server_key" {
  key_name = "api-server-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC/y20VjbLB2um4pQ5vfAJNFYMEGh6DIwrKbdAIkghzxG6bbhR54ApBnUrtnK+rHBZ0NQ+yHAwh5wcmEwuhBw2vi9hSDIbN9MuW+DAaam/a1Xq9dJ/2P7p3gNwFCMxUGJpfg997Fjdqt2V6l/hvVEZ8YE/gTGWgieQXH36QDGdpOcZYvtYx4C6FJQtEZrYPDwl1KIub8wi7oeNBTl2slYYJAc45YrIvUvt06rtIrThQriw+vzIKS5IecQZA67AzDXZNCcCQ6IuGc17BU0Q5WP6Vbjmbq8/ucgZ0p4qwnMRj3WJaXras5BEWlpw4jeN+/dpxetniX40PvrXbHfTvKgi7 jerzylasyk@Jerzys-MacBook-Pro.local"
}

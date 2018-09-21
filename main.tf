terraform {
  backend "local" {
    path = "terraform.d/terraform.tfstate"
  }
}

provider "onix" {}

resource "onix_itemtype" "AWS Environment" {
  name        = "AWS Environment"
  description = "AWS Environment"
}



resource "onix_itemtype" "AWS VPC" {
  name        = "AWS VPC"
  description = "AWS VPC"
}

resource "onix_itemtype" "AWS Subnet" {
  name        = "AWS Subnet"
  description = "AWS Subnet"
}


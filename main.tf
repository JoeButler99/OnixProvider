terraform {
  backend "local" {
    path = "terraform.d/terraform.tfstate"
  }
}

provider "onix" {}

resource "onix_itemtype" "AWS_Environment" {
  name        = "AWS Environment"
  description = "AWS Environment"
}

resource "onix_itemtype" "AWS_VPC" {
  name        = "AWS_VPC"
  description = "AWS_VPC"
}

resource "onix_itemtype" "AWS_Subnet" {
  name        = "AWS_Subnet"
  description = "AWS_Subnet"
}

resource "onix_item" "CTRL_VPC_ID" {
  name        = "CTRLVPCID"
  description = ""
  itemtype    = "${onix_itemtype.AWS_VPC.id}"
}

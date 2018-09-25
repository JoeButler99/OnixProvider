terraform {
  backend "local" {
    path = "terraform.d/terraform.tfstate"
  }
}

provider "onix" {}

resource "onix_itemtype" "AWS_Environment" {
  key         = "AWSENV"
  name        = "AWS Environment"
  description = "AWS Environment"
}

resource "onix_itemtype" "AWS_VPC" {
  key         = "AWSVPC"
  name        = "AWS VPC"
  description = "AWS VPC Details"
}

resource "onix_itemtype" "AWS_Subnet" {
  key         = "AWSSUB2"
  name        = "AWS Subnet"
  description = "AWS Subnet Details"
}

resource "onix_item" "CTRL_VPC_ID" {
  key         = "CTRLVPCID"
  name        = "CTRL VPC ID"
  description = "CTRL VPC ID"
  itemtype    = "${onix_itemtype.AWS_VPC.key}"

  meta = <<JSON
{ "I WORK": true }
JSON
}

resource "onix_itemtype" "CtrlSSHKey" {
  key         = "EnvSshKey"
  name        = "Environment SSH Key"
  description = "Public Key for SSH used with control VMS"
}

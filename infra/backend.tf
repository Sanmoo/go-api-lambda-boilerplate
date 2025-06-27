terraform {
  backend "s3" {
    bucket = "deployer-terraform-states"
    key    = "media-tracker/common-infra"
    region = "us-east-1"
  }
}
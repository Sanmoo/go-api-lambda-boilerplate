terraform {
  backend "s3" {
    bucket = "deployer-terraform-states"
    key    = "media-tracker/lambdas/books"
    region = "us-east-1"
  }
}
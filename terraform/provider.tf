terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.38.0"
    }
    
    google = {
      source = "hashicorp/google"
      version = "4.11.0"
    }
  }
}

provider "google" {
#   credentials = file(format("%s/%s/%s",path.module, "terraform", var.gcp_credentials))
  credentials = file("deployment-account.json")
  project     = var.project
  region      = var.gcp_region
}
terraform {
  backend "gcs" {
    bucket = "obsec-tf-state"
    prefix = "tf/state"
  }
}
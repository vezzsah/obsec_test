resource "google_storage_bucket" "my-bucket" {
  name                     = "bucket-demo-wip-cred"
  location                 = "us-central1"
  project                  = "obsec-challenge"
  force_destroy            = true
  public_access_prevention = "enforced"
}


resource "google_storage_bucket" "tf-import" {
  name          = "obsec-tf-import-node-pool"
  location      = var.region
  force_destroy = true

  lifecycle_rule {
    condition {
      age = 3
    }
    action {
      type = "Delete"
    }
  }
}
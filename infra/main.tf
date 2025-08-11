resource "google_service_account" "default" {
  account_id   = "gke-sa-obsec"
  display_name = "GKE Service Account"
  project      = var.project
}

resource "google_container_cluster" "primary" {
  name                      = "obsec-gke-cluster"
  location                  = var.zone
  default_max_pods_per_node = 2

  release_channel {
    channel = "STABLE"
  }
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "kubernetes_namespace" "obsec-gke-namespace" {
  metadata {
    name = var.project
  }
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "my-node-pool"
  location   = google_container_cluster.primary.location
  cluster    = google_container_cluster.primary.name
  node_count = 1
  autoscaling {
    max_node_count = 1
  }

  node_config {
    preemptible     = true
    machine_type    = "e2-micro"
    disk_size_gb    = 10
    service_account = google_service_account.default.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }

}
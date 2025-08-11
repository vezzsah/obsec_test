resource "google_service_account" "gke-sa" {
  account_id   = "gke-sa-obsec"
  display_name = "GKE Service Account"
  project      = var.project
}

resource "google_container_cluster" "primary" {
  name                      = "obsec-gke-cluster"
  location                  = var.zone
  default_max_pods_per_node = 8
  remove_default_node_pool  = true
  initial_node_count        = 1
  network                   = google_compute_network.vpc.self_link
  subnetwork                = google_compute_subnetwork.private.self_link
  networking_mode           = "VPC_NATIVE"
  deletion_protection       = false

  addons_config {
    http_load_balancing {
      disabled = false
    }
    horizontal_pod_autoscaling {
      disabled = false
    }
  }

  release_channel {
    channel = "REGULAR"
  }

  workload_identity_config {
    workload_pool = "${var.project}.svc.id.goog"
  }

}

resource "google_container_node_pool" "obsec-challenge-pool" {
  
  name     = "obsec-challenge-nodes"
  location = google_container_cluster.primary.location
  cluster  = google_container_cluster.primary.name
  autoscaling {
    min_node_count = 1
    max_node_count = 3
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }


  network_config {
    create_pod_range    = true
    pod_ipv4_cidr_block = google_compute_subnetwork.private.ip_cidr_range

  }

  node_config {
    preemptible  = false
    machine_type = "e2-micro"

    service_account = google_service_account.gke-sa.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
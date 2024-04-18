# GKE cluster
data "google_container_engine_versions" "gke_version" {
  location = var.region
  version_prefix = "1.27."
}
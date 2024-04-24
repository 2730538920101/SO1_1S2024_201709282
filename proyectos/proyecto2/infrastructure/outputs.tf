output "region" {
  value       = var.region
  description = "GCloud Region"
}

output "project_id" {
  value       = var.project_id
  description = "GCloud Project ID"
}

output "kubernetes_cluster_name" {
  value       = google_container_cluster.primary.name
  description = "GKE Cluster Name"
}

output "kubernetes_cluster_host" {
  value       = google_container_cluster.primary.endpoint
  description = "GKE Cluster Host"
}

# Output para la URL del servicio de Cloud Run para el frontend
output "frontend_service_url" {
  value = google_cloud_run_service.frontend_service.status[0].url
  description = "La URL del servicio de Cloud Run para el frontend"
}

# Output para la URL del servicio de Cloud Run para el backend
output "backend_service_url" {
  value = google_cloud_run_service.backend_service.status[0].url
  description = "La URL del servicio de Cloud Run para el backend"
}


output "artifact_registry_repository_url" {
  description = "Artifact Registry repository URL"
  value       = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.iscool_gpt.repository_id}"
}

output "staging_url" {
  description = "Staging Cloud Run service URL"
  value       = google_cloud_run_service.iscool_gpt_staging.status[0].url
}

output "production_url" {
  description = "Production Cloud Run service URL"
  value       = google_cloud_run_service.iscool_gpt_production.status[0].url
}

output "staging_service_name" {
  description = "Staging Cloud Run service name"
  value       = google_cloud_run_service.iscool_gpt_staging.name
}

output "production_service_name" {
  description = "Production Cloud Run service name"
  value       = google_cloud_run_service.iscool_gpt_production.name
}

output "cloud_run_service_account_email" {
  description = "Cloud Run service account email"
  value       = google_service_account.cloud_run_sa.email
}

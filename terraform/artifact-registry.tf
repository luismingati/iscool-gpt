resource "google_artifact_registry_repository" "iscool_gpt" {
  location      = var.region
  repository_id = var.artifact_registry_repository
  description   = "Docker repository for isCool GPT API"
  format        = "DOCKER"

  labels = {
    environment = "multi"
    app         = "iscool-gpt"
  }
}

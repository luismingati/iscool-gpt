resource "google_cloud_run_service" "iscool_gpt_staging" {
  name     = "iscool-gpt-staging"
  location = var.region

  template {
    spec {
      service_account_name = google_service_account.cloud_run_sa.email

      containers {
        # This will be updated by GitHub Actions on deployment
        # Initial placeholder image
        image = "gcr.io/cloudrun/hello"

        ports {
          container_port = 8080
        }

        env {
          name  = "GEMINI_API_KEY"
          value = var.gemini_api_key
        }

        env {
          name  = "RATE_LIMIT_REQUESTS"
          value = var.rate_limit_requests
        }

        env {
          name  = "RATE_LIMIT_WINDOW"
          value = var.rate_limit_window
        }

        env {
          name  = "ENVIRONMENT"
          value = "staging"
        }

        resources {
          limits = {
            cpu    = "1000m"
            memory = "512Mi"
          }
        }
      }
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "10"
        "autoscaling.knative.dev/minScale" = "0"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

# Make the service publicly accessible
resource "google_cloud_run_service_iam_member" "staging_public_access" {
  service  = google_cloud_run_service.iscool_gpt_staging.name
  location = google_cloud_run_service.iscool_gpt_staging.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

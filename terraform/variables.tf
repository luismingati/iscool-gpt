variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP Region"
  type        = string
  default     = "southamerica-east1"
}

variable "credentials_file" {
  description = "Path to the service account key file"
  type        = string
}

variable "artifact_registry_repository" {
  description = "Artifact Registry repository name"
  type        = string
  default     = "iscool-gpt"
}

variable "gemini_api_key" {
  description = "Gemini API Key for the application"
  type        = string
  sensitive   = true
}

variable "rate_limit_requests" {
  description = "Rate limit requests per window"
  type        = string
  default     = "10"
}

variable "rate_limit_window" {
  description = "Rate limit time window"
  type        = string
  default     = "60s"
}


resource "google_project" "go_project" {
  name       = "go-app"
  project_id = "go-demo-app"
}

resource "google_service_account" "go-deployment-account" {
  account_id   = "go-deployment-account"
  display_name = "Go deployment account"
}

resource "google_project_service" "run" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "build" {
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "artifactregistry" {
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "sqladmin" {
  service            = "sqladmin.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "containerregistry" {
  service            = "containerregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "iamcredentials" {
  service            = "iamcredentials.googleapis.com"
  disable_on_destroy = false
}

resource "google_service_account_key" "go-deployment-account_key" {
  service_account_id = google_service_account.go-deployment-account.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}

resource "google_project_iam_binding" "service_account_role_service_agent" {
  project = google_project.go_project.project_id
  role    = "roles/cloudbuild.serviceAgent"
  members = [
    "serviceAccount:${google_service_account.go-deployment-account.email}"
  ]
}

resource "google_project_iam_binding" "service_account_role_cloudsql_admin" {
  project = google_project.go_project.project_id
  role    = "roles/cloudsql.admin"
  members = [
    "serviceAccount:${google_service_account.go-deployment-account.email}"
  ]
}

resource "google_project_iam_binding" "service_account_roles_cloudsql_client" {
  project = google_project.go_project.project_id
  role    = "roles/cloudsql.client"
  members = [
    "serviceAccount:${google_service_account.go-deployment-account.email}"
  ]
}

resource "google_project_iam_binding" "service_account__role_owner" {
  project = google_project.go_project.project_id
  role    = "roles/owner"
  members = [
    "serviceAccount:${google_service_account.go-deployment-account.email}"
  ]
}

resource "google_sql_database_instance" "gcp_sql_postgres" {
  name             = var.gcp_pg_name
  database_version = var.gcp_pg_database_version
  region           = var.gcp_pg_region
  deletion_protection = false

  settings {
    tier = var.gcp_pg_tier

    database_flags {
      name  = var.gcp_pg_db_flag_name
      value = var.gcp_pg_db_flag_value
    }
  }
  
  timeouts {
    create = "60m"
  }
}

resource "google_sql_user" "users" {
  name     = "postgres"
  instance = google_sql_database_instance.gcp_sql_postgres.name
  password = "postgres"
}

resource "google_cloud_run_service" "go-app_service" {
  name     = "go-app-service"
  location = var.gcp_pg_region

  template {
    spec {
      containers {
        image = "gcr.io/${google_project.go_project.project_id}/go-run-service/${google_project.go_project.name}:latest"
      }
    }
  }
}
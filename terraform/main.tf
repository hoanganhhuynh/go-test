
# resource "google_project" "go_project" {
#   name       = "Go-App"
#   project_id = "go-app-396907"
# }



# resource "google_service_account_key" "go_service_account_key" {
#   service_account_id = google_service_account.go_deployment_account.name
# }

resource "google_project_service" "cloud-resource-manager" {
  service            = "cloudresourcemanager.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "cloud-service-usage" {
  service            = "serviceusage.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "run" {
  service            = "run.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "build" {
  service            = "cloudbuild.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "artifactregistry" {
  service            = "artifactregistry.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "sqladmin" {
  service            = "sqladmin.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "containerregistry" {
  service            = "containerregistry.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_project_service" "iamcredentials" {
  service            = "iamcredentials.googleapis.com"
  project            = var.go_app_project_id
  disable_on_destroy = false
}

resource "google_service_account" "go_deployment_account" {
  account_id   = "go-deployment-account"
  display_name = "Go deployment account"
  project      = var.go_app_project_id
}

# resource "google_service_account_key" "go-deployment-account_key" {
#   service_account_id = google_service_account.go-deployment-account.name
#   public_key_type    = "TYPE_X509_PEM_FILE"
# }

# resource "google_project_iam_binding" "service_account_role_project_creator" {
#   project = google_project.go_project.project_id
#   role    = "roles/resourcemanager.projectCreator"
#   members = [
#     "serviceAccount:${google_service_account.go_deployment_account.email}"
#   ]
# }

resource "google_project_iam_binding" "service_account_role_service_agent" {
  project = var.go_app_project_id
  role    = "roles/cloudbuild.serviceAgent"
  members = [
    "serviceAccount:${google_service_account.go_deployment_account.email}"
  ]
}

resource "google_project_iam_binding" "service_account_role_cloudsql_admin" {
  project = var.go_app_project_id
  role    = "roles/cloudsql.admin"
  members = [
    "serviceAccount:${google_service_account.go_deployment_account.email}"
  ]
}

resource "google_project_iam_binding" "service_account_roles_cloudsql_client" {
  project = var.go_app_project_id
  role    = "roles/cloudsql.client"
  members = [
    "serviceAccount:${google_service_account.go_deployment_account.email}"
  ]
}

resource "google_project_iam_binding" "service_account_role_owner" {
  project = var.go_app_project_id
  role    = "roles/owner"
  members = [
    "serviceAccount:${google_service_account.go_deployment_account.email}"
  ]
}

resource "google_sql_database_instance" "gcp_sql_postgres" {
  name             = var.gcp_pg_name
  database_version = var.gcp_pg_database_version
  region           = var.gcp_pg_region
  project      = var.go_app_project_id
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
        image = "gcr.io/${var.go_app_project_id}/go-run-service/${var.go_app_name}:latest"
      }
    }
  }
}
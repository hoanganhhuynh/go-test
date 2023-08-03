
resource "google_sql_database_instance" "gcp_sql_postgres" {
  name             = var.gcp_pg_name
  database_version = var.gcp_pg_database_version
  region           = var.gcp_pg_region
  deletion_protection=false

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
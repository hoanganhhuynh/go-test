

locals {
  gcp_pg_name_timestamp = "${var.gcp_pg_name}-${timestamp()}"
}

resource "google_sql_database_instance" "gcp_sql_postgres" {
  name             = var.gcp_pg_name_timestamp
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
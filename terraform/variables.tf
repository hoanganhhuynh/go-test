
####################################
#   POSTGRES variables definition  #
####################################

variable "gcp_pg_name" {
  type = string
  default = "pg-instance"
}

variable "gcp_pg_name_timestamp" {
  type = string
  default = ""
}

variable "gcp_pg_database_version" {
  type = string
  default = "POSTGRES_15"
}

variable "gcp_pg_region" {
  type = string
  default = "us-central1"
}

variable "gcp_pg_tier" {
  type = string
  default = "db-f1-micro"
}

variable "gcp_pg_db_flag_name" {
  type = string
  default = "cloudsql.logical_decoding"
}

variable "gcp_pg_db_flag_value" {
  type = string
  default = "on"
}

variable "db_name" {
  type = string
  default = "GoLearn"
}

variable "db_user" {
  type = string
  default = "postgres"
}

variable "db_password" {
  type = string
  default = "postgres"
}

####################################
# GCP provider variable definition #
####################################

variable "project" {
  type        = string
  description = "GCP Project ID"
}

variable "gcp_region" {
  type        = string
  description = "GCP region"
}

variable "gcp_credentials" {
  type = string
}
#!/bin/bash -e

gcloud sql databases create "${DB_NAME}" --instance=pg-instance
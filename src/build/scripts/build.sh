#!/bin/bash -e

#OUTPUT_DIR=".output/"
echo "${OUTPUT_DIR}"

# docker-compose run --rm sdk go mod tidy
docker-compose run --rm bash cp -r "./build/sqls/" "${OUTPUT_DIR}"


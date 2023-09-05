#!/bin/bash -e

for file in ./build/scripts/data/*; do
    psql -f "$file"
done
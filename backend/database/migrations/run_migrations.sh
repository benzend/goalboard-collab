#!/bin/bash

for absolutepath in /code/database/migrations/*.up.sql; do
  [ -e "$absolutepath" ] || continue
  filename=$(basename -- "$absolutepath")
  filename="${filename%.*}"
  . /code/database/migrations/run_migration.sh $filename
done
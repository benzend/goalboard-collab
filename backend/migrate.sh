#!/bin/bash
# Runs migrations on PostgreSQL

# Reminder: If you are editing this file, please contribute your changes back to the community!
# Contribute here: https://github.com/Schniz/migrate.sh/edit/master/migrate.sh

set -e

# Load environment variables from .env file and check essential variables
source /code/.env

if [ -z "$user" ] || [ -z "$password" ] || [ -z "$host" ] || [ -z "$dbname" ]; then
  echo "ERROR: Database configuration variables are missing in .env file."
  exit 1
fi

DATABASE_URL="postgresql://${user}:${password}@${host}/${dbname}"
MIGRATIONS_DIR="/code/database/migrations"

make_new_migration() {
  local desc=$(echo "$@" | sed -E "s/[[:space:]]+/_/g" | awk '{print tolower($0)}')
  local file_path="${MIGRATIONS_DIR}/$(get_timestamp)_${desc}.sql"
  mkdir -p "${MIGRATIONS_DIR}"
  touch "${file_path}"
  echo "  --> New migration created at ${file_path}"
}

verify_database_url() {
  if [ -z "$DATABASE_URL" ]; then
    echo "  --> ERROR: DATABASE_URL is not defined"
    exit 1
  fi
}

pending_migrations() {
  local up_migrations=$(psql "$DATABASE_URL" -t -c "select filename from migration" | sed "s@[[:space:]]@@g" | grep .)
  local all_migrations=$(cd "$MIGRATIONS_DIR" && ls *.sql)
  local migrations_to_run=("${all_migrations[@]}")

  for item in ${up_migrations[@]}; do
    migrations_to_run=("${migrations_to_run[@]/$item/}")
  done

  echo "${migrations_to_run[@]}"
}

run_migrations() {
  local pending_migrations=($(pending_migrations))
  if [ ${#pending_migrations[@]} -eq 0 ]; then
    echo "  --> Nothing to migrate!"
  else
    create_migration_table
    echo "PENDING MIGRATIONS:"
    echo "${pending_migrations[@]}"
    echo "==================="
    for migration in "${pending_migrations[@]}"; do
      local contents=$(<"${MIGRATIONS_DIR}/${migration}")
      local contents_with_migration_result="
        BEGIN;
        ${contents};
        INSERT INTO migration (filename) VALUES ('$(basename "${migration}")');
        COMMIT;
      "
      echo "  --> Running ${migration}"
      echo "${contents_with_migration_result}" | psql "$DATABASE_URL" > /dev/null
    done
  fi
  dump_schema
}

dump_schema() {
  local schema_file="/code/database/schema.sql"
  pg_dump --schema-only "$DATABASE_URL" > "${schema_file}.tmp"
  pg_dump --data-only -t migration "$DATABASE_URL" >> "${schema_file}.tmp"
  grep -Ev '^(--|)$' "${schema_file}.tmp" | grep -v '^\-\-\- ?\(.*\)$' > "$schema_file"
  rm "${schema_file}.tmp"
  echo "  --> Database schema saved to ${schema_file}"
}


create_migration_table() {
  psql "$DATABASE_URL" <<< "
  CREATE TABLE IF NOT EXISTS migration (
    id         SERIAL PRIMARY KEY,
    filename   VARCHAR (120) NOT NULL
  );
  "
}

reset_db() {
  psql "$DATABASE_URL" <<< "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
}

show_help() {
  echo "Usage:"
  echo "  $0 new <description> - Create a new migration file"
  echo "  $0 schema:dump       - Dump the database schema"
  echo "  $0 schema:load       - Load the database schema"
  echo "  $0 up                - Run pending migrations"
  echo "  $0 danger:reset      - Resets the database state"
  echo "  $0 help              - Show this help message"
}

get_timestamp() {
  date +%Y%m%d%H%M%S
}

main_migrate() {
  local action=${1:-"help"}
  shift
  case "$action" in
    new)
      make_new_migration "$@"
      ;;
    schema:dump)
      verify_database_url
      dump_schema
      ;;
    schema:load)
      verify_database_url
      psql "$DATABASE_URL" -f /code/database/schema.sql
      ;;
    danger:reset)
      verify_database_url
      reset_db
      ;;
    up)
      verify_database_url
      run_migrations
      ;;
    *)
      show_help
      exit 1
  esac
}

main_migrate "$@"

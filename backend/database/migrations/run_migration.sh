#!/bin/bash
source /code/.env

psql postgresql://$user:$password@$host/$dbname -a -f /code/database/migrations/"$@".sql
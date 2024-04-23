#!/bin/bash

# PostgreSQL connection parameters
HOST="db"
PORT="5432"
DATABASE="postgres"
USERNAME="username"
PASSWORD="password"

# Directory containing SQL files
SQL_DIR="path_to_your_sql_files_directory"

# Maximum number of files or directories to process
MAX_COUNT=50

# Connect to PostgreSQL and execute SQL files
COUNT=0
for SQL_FILE in $(find "$SQL_DIR" -type f -name "*.sql" -not -path "*/vendor/*" | head -n $MAX_COUNT); do
    echo "Executing SQL file: $SQL_FILE"
    psql -h "$HOST" -p "$PORT" -d "$DATABASE" -U "$USERNAME" -W "$PASSWORD" -f "$SQL_FILE"
    if [ $? -eq 0 ]; then
        echo "SQL file executed successfully"
    else
        echo "Error executing SQL file: $SQL_FILE"
        exit 1
    fi
    COUNT=$((COUNT + 1))
done

echo "Processed $COUNT SQL files"
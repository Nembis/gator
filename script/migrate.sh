#!/bin/bash

# Check if an argument was provided
if [ -z "$1" ]; then
    echo "Usage: $0 <up|down>"
    exit 1
fi

# Check for .env file and load it
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Navigate to directory and run command
cd sql/schema || { echo "Directory sql/schema not found"; exit 1; }
goose postgres "$DB_URL" "$1"


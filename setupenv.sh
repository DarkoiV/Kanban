#!/bin/sh

default_db_host=psdb

# -b means build mode
if ! test "$1" = "b"; then
	test -d "build/latest" && cd "build/latest" 
	default_db_host=localhost
fi

echo Creating new .env

# Create .env
touch .env
echo '# ENV used by kanban app' > .env

echo
read -p 'Application port[9000]: ' APP_PORT
echo "APP_PORT=${APP_PORT:-9000}" >> .env

echo '\n# Database' >> .env

echo
read -p 'Database name[kanbanDB]: ' DB_NAME
echo "DB_NAME=${DB_NAME:-kanbanDB}" >> .env

read -p "Database host[$default_db_host]: " DB_HOST
echo "DB_HOST=${DB_HOST:-$default_db_host}" >> .env

read -p 'Database port[5432]: ' DB_PORT
echo "DB_PORT=${DB_PORT:-5432}" >> .env

read -p 'Database user[kanban_client]: ' DB_USER
echo "DB_USER=${DB_USER:-kanban_client}" >> .env

read -p 'Database pswd[kanban4321]: ' DB_PASS
echo "DB_PASS=${DB_PASS:-kanban4321}" >> .env

echo
echo "----- FILE -- .env -- FILE -----"
cat .env
echo "----- FILE -- .env -- FILE -----"
echo
echo Make sure env are correct!

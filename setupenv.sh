#!/bin/sh

# First argument is target directory
if ! test $1 = ""; then
	mkdir -p "$1" && cd $1
fi

echo "Creating new .env"

touch .env
echo "# ENV used by kanban app" > .env

read -p "Use postgres[y]?: " postgres
if test "$postgres" = "y"; then
	echo 'DIALECT="postgres"' >> .env

	read -p "Database name: " DB_NAME
	echo "DB_NAME=\"$DB_NAME\"" >> .env

	read -p "Database host: " DB_HOST
	echo "DB_HOST=\"$DB_HOST\"" >> .env

	read -p "Database port: " DB_PORT
	echo "DB_PORT=\"$DB_PORT\"" >> .env

	read -p "Database user: " DB_USER
	echo "DB_USER=\"$DB_USER\"" >> .env

	read -p "Database pswd: " DB_PASS
	echo "DB_PASS=\"$DB_PASS\"" >> .env
else
	echo "Using sqlite"
	echo 'DIALECT="sqlite"' >> .env
fi

echo
echo "----- FILE -- .env -- FILE -----"
cat .env
echo "----- FILE -- .env -- FILE -----"
echo
echo "Make sure env are correct!"

#!/bin/sh

# Create build folder
test -d "build" || mkdir build

# Create new folder for current build
FOLDER_NAME=$(date | awk '{print $2 "-" $3 "T" $5}' | sed -r 's/:/./g')
mkdir build/$FOLDER_NAME

# Build go binary
(cd backend; go get -d -v && go build -o ../build/$FOLDER_NAME/server)

# Build Vue SPA
(cd frontend; npm run build && test -d dist && mv dist ../build/$FOLDER_NAME)

# Link latest build
(cd build; test -d latest && unlink latest; ln -s $FOLDER_NAME latest)

# Ask for .env
(cd build/$FOLDER_NAME
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

echo "----- FILE -- .env -- FILE -----"
cat .env
echo "----- FILE -- .env -- FILE -----"
)

# Create image
read -p "Create docker image[y]?" dockerImage
if test "$dockerImage" = "y"; then
	docker build -t darkoiv/kanban .
fi

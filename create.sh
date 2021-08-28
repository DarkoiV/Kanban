#!/bin/sh

# Create build folder
test -d "build" || mkdir build

# Create new folder for current build
FOLDER_NAME=$(date | awk '{print $2 "-" $3 "T" $5}' | sed -r 's/:/./g')
mkdir build/$FOLDER_NAME

# Build go binary
(cd backend; go get -d -v && go build -o ../build/$FOLDER_NAME/server)

# Build Vue SPA
(cd frontend; test -d node_modules || npm install; npm run build && mv dist ../build/$FOLDER_NAME)

# Link latest build
(cd build; test -d latest && unlink latest; ln -s $FOLDER_NAME latest)

# Ask for .env
./setupenv.sh "build/$FOLDER_NAME"

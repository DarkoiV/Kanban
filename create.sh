#!/bin/sh

# Create build folder
test -d "build" || mkdir build

# Create new folder for current build
FOLDER_NAME=$(date | awk '{print $2 "-" $3 "T" $5}' | sed -r 's/:/./g')
mkdir build/$FOLDER_NAME

# Build go binary
(cd backend; go build -o ../build/$FOLDER_NAME/server)

# Build Vue SPA
(cd frontend; npm run build && test -d dist && mv dist ../build/$FOLDER_NAME)

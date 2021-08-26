# Kanban
Basic kanban board web-app.  

***

## Build binary
Make sure you have installed `go 1.16+` and `npm 7.20+` 
```bash
# Clone
git clone https://github.com/DarkoiV/Kanban.git
cd Kanban

# Build binary, Vue page and create .env file
./create.sh

# Run server (runs on port 9000)
cd build/latest
./server
```
Each build will create new folder!

***

## Build docker container
```bash
# Clone
git clone https://github.com/DarkoiV/Kanban.git
cd Kanban

# Create .env
./setupenv.sh

# Run docker build
docker build -t darkoiv/kanban .

# Run docker image (binds port 9000)
docker run -p "9000:9000" darkoiv/kanban
```

# Kanban
Basic kanban board web-app.  

***
Creating new board and lists
![Creating-Board](https://user-images.githubusercontent.com/25897608/133525412-5eeb5366-8750-48b6-b870-bcd46bf7101d.gif)

***
## Build binary
Make sure you have installed `go 1.16+` and `npm 7.20+` 
```bash
# Clone
git clone https://github.com/DarkoiV/Kanban.git
cd Kanban

# Build binary, Vue page and create .env file
./create.sh

# Run server (by default runs on port 9000)
cd build/latest
./server
```
Each build will create new folder!

***

## Build docker
```bash
# Clone
git clone https://github.com/DarkoiV/Kanban.git
cd Kanban

# Create .env
./setupenv.sh

# Run docker-compose
docker-compose up

```

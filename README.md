# Kanban
Basic kanban board web-app.  

***
Creating new board and lists
![Creating-Board](https://user-images.githubusercontent.com/25897608/133525412-5eeb5366-8750-48b6-b870-bcd46bf7101d.gif)

Moving, editing and creating tasks.
![Task](https://user-images.githubusercontent.com/25897608/133529233-5e719c21-60b1-452d-b167-99e86a01c41b.gif)

Deleting tasks and lists
![Deleting](https://user-images.githubusercontent.com/25897608/133529264-4b2a7d90-e09a-4ed6-bc82-fb368e429901.gif)


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

# GoFilez
GGWP fastest and most functional file server in sweden :)

## Idea
Sync files on accross all devices. Chose what directories to sync. Interaction with files via web app, gui app or file browser. 

## TODO

* [ ] [Wails desktop app](https://github.com/wailsapp/wails)
* [ ] [Gin website app](https://github.com/gin-gonic/gin)

## Setup

Prerequisite docker

### Build image

` docker build --tag  gofilez:latest .`

### Start conatiner

 `docker compose up` 

### Enter command line of running container
`docker-compose gofilez sh`
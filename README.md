# example-golang-http-server

This repo is example code for explaining Docker for Kata.ai internal team

## Step-by-step
1. Clone this repo
2. Run `docker build -t example-docker .`
3. Run `docker image ls`
4. Run `docker run -d -p 8080:8080 example-docker`
5. Run `docker container ls`
6. Run `curl http://localhost:8080`
7. Run `curl http://localhost:8080/whoami`
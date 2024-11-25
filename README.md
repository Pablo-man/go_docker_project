# Hello world whit go
## Results
<p align="center">
    <img src="./public/img/result.png" alt="Hello from go">
</p>

Web application that deploys a web server and builds a simple API with a route that displays the message **hello world with go!**

## :open_book: How to use
### Pre-requisites
* go language
* code editor
---

* Clone the repository
    ```
    git clone https://github.com/Pablo-man/go_docker_project.git
    ```
* Open the program with the code editor of your choice
* Open a terminal that points to the root of the project
* Run the command:

    `go run main.go`
* Visit your `localhost:8080` to view the results

    > [!TIP]
    > By default the application will be deployed on port `8080`, if it is necessary to change it to another one you must modify it from the `main.go` file and in the `http.ListenAndServe(":<PORTNUMBER>", nil)` section place the desired port

## :rocket: How to run with docker
### Pre-requisites
* Docker - DockerDesktop installed
* DockerHub account
---
Visit the following link to learn about the process of generating the project image

:whale2: [GO](https://hub.docker.com/repository/docker/pamendeza/go_docker_project "Docker steps")

## :tennis: PAAS Deploy(Render) without Docker
Type of deployment was:

![Render Service](./public/img/type.png "Service")

Configurations:
* Build and start commmands to compile and execute the web application on render.

![Render Config](./public/img/conf.png "Configuration")

![Render Config](./public/img/conf1.png "Configuration")


State:

![Render Service](./public/img/renderDeploy1.png "Service")

![Render Service](./public/img/renderDeploy.png "Service")


:cake:[Hello World](https://go-docker-project.onrender.com/ "click for visit")
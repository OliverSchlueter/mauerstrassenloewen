---
icon: apps
order: 200
---

# Architecture

We have decided to adopt a microservices architecture for our application. 
Inter-service communication will primarily use HTTP/REST or CloudEvents/NATS, while asynchronous messaging will be explored for cases that benefit from decoupling and eventual consistency. 
Services will manage their own databases or schemas to ensure loose coupling and domain ownership. 
To support this architecture, we will containerize each service using Docker and orchestrate them locally using Docker Compose, with the option to explore Kubernetes if time permits. 
Additionally, we will implement basic observability features such as health checks, centralized logging, and monitoring, and we will set up CI/CD pipelines using GitHub Actions to support automated testing and deployment.

To get the reasons for each decision, please read the [Architecture Decision Records](adr/adr.md).

## Overview

![Architecture Overview](static/architecture_overview.png)

*You can find the exported excalidraw file [here](static/architecture.excalidraw).*

## Services

Now let's take a look at the services we have in our architecture.

### Reverse proxy

We decided to write our own reverse proxy (RP) for learning purposes and to have a better control over the routing and the features we want to implement.
The RP is responsible for routing the requests to the correct service based on the URL path. It might also have load balancing capabilities in the future.

The RP is the entry point for all the incoming requests. It will route each request to the correct service based on the URL path.

We configured the following rules for the RP:

- All requests starting with `/msl` will be routed to the `backend` service
- All requests starting with `/simulation` will be routed to the `simulation` service
- All other requests will be routed to the `frontend` service

Examples:

- `http://localhost:8080/msl/api/v1/chatbot` will be redirected to `http://backend:8082/api/v1/chatbot`
- `http://localhost:8080/simulation/api/v1/simulation` will be redirected to `http://simulation:8083/api/v1/simulation`
- `http://localhost:8080/` will be redirected to `http://frontend:8081/`
- `http://localhost:8080/login` will be redirected to `http://frontend:8081/login`

If a service is not available, the RP will return a `502` (Bad Gateway) status code.

### Frontend

The frontend service is responsible for serving the web application. It's basically a file server that serves the static files of the web application (HTML, CSS, JS and assets). 
Before compiling the frontend service, the actual frontend needs to be built and the output needs to be copied to a directory in the frontend service.

The frontend service is also responsible for serving the documentation and OpenAPI specifications.

### Backend

The backend service is responsible for serving the API and the business logic of the application. Here we have the core functionality of the application.
For example, user management, authentication and authorization are handled in the backend service.

### AI Worker

Computing AI task is a resource-intensive process. We decided to separate the AI tasks from the main application and run them in a separate service. 
The AI worker service has no HTTP API, it only listens to NATS messages and processes them.

The AI worker has connections to several AI providers (ollama and OpenAI) and supports RAG using a vector database.

The AI worker is responsible for processing the AI tasks and returning the results to the main application.

### Simulation

The simulation service is responsible for running the simulations. It has an HTTP API to manage simulations.
 
!!!warning Sunset notice
This service has been sunset to focus on the core features of the application.
!!!

### Monitoring

The monitoring service is responsible for collecting metrics and logs from all the services. It uses Prometheus and Grafana to collect and visualize the metrics.

It listens to all NATS messages and collects the metrics from all the services.
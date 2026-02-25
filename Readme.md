# URL Shortener API

A simple RESTful API built to shorten long URLs, retrieve original URLs, update or delete short URLs, and track access statistics.

This project demonstrates CRUD operations, RESTful design principles, and access tracking using a database-backed service.

# Features

1. Create short URLs
2. Retrieve original URLs
3. Delete short URLs
4. Track access count
5. RESTful API design

# Tech Stack

1. <b>Backend:</b> Go
2. <b>Database:</b> MongoDB
3. <b>Containerization:</b> Docker
4. <b>Orchestration:</b> Kubernetes

# Running Locally

1. Clone the repository

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

2. Start MongoDB (docker):

```bash
docker run -d --name mongodb -p 27017:27017 mongo:7
```

3. Export ENV vars

```bash
export MONGO_URI="mongodb://localhost:27017"
export MONGO_DB="db_name"
export PORT=5000
```

4. Run application

```bash
go run cmd/main.go
```

Server running at `localhost:5000`

# 🐋Running with docker

Run: `docker compose up --build`

# Running with Kubernetes

```bash
kubectl apply -f k8s/
```

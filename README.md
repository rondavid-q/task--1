# Microservices Example: Node.js, Python, and Go

This repository contains a simple example of microservices architecture implemented using Node.js, Python, and Go. The services communicate with each other to perform a data processing task.

## Services

### 1. Node.js Service

- **Description:** Accepts requests and forwards them to the Python and Go services.
- **Endpoint:** `/process-data`
- **Dependencies:** Express.js, Axios

#### How to Run:

```bash
cd nodejs-service
npm install
node input.js


2. Python Service

- **Description:** Performs a simple data processing task (sorting a list of numbers).
- **Endpoint:** `/process-data`
- **Dependencies:** Flask

```bash
cd python-service
pip install -r requirements.txt
python data.py



3. Go Service

- **Description:** Logs the requests received and responses sent, and returns a confirmation back to the Node.js service.
- **Endpoint:** `/log-request`


```bash
cd go-service
go run process.go


Note: For all the aboce you need Node.js, python and go to be installed in system.


ORCHESTRATION USING DOCKER AND DOCKER COMPOSE

From the root folder
```bash

docker-compose build
docker-compose up -d

TO SEE THE LOGS
docker-compose logs
or
docker ps

Get the container ID of each service
docker logs -f <container-ID>


TESTING THE FLOW

From the Terminal

NOTE: USING Curl - NEED curl installed in your system
If you are testing locally

curl -X POST -H "Content-Type: application/json" -d '{"numbers": [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]}' http://127.0.0.1:8000/process-data     

If you are testing on a remote server
Note: Please open port 8000-8002 in your remote server

curl -X POST -H "Content-Type: application/json" -d '{"numbers": [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]}' http://<remote-ip>:8000/process-data 

OUTPUT:

{"pythonResponse":[1,1,2,3,3,4,5,5,5,6,9],"goResponse":{"message":"Request logged and processed successfully at 2023-12-26 15:52:52"}}
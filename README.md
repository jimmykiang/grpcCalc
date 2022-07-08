// start docker server.
docker build -t grpcserver . -f DockerfileServer
docker run -p 50051:50051 grpcserver


docker build -t grpcclient . -f DockerfileClient

// run client connection to server without docker.
./bin/calculator/client

docker-compose up --exit-code-from grpcclient


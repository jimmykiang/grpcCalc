// start docker server.
docker build -t grpcserver . -f DockerfileServer
docker run -p 50051:50051 grpcserver


docker build -t grpcclient . -f DockerfileClient

// run client connection to server without docker.
./bin/calculator/client

docker-compose up --exit-code-from grpcclient



docker build -t jugon666/grpcserver:latest . -f DockerfileServer
docker push jugon666/grpcserver:latest

docker build -t jugon666/grpcclient:latest . -f DockerfileClient
docker push jugon666/grpcclient:latest


// first server must be running.
apply -f serverDeploy.yaml

// apply clientDeploy after server is in running state.
apply -f clientDeploy.yaml

// check result of grpc roundTrip client-server-client.
kubectl logs grpcclient-5466f78b44-q96sr


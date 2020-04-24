mkdir -p ../build
cd ../cmd/client_grpc
go build .
mv client_grpc.exe ../../build
cd ../../build
./client_grpc.exe -server=localhost:9090

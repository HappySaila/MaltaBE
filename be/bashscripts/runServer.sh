rm -rf ../build
mkdir -p ../build
cd ../cmd/server
go build .
mv ./server.exe ../../build
cd ../../build
./server.exe -grpc-port=9090 -db-host=bgtusc51sgcqzmjlcisq-mysql.services.clever-cloud.com:3306 -db-user=uj8arcdthnwjt1sj -db-password=sLvaummcX1e7viu4kiLo -db-schema=bgtusc51sgcqzmjlcisq

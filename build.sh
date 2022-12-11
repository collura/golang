go build application.go
mv application _devops/
cd _devops
docker build -t golang-1 .
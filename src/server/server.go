package server

import (
	"log"
	"net"
	"petProj/src/attendanceProtos"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)
func StartServer(db *gorm.DB){
l, err := net.Listen("tcp", ":8089")
if err != nil {
	log.Fatalf("failed to listen: %v", err)
}
serverRegistrar := grpc.NewServer()

employeeService := NewEmployeeServer(db)

protos.RegisterEmployeeServiceServer(serverRegistrar, employeeService)

err = serverRegistrar.Serve(l)
if err != nil {
	log.Fatalf("failed to serve: %v", err)
}
}
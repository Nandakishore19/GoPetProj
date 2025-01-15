package client

// TODO: Create a client for the server

import (
	"log"
	protos "petProj/src/attendanceProtos"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn *grpc.ClientConn
	Client protos.EmployeeServiceClient
}

func NewGRPCClient(address string)(*GRPCClient, error) {
	conn, err := grpc.NewClient(address)
	if err != nil {
		return nil, err
	}
	employeeServiceClient := protos.NewEmployeeServiceClient(conn)
	return &GRPCClient{
		conn: conn,
		Client: employeeServiceClient,
	},err
}

func (c *GRPCClient) Close(){
	if err := c.conn.Close(); err != nil {
		log.Printf("Failed to close gRPC connection: %v", err)
	}
}
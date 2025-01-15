package client

import (
	"context"
	"log"
	"petProj/src/attendanceProtos"
	"time"

)

func CallAddEmployees(grpcClient *GRPCClient) error {
	ctx,cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()

	addEmpRequest := &protos.AddEmployeeRequest{
		EmployeeId: 1,
		Email: "",
		MacAddress: "",
		Name: "",
	}
	response, err := grpcClient.Client.AddEmployee(ctx, addEmpRequest)
	if err != nil {
		return err
	}
	log.Printf("Response from Add Employees Endpoint: %v",response)
	return nil
}

func CallGetEmployees(grpcClient *GRPCClient, mac_address []string) error {
	ctx, cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()

	getEmpRequest := &protos.GetEmployeeRequest{
		MacAddress: mac_address,
	}
	response, err := grpcClient.Client.GetEmployees(ctx,getEmpRequest)
	if err != nil {
		return err
	}
	log.Printf("Response from Get Employees Endpoint: %v",response)
	return nil
}

func CallLogAttendanceEndpoint(grpcClient *GRPCClient) error {
	ctx, cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()

	logAttendanceRequest := &protos.EmployeeAttendanceRequest{
		EmployeeId: 1,
	}
	response, err := grpcClient.Client.LogEmployeeAttendance(ctx,logAttendanceRequest)
	if err != nil {
		return err
	}
	log.Printf("Response from Log Employee Attendance Endpoint: %v", response)
	return nil
}
package server

import (
	"context"
	"fmt"
	"petProj/src/attendanceProtos"
	"time"
	// "petProj/src/models"
	"petProj/src/repositories"
	"gorm.io/gorm"

	// "petProj/src/models"
)


type myEmployeeServer struct {
	protos.UnimplementedEmployeeServiceServer
	attendanceRepo *repositories.AttendanceRepository
	employeeRepo *repositories.EmployeeRepository
}

func NewEmployeeServer(db *gorm.DB) *myEmployeeServer {
	return &myEmployeeServer{
		attendanceRepo: repositories.NewAttendanceRepository(db),
		employeeRepo: repositories.NewEmployeeRepository(db),
	}
}

func (s *myEmployeeServer) GetEmployees(ctx context.Context,req *protos.GetEmployeeRequest)(*protos.GetEmployeeResponse, error){
	// emps:= []*protos.Employee{}
	dbEmps,err := s.employeeRepo.GetEmployees(req.MacAddress);
	emps := make([]*protos.Employee, len(dbEmps))
	if err != nil {
		return nil, err
	}
	for i, dbEmp := range dbEmps {
		emps[i] = &protos.Employee{
			EmployeeId:  int32(dbEmp.EmployeeID),
			Name:       dbEmp.Name,
			MacAddress: dbEmp.MacAddress,
		}
	}
	if err != nil {
		fmt.Println("Couldn't retrieve List of Employees")
	}

	return &protos.GetEmployeeResponse{
		Employees: emps,
	},nil
}

func (s *myEmployeeServer) AddEmployee(ctx context.Context,req *protos.AddEmployeeRequest)(*protos.AddEmployeeResponse, error){

	err := s.employeeRepo.AddEmployee(int(req.EmployeeId), req.Email, req.MacAddress, req.Name)
	if err != nil {
		fmt.Printf("Failed to add employee: %v\n", err)
		return &protos.AddEmployeeResponse{
			Success: false,
			Message: "failed to add employee",
		}, err
	}

	return &protos.AddEmployeeResponse{
		Success: true,
		Message: "success",
	},nil

}

func (s *myEmployeeServer) LogEmployeeAttendance(ctx context.Context,req *protos.EmployeeAttendanceRequest)(*protos.EmployeeAttendanceResponse, error){
	currentTime := time.Now()

	date := currentTime.Format("2002-19-11")
	err := s.attendanceRepo.LogEmployeeAttendance(int(req.EmployeeId),date)
	if err != nil {
		fmt.Printf("Failed to Log Employee Attendance: %v\n", err)
		return &protos.EmployeeAttendanceResponse{
			Success: false,
			Message: "Failed to log Employee",
		},nil
	}
	return &protos.EmployeeAttendanceResponse{
		Success: true,
		Message: "Success",
	},nil
}
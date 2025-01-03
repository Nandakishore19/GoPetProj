package main

import (
	// "fmt"
	"log"
	"petProj/src/repositories"

	"google.golang.org/grpc"
	// "time"
	"petProj/src/attendanceProtos"
	"petProj/src/server"
	"net"
)


func main() {
	db,err := repositories.NewDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	l, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverRegistrar := grpc.NewServer()

	employeeService := server.NewEmployeeServer(db)

	protos.RegisterEmployeeServiceServer(serverRegistrar, employeeService)

	err = serverRegistrar.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// empRepo := repositories.NewEmployeeRepository(db)
	// attendanceRepo := repositories.NewAttendanceRepository(db)

	// employeesToAdd := []struct {
	// 	id        int
	// 	email     string
	// 	macAddress string
	// 	name      string
	// }{
	// 	{1, "nanda@kishore.com", "00:00:00:00:00:11", "Nanda"},
	// 	{2, "kishore@kishore.com", "00:00:00:01:00:00", "Kishore"},
	// 	{3, "kishore@nanda.com", "00:00:00:02:00:00", "Adnan"},
	// }

	// for _, e := range employeesToAdd {
	// 	if err = empRepo.AddEmployee(e.id, e.email, e.macAddress, e.name); err != nil {
	// 		log.Fatalf("failed to add employee %s: %v", e.name, err)
	// 	}
	// }
	// if err != nil {
	// 	log.Fatalf("failed to add employee: %v", err)
	// }

	// log.Println("Employee added successfully")
	// fmt.Println("Employee added successfully")

	// employees, err := empRepo.GetEmployees([]string{"00:00:00:00:00:11", "00:00:00:01:00:00"})
	// if err != nil {
	// 	log.Fatalf("failed to get employees: %v", err)
	// }
	// for _, employee := range employees {
	// 	log.Printf("Employee ID: %d, Email: %s, MAC Address: %s, Name: %s", employee.EmployeeID, employee.Email, employee.MacAddress, employee.Name)
	// }

	// EmployeeId := 1
	// date := time.Now().Format("2006-01-02")
	// err = attendanceRepo.LogEmployeeAttendance(EmployeeId, date)
	// if err != nil {
	// 	log.Fatalf("failed to log employee attendance: %v", err)
	// }
	// log.Println("Employee attendance logged successfully for employee ID:", EmployeeId)

}
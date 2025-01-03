generate_grpc_code:
	protoc \
	--go_out=attendanceProtos \
	--go_opt=paths=source_relative \
	--go-grpc_out=attendanceProtos \
	--go-grpc_opt=paths=source_relative \
	attendance_management.proto
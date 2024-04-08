package server

import (
	"context"
	"fidelis.com/communication/dbtools"
	"fidelis.com/communication/model"
	"fidelis.com/communication/protocol"
)

type GrpcController struct {
	dbConnection *dbtools.DBInitializer
	protocol.UnimplementedStudentServiceServer
}

func GrpcServerInitializer(dataSourceName string) (*GrpcController, error) {
	dbtInit, err := dbtools.Connect(dataSourceName)
	if err != nil {
		return &GrpcController{}, nil
	}
	return &GrpcController{
		dbConnection: dbtInit,
	}, nil
}

func (controller *GrpcController) GetStudentByID(context context.Context, id *protocol.GetStudentByIDRequest) (*protocol.Student, error) {
	newStudent := controller.dbConnection.GetStudentById(id.GetID())

	return convertToGrpcStudent(newStudent), nil
}

func (controller *GrpcController) GetStudentBYName(name *protocol.GetStudentBYNameRequest, grpcStudents protocol.StudentService_GetStudentBYNameServer) error {
	students := controller.dbConnection.GetStudentsByName(name.GetName())

	for _, student := range students {
		err := grpcStudents.Send(convertToGrpcStudent(student))
		if err != nil {
			return err
		}
	}

	return nil
}

func convertToGrpcStudent(mStudend model.Student) *protocol.Student {

	return &protocol.Student{
		Name: mStudend.Name,
		Age:  mStudend.Age,
		ID:   mStudend.ID,
	}
}

package main

import (
	"context"
	"fidelis.com/communication/configuration"
	"fidelis.com/communication/protocol"
	"fidelis.com/communication/server"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	config, err := configuration.New()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("Failed to load dataSourceName")
	}

	var option string
	flag.StringVar(&option, "admin", "server", "communication between server and client")
	flag.Parse()

	switch option {
	case "server":
		runGrpcServer(config.DataSourceName)
	case "client":
		runGrpClient()
	}
}

func runGrpClient() {
	connection, err := grpc.NewClient("127.0.0.1:3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer connection.Close()

	client := protocol.NewStudentServiceClient(connection)

	fmt.Println("\n\nWelcome, Please enter number to select your requests")
	fmt.Println("1:Select All Students based Name")
	fmt.Println("2:Select a Student Based ID Number")

	fmt.Println("\nPlease enter your Number:")
	var input string
	fmt.Scanln(&input)

	if strings.EqualFold(input, "1") {
		value := ""
		fmt.Println("\nEnter your name")
		fmt.Scanln(&value)
		students, err := client.GetStudentBYName(context.Background(), &protocol.GetStudentBYNameRequest{
			Name: value,
		})

		if err != nil {
			log.Fatalln(err.Error())
		}

		for {
			student, err := students.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(student)
		}
		fmt.Println("")

		return
	} else if strings.EqualFold(input, "2") {
		fmt.Println("\nEnter your ID:")
		fmt.Scanln(&input)

		id, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		student, err := client.GetStudentByID(context.Background(), &protocol.GetStudentByIDRequest{
			ID: int64(id),
		})

		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(student)
		fmt.Println("")
	} else {
		fmt.Println("You choose an invalid command")
	}
}

func runGrpcServer(dataSourceName string) {
	fmt.Println("Starting Server ...........")
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Server is listenening ...")

	var options []grpc.ServerOption

	newServer := grpc.NewServer(options...)

	studentServer, err := server.GrpcServerInitializer(dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	protocol.RegisterStudentServiceServer(newServer, studentServer)

	err = newServer.Serve(listener)

	if err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"fmt"
	"os"
	"GOCHAT/server"
	"GOCHAT/client"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("Usage: go run main.go [server|client]")
	} 

	switch os.Args[1]{
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Invalid option. Use 'server' or 'client'.")
	}
	
}

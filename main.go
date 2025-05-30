package main

import (
	"fmt"
	"os"
	"BetterGOChat/server"
	"BetterGOChat/client"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("Usage: go run main.go [server|client]")
	} 

	switch os.Args[1]{
	case "server":
		go server.StartWebSocketServer()
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Invalid option. Use 'server' or 'client'.")
	}
	
}

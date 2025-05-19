package server

import(
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"os"
)

type Client struct{
	conn net.Conn
	username string
}

var(
	clients = make(map[net.Conn]Client) //map of all clients connected to server
	broadcast = make(chan string) //Creates a channel that allows the server to send and recieve data from clients
	mutex = &sync.Mutex{} //Syncronizes the mutex data so users dont conflict
)

func StartServer(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	ipaddy := getThisIP()
	fmt.Printf("Server started on %s:8080", ipaddy)

	//^^ Server begins listening for tcp connections on port 8080, sets listener to close when server is shut down

	go handleBroadcast() //opens a new thread to handle the server broadcasts

	for{
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Connection error: ", err)
			continue
		}
		fmt.Println("Client connected: ", conn.RemoteAddr())
		go handleConnection(conn)
	}
	//^^ continuously loops to handle accepting new connections into the server - once it 
	// catches a connection it opens a new thread to handle that specific user in the server.
}

func handleConnection(conn net.Conn){
	defer conn.Close() //closes the connection on function completion

	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Failed to read username: ", err)
		return
	}
	username = strings.TrimSpace(username)
	//^^ User inputs a username into the command prompt and it is registered with the server

	mutex.Lock()
	clients[conn] = Client{conn: conn, username: username}
	mutex.Unlock() 

	//^^ Locks the Client struct so only this thread can access it, adds a new user to it, then unlocks
	//so other threads can access it again - prevents deadlock

	broadcast <- fmt.Sprintf("%s has joined the chat! \r\n", username)

	for{
		message, err := reader.ReadString('\n')
		if err != nil{
			fmt.Printf("%s has disconnected. \r\n", username)
			broadcast <- fmt.Sprintf("%s has left the chat. \r\n", username)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			return
		}
		broadcast <- fmt.Sprintf("%s: %s", username, message)
	}
	//^^ For loop to continuously post user messages to the server, as well as broadcast if a user has
	//disconnected from the server - locks the Clients struct before deleting user then unlocks to avoid
	//deadlock in struct
}

func handleBroadcast(){
	for{
		msg := <- broadcast
		mutex.Lock()
		for _, client := range clients{
			_, err := client.conn.Write([]byte(msg))
			if err != nil{
				fmt.Println("Error sending message to client: ", err)
			}
		}
		wsBroadcast <- msg
		mutex.Unlock()
	}
	//^^Loops over and continuously updates clients with the current messages
}

func getThisIP()string{
	hostname, err := os.Hostname()
	if err != nil{
		fmt.Println("error getting hostname: ", err)
		return "COULD NOT FIND HOSTNAME"
	}
	addresses, err := net.LookupIP(hostname)
	if err != nil{
		fmt.Println("Error getting IP addresses: ", err)
		return "COULD NOT FIND IP"
	}
	for _, addr := range addresses{
		if ipv4 := addr.To4(); ipv4 != nil{
			return ipv4.String()
		}
	}
	return "ERROR GETTING IP"
}
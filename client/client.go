package client

import(
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"path/filepath"
	"github.com/gorilla/websocket"
)

var(
	wsBroadcast = make(chan string)
	wsConn *websocket.Conn
)

func StartClient(){
	var ip string
	fmt.Print("Enter server IP: ")
	fmt.Scanln(&ip)

	conn, err := net.Dial("tcp", ip)
	if err != nil{
		fmt.Println("Connection failed: ", err)
		return
	}
	fmt.Println("Connected to chatroom!")

	fmt.Print("Enter your username: ")
	username := bufio.NewReader(os.Stdin)
	name, _ := username.ReadString('\n')
	conn.Write([]byte(name))

	go func(){
		reader := bufio.NewReader(conn)
		for{
			message, err := reader.ReadString('\n')
			if err != nil{
				fmt.Println("Lost connection to server: ", err)
				return
			}
			wsBroadcast <- message
		}
	}()
	
	go connectWebSocket()
	go serveGui()
	openBrowser("http://localhost:3000") //serves the GUI
	select{}
}

func serveGui(){
	workDir, _ := os.Getwd()
	uiPath := filepath.Join(workDir, "ui")
	http.Handle("/", http.FileServer(http.Dir(uiPath)))
	fmt.Println("Serving GUI on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("Failed to start HTTP Server")
	}
}

func openBrowser(url string){
	var err error
	switch runtime.GOOS{
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	
	case "darwin":
		err = exec.Command("open", url).Start()
	
	} 
	if err != nil{
		fmt.Println("Failed to open browser: ", err)
	}
}

func connectWebSocket(){
	var err error
	wsConn, _, err = websocket.DefaultDialer.Dial("ws://localhost:8081/ws", nil)
	if err != nil{
		fmt.Println("WebSocket connection failed: ", err)
		return
	}

	go func(){
		for{
			_, message, err := wsConn.ReadMessage()
			if err != nil{
				fmt.Println("WebSocket error: ", err)
				return
			}
			fmt.Println(string(message))
		}
	}()

	go func(){
		for msg := range wsBroadcast{
			err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil{
				fmt.Println("Error sending to WebSocket: ", err)
				return
			}
		}
	}()
}
package client

import(
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
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


	openBrowser("http://localhost:3000") //serves the GUI

	go func(){
		for{
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil{
				fmt.Println("Disconnected from sever.")
				return
			}
			fmt.Print(message)
		}
	}()

	for{
		input := bufio.NewReader(os.Stdin)
		msg, _ := input.ReadString('\n')
		conn.Write([]byte(msg))
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

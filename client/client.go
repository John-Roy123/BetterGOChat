package client

import(
	"bufio"
	"fmt"
	"net"
	"net/http"
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

	go serveGui()
	openBrowser("http://localhost:3000") //serves the GUI
}

func serveGui(){
	http.Handle("/", http.FileServer(http.Dir("./ui")))
	fmt.Println("Serving GUI on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("Failed to start HTTP server", err)
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

package server

import(
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var(
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}
	wsClients = make(map[*websocket.Conn]bool)
	wsMutex = &sync.Mutex{}
	wsBroadcast = make(chan string)
)

func StartWebSocketServer(){
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Websocket server started on :8081")
	go handleWebSocketBroadcast()
	http.ListenAndServe(":8081", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println("WebSocket upgrade failed", err)
		return
	}

	wsMutex.Lock()
	wsClients[conn] = true
	wsMutex.Unlock()
	fmt.Println("New WebSocket client connected")

	for{
		_, msg, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("WebSocket client disconnected")
			wsMutex.Lock()
			delete(wsClients, conn)
			wsMutex.Unlock()
			conn.Close()
			return
		}
		wsBroadcast <- string(msg)
	}
}

func handleWebSocketBroadcast(){
	for{
		message := <-wsBroadcast
		fmt.Println("Broadcasting message to WebSocket clients: ", message)

		wsMutex.Lock()

		for client := range wsClients{
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil{
				fmt.Println("Error sending message to client: ", err)
				client.Close()
				delete(wsClients, client)
			}
		}
		wsMutex.Unlock()
	}
}
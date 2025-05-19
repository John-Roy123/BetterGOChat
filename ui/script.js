let socket

window.onload = () => {
    socket = new WebSocket("ws://localhost:8081")

    socket.onmessage = (event) =>{
        const output = document.getElementById("output")
        output.innerHTML += event.data + "<br>"
        output.scrollTop = output.scrollHeight
    }

    socket.onopen = () =>{
        console.log("Connected to the websocket server.")
    }
    socket.onerror = (error) => {
        console.log("Websocket error: ", error)
    }
}

function sendMessage(){
    const message = document.getElementById("message").value
    if(message.trim()){
        socket.send(message)
        document.getElementById("message").value = ""
    }

}
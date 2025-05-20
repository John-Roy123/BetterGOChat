let socket
let username

window.onload = () => {
    const urlParams =  new URLSearchParams(window.location.search)
    const serverIP = urlParams.get("serverip")
    username = urlParams.get("user")
    console.log(serverIP)
    socket = new WebSocket(`ws://${serverIP}:8081/ws`)

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
    const message = '['+username+']: ' + document.getElementById("message").value
    if(message.trim()){
        socket.send(message)
        document.getElementById("message").value = ""
    }

}


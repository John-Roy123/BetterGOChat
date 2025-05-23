let socket
let username
const inputPrompt = document.getElementById("message")

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
    const message = '['+username+']: ' + inputPrompt.value
    if(message.trim()){
        socket.send(message)
        inputPrompt.value = ""
    }

}

document.querySelector('#message').addEventListener('keypress', function(e){
    if(e.key === 'Enter'){
        sendMessage()
    }
})
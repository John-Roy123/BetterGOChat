let socket
const inputPrompt = document.getElementById("message")
const urlParams =  new URLSearchParams(window.location.search)
const serverIP = urlParams.get("serverip")
const username = urlParams.get("user")

window.onload = () => {
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

function sendMessage() {
    const rawMessage = inputPrompt.value
    const message = '['+username+']: ' + rawMessage
    if(message.trim()) {
        socket.send(message);
        inputPrompt.value = "";
    }
}


document.querySelector('#message').addEventListener('keypress', function(e){
    if(e.key === 'Enter'){
        sendMessage()
    }
})
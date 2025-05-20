## To Run
Install the dependent libraries:
```
go get github.com/gorilla/websocket
```
Then from there cd into the project folder and run 
```
go run main.go {client|server}
```
depending on whether you want to host the server on your computer, write server, if you just want to connect to an existing server, write client.  From there, you will be prompted to enter the server IP, which will be in the server console log. Enter the server IP as such: xxx.xxx.xxx.xxx:8080 to connect to the server on port 8080  
After that, enter your username and you will join the server!

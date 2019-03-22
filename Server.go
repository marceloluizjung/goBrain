package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
import "os/exec"

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8091")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
	fmt.Print("Message Received:", string(message))
	
	//Command here
	var entrada string = string(message)
	var compare string = "exp"
	fmt.Print("Teste Received:", entrada)
	if (entrada == compare) {
	fmt.Print("Your comand ok")
	cmd := exec.Command(`explorer`, `/select,`, `C:\`)
    cmd.Run()
	}
    // sample process for string received
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}
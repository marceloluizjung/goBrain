package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8091")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {

  // read in input from stdin
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Text to send: ")
  text, _ := reader.ReadString('\n')
  // send to socket
  fmt.Fprintf(conn, text + "\n")
    
  // will listen for message to process ending in newline (\n)
  message, _ := bufio.NewReader(conn).ReadString('\n')
  // output message received
	fmt.Print("Message Received:", string(message))
	
	//Command here
	

  }



}
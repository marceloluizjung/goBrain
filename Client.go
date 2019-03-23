package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

    // connect to this socket
    conn, _ := net.Dial("tcp", "127.0.0.1:8091")
    for {   
    
    // will listen for message to process ending in newline (\n)
    messaget, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Message Received:", string(messaget))    

    switch messaget {
    case "1":
      fmt.Println("1 Key Pressed")
      break
    case "2":
      fmt.Println("2 Key Pressed")
      break
    }


    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
  }

 

}
package main

import "net"
import "fmt"
import "bufio"
//import "os"

func main() {

    // connect to this socket
    conn, _ := net.Dial("tcp", "127.0.0.1:8091")
    for {   
    
    // will listen for message to process ending in newline (\n)
    
    command, _ ,_:= bufio.NewReader(conn).ReadRune()
    //messaget, _ := bufio.NewReader(conn).ReadString(' ')
    
    // output message received
    fmt.Print("Command Received: ")    

   
    switch command {
    case '1':
      fmt.Println("teste1")
      fmt.Fprintf(conn, "true" + "\n")
      break
    case '2':
      fmt.Println("teste2")
      fmt.Fprintf(conn, "true" + "\n")
      break
    case '3':
      fmt.Println("teste3")
      fmt.Fprintf(conn, "true" + "\n")
      break
    default:
      fmt.Println("Command undefined")   
      fmt.Fprintf(conn, "false" + "\n")
    break
    }

    /*
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
  */
  }

 

}
package main

import (
	"TeliGo/cmds"
	"TeliGo/config"
	"TeliGo/utils"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// load configuration
func init() {
	config.Load()
}

// listener
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:23")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("[TeliGo] Listening on localhost:23")

	cmds.InitCommands()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go ConnHandler(conn)
	}
}

// connection handler
func ConnHandler(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted connection from:", conn.RemoteAddr())

	// buffer reads the first 64 bytes
	buffer := make([]byte, 64)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// buffer := string
	input := strings.TrimSpace(string(buffer))
	fmt.Println("Received first 64 bytes:", input)
	fmt.Println("-----------------------------------")
	reader := bufio.NewReader(conn)

	fmt.Fprintf(conn, ("\033[H\033[2J"))
	fmt.Fprintf(conn, "%s", utils.Gradient("[+] Welcome to Project TeliGo. An open sourced telnet project", []string{"ffffff", "00FFF7", "ffffff", "ffffff", "00FFF7"})+"\r\n")
	fmt.Fprintf(conn, "%s", utils.Gradient("[+] Developed by: YourPOV ", []string{"ffffff", "00FFF7", "ffffff", "ffffff", "00FFF7"})+"\r\n")

	for {
		if config.Users != nil && len(config.Users.Users) > 0 {
			fmt.Fprintf(conn, "\r\n\x1b[97m%s\x1b[96m@\x1b[97mTeli\x1b[96mGo\x1b[97m# ", config.Users.Users[0].Username)
		} else {
			fmt.Fprintf(conn, "\r\n\x1b[97mNix\x1b[96m@\x1b[97mTeli\x1b[96mGo\x1b[97m#\x1b[97m ")
		}

		input, err := reader.ReadString('\n')
		log.Println(input)
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		// prevent whitespace
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)

		// check command exists
		cmd := cmds.CanExecute(args[0])
		if cmd != nil {
			// we found our command
			cmd.Execute(args[1:], conn)
		} else {
			// our command isn't found
			fmt.Fprintf(conn, "\r\n\x1b[91mUnknown command\x1b[97m: %s\r\n", args[0])

		}
	}
}

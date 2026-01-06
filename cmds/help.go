package cmds

import (
	"fmt"
	"net"
)

func Help(args []string, conn net.Conn) {
	fmt.Fprint(conn, "\033[H\033[2J")
	write(conn, ("--- Available Commands ---\r\n"))
	for name, cmd := range Commands {
		write(conn, fmt.Sprintf("%s: %s\r", name, cmd.Description))
	}
}

package cmds

import (
	"net"
	"time"
)

func Logout(args []string, conn net.Conn) {
	write(conn, ("Disconnecting From Terminal"))
	time.Sleep(2 * time.Second)
	conn.Close()
	return
}

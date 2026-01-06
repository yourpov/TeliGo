package cmds

import (
	"net"
)

func PingPong(args []string, conn net.Conn) {
	write(conn, ("\x1b[96m[\x1b[97mTeliGo\x1b[96m] \x1b[96mPong\x1b[96m!\x1b[97m, \x1b[96mCommand handler is \x1b[92mlive\r\n\x1b[0m"))
}

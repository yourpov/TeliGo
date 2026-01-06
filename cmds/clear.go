package cmds

import (
	"net"
)

func Clear(args []string, conn net.Conn) {
	write(conn, ("\033[H\033[2J"))
	write(conn, ("                              \x1b[97mWelcome to  \x1b[97mTeli\x1b[96mGo\x1b[97m!\r\n\r\n\r\n"))
}

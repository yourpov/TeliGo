package cmds

import (
	"fmt"
	"net"
	"strings"
)

func Echo(args []string, conn net.Conn) {
	input := strings.Join(args, " ")
	write(conn, fmt.Sprintf("\x1b[96m[\x1b[97mTeliGo\x1b[96m] \x1b[97mEchoed\x1b[97m: \x1b[96m%s\r\n", input))
}

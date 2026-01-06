package cmds

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Command struct {
	Name        string
	Alias       []string
	Description string
	Execute     func(args []string, conn net.Conn)
}

var (
	Commands        = make(map[string]*Command)
	ShowDescription = true
)

func InitCommands() {
	LoadCommand(&Command{Name: "Clear", Alias: []string{"cls", "c"}, Description: "Clears the terminal", Execute: Clear})
	LoadCommand(&Command{Name: "Logout", Alias: []string{"exit"}, Description: "Logs out of terminal", Execute: Logout})
	LoadCommand(&Command{Name: "Help", Alias: []string{"?"}, Description: "Displays command catagorys", Execute: Help})
	LoadCommand(&Command{Name: "Echo", Alias: []string{"say"}, Description: "Echo a command", Execute: Echo})
	LoadCommand(&Command{Name: "Pingpong", Alias: []string{"ppong"}, Description: "Check if command handler is live", Execute: PingPong})
	LoadCommand(&Command{Name: "Animate", Alias: []string{"animate"}, Description: "Animated text", Execute: Animate})
	PrintCommandsDescription()
}

func LoadCommand(c *Command) {
	if _, f := Commands[c.Name]; f {
		log.Fatalf("\x1b[91m[\x1b[97mTeliGo\x1b[91m]\x1b[97m: \x1b[91mfound conflicting command names in (\x1b[97mcmd_loader.go\x1b[91m) Command\x1b[97m: %s\x1b[0m", c.Name)
	}
	Commands[c.Name] = c
}

func CanExecute(cmd string) *Command {
	for _, cmds := range Commands {
		if strings.EqualFold(cmds.Name, cmd) {
			return cmds
		}
	}
	// if orignal command name isn't found then loop over commands
	for _, v := range Commands {
		for _, s := range v.Alias {
			if strings.EqualFold(cmd, s) {
				// found command
				return v
			}
		}
	}

	return nil
}

// print command description
func PrintCommandsDescription() {
	if !ShowDescription {
		return
	}
	fmt.Printf("-----------------------------------\r\n")
	for name, cmd := range Commands {
		fmt.Printf("- %s - %s\n", name, cmd.Description)
	}
	fmt.Printf("-----------------------------------\r\n")
}

func write(c net.Conn, data string) {
	c.Write([]byte(fmt.Sprintf("%s\r\n", data)))
}

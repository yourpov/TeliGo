# TeliGo

little telnet server i put together in go. i was just messing around with network stuff some time ago and wanted to see how hard it'd be to make a cnc from scratch, releasing it now years later lol.

## what it does

- runs a telnet server on localhost:23
- has some basic commands with aliases
- gradient text because why not
- super simple command system if you wanna add stuff

## running it

you need go 1.21+ installed

```bash
go build
./TeliGo
```

or just use `screen.bat` if you're on windows

then connect with:

```bash
telnet localhost 23
```

## commands

- `clear` (or `cls`, `c`) - clears screen
- `logout` (or `exit`) - disconnects
- `help` (or `?`) - shows commands
- `echo` (or `say`) - echoes input back
- `pingpong` (or `ppong`) - command response test
- `animate` - some animated text

## adding commands

pretty straightforward. add a file in `cmds/`, write your function:

```go
func MyCommand(args []string, conn net.Conn) {
    fmt.Fprintf(conn, "whatever you want\r\n")
}
```

then register it in `cmd_loader.go`:

```go
LoadCommand(&Command{
    Name: "MyCommand", 
    Alias: []string{"CommandAlias1", "CommandAlias2"}, 
    Description: "does something idk", 
    Execute: MyCommand
    })
```

## config

users are in `config/settings.json` - currently just has a test account (Admin/Admin56).

---

it's open source so do whatever you want with it

**disclaimer:** this is NOT production ready. no encryption, basic auth etc

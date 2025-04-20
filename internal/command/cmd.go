package cmd

import "strings"

type cmdRegistry struct {
	meta map[string]*cmdMeta
}

type cmdMeta struct {
	name string
	args []string
	// execute //todo: write type
}

func ExecuteCommand(message string) {
	c := parseMessage(message);
	
}

func parseMessage(msg string) (c *cmdMeta) {
	var msgArgs []string = strings.Split(msg, " ")
	c = &cmdMeta{
		name: msgArgs[0],
		args: msgArgs[1:],
	}
	return
}

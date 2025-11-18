package cmd

import (
	"fmt"
	"log"
	"strings"
)

type CommandHandler interface {
	Execute(args []string) (any,error)
}

var registry = make(map[string]CommandHandler)

func Register(name string, handler CommandHandler){
	registry[name] = handler;
	log.Printf("Registered command: %s", name);
}

func Execute(msg string) (any,error) {
	name , args := parseMessage(msg);
	handler , ok := registry[name];
	if !ok {
		return nil,fmt.Errorf("unknown command: %s", name);
	}
	return handler.Execute(args);
}

func parseMessage(msg string) (name string, args []string) {
	msg = strings.ToUpper(msg);
	var msgArgs []string = strings.Fields(msg);
	name = msgArgs[0];
	args = msgArgs[1:];
	return
}

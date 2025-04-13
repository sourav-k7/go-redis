package services

import (
	"slices"
	"strings"

	"github.com/sourav-k7/go-redis/internal/constant"
	"github.com/sourav-k7/go-redis/internal/models"
)

func ExecuteCommand(message string) {
	var cmds []string = strings.Split(message, " ")
	cmdType := findCmdClass(cmds[0])
	if cmdType == constant.StringType {
		stringTypeCmdHandler(cmds[0], cmds[1], cmds[2])
	}
}

func findCmdClass(cmd string) constant.CommandClass {
	stringClassCmds := []string{"SET", "GET"}
	if slices.Contains(stringClassCmds, cmd) {
		return constant.StringType
	}
	return constant.DefaultType
}

func stringTypeCmdHandler(cmd, key, value string) {
	models.StringCache.Set(key,value);
}

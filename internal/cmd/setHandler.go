package cmd

import (
	"fmt"

	"github.com/sourav-k7/go-redis/internal/store"
)

type SetHandler struct{}

func (s *SetHandler) Execute(args []string) (any, error) {
	if len(args) < 2 {
		return nil,fmt.Errorf("SET requires key and value")
	}
	store.Set(args[0],args[1]);
	fmt.Printf("[SetHandler] %s = %s\n", args[0], args[1])
	return "OK",nil
}

func init() {
	Register("SET", &SetHandler{})
}

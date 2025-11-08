package cmd

import (
	"fmt"

	"github.com/sourav-k7/go-redis/internal/store"
)

type StringHandler struct{}

func (s *StringHandler) Execute(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("SET requires key and value")
	}
	store.Set(args[0],args[1]);
	fmt.Printf("[StringHandler] %s = %s\n", args[0], args[1])
	return nil
}

func init() {
	Register("SET", &StringHandler{})
}

package cmd

import (
	"fmt"

	"github.com/sourav-k7/go-redis/internal/store"
)

type GetHandler struct{}

func (s *GetHandler) Execute(args []string) (any, error) {
	if len(args) < 1 { 
		return nil, fmt.Errorf("GET requires key")
	}
	fmt.Printf("[GetHandler] key: %s\n", args[0])
	value, exist := store.Get(args[0]);
	if !exist {
		return "", nil;
	}
	return value, nil
}

func init() {
	Register("GET", &GetHandler{})
}

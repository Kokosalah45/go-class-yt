package hello

import (
	"fmt"
	"strings"
)

func Say (name []string) string{
	return fmt.Sprintf("Hello, %v", strings.Join(name, " ,")) 
}
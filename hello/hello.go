package hello

import (
	"fmt"
	"strings"
)

func Say(names []string) string {
	if len(names) == 0 || names == nil {
		names = []string{"World"}
	}
	return fmt.Sprintf("Hello, %v", strings.Join(names, ", "))
}

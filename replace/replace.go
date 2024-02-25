package replace

import (
	"bufio"
	"strings"
)

func Replace(old , new string , s *bufio.Scanner ) string {

	result := ""


	if old == "" || new == ""  {
		return result
	
	}

	for s.Scan(){
		line := s.Text()
		result += strings.Replace(line, old, new, -1)
	}
	
	return result
}
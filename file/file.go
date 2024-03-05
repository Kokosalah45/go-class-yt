package file

import (
	"bufio"
	"encoding/json"
	"os"
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



func Occurence(s *bufio.Scanner) map[string]int {

	occurenceMap := make(map[string]int)

	for s.Scan(){
		word := s.Text()
		occurenceMap[word]++
		
	}
	println("LOL")
	file , err := os.Open("./occurence.txt"); 
	if err != nil {
		if	buff , err := json.Marshal(occurenceMap); err != nil {
			file.Write(buff)
		}
	}else{
		println("file doesn't exists")
		file , err := os.Create("./occurence.txt")
		if err != nil {
			if	buff , err := json.Marshal(occurenceMap); err != nil {
				file.Write(buff)
			}
		}
	}
	return occurenceMap
}
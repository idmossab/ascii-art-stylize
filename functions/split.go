package function

import (
	//"strings"
)

func Split(str string) []string {
	slice := []string{}
	newStr := ""

	for i := 0; i < len(str); i++ {
		if str[i] == '\r' {
			if newStr != "" {
				slice = append(slice, newStr)
				newStr = ""
			}
			slice = append(slice, "")
			
			i += 1

		} else { 
			newStr += string(str[i])
		}
	}

	if newStr != "" {
		slice = append(slice, newStr)
	}
	return slice
}
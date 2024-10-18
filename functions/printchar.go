package function


func PrintWords(words []string, slice [][]string) string {
	str := ""
	for _, w := range words { // here to manipulate element by element
		for i := 1; i <= Hight_symbole; i++ { // to print the lines
			if len(w) == 0 { // here each "" means it is a newLine
				str += "\r\n"
				break // here to print \n only once
			}
			for _, e := range w {
				if int(e)-Min_printble >= 0 && int(e)-Min_printble <= len(slice)-1 {
					str += slice[int(e)-Min_printble][i] // print a line of each letter of the word
				} else { // if we encounter a special charter we return an error
					return "Special charactere is not allowed."
				}
			}
			if i < Hight_symbole { // after each line return to the line
				str += "\r\n"
			}
		}
	}
	if !ContainChar(str) {
		str = str[:len(str)-1]
	}

	return str
}

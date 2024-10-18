package function



func ContainChar(s string) bool {
	// check if the string contains char or just \n
	for _, r := range s {
		if int(r) >= Min_printble && int(r) <= Max_printble {
			return true
		}
	}
	return false
}
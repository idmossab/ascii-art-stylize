package function

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSymbols(banner string) ([][]string, error) {
	file, err := os.Open("./symbols/" + banner + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	symbole := []string{}
	symboles := [][]string{}

	for scanner.Scan() {
		symbole = append(symbole, scanner.Text())
		count++
		if count == Hight_symbole_with_ligne {
			symboles = append(symboles, symbole)
			symbole = []string{}
			count = 0
		}
	}

	if len(symboles) < Nbr_char_printble {
		return nil, fmt.Errorf("please make sure all characters are present in the file")
	}
	return symboles, nil
}

package web

import (
	"bufio"
	"os"
)

func Font(s string) map[int][]string {
	file, err := os.Open("draw/" + s + ".txt") // open file
	if err != nil {                            // hundel if was somme err
		return nil
	}
	artAlpha := make(map[int][]string) // ceat map
	defer file.Close()                 // close if func finsh
	i, j := 0, 32
	lines := []string{}
	scanner := bufio.NewScanner(file) // make scan in file
	for scanner.Scan() {
		
		lines = append(lines, scanner.Text()) // scan line by line
		i++
		if i%9 == 0 {
			artAlpha[j] = lines[1:] // add to map
			j++
			lines = []string{} // clean array
		}
	}
	return artAlpha
}

func Checkout(s string) (bool, rune) { // check if there are unvalide ascii in text
	for _, i := range s {
		if i > 126 || i < 32 {
			return false, i
		}
	}
	return true, 0
}

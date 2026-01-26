package print

import "fmt"

func indexToPrint(r rune) int {
	return int(r - 32)
}

func AsciiArt(data string, dataList []string) []string {
	runs := []rune(data)
	newList := []string{}

	for x := 0; x < len(runs); x++ {
		run := runs[x]
		if run == '\n' {
			newList = append(newList, "\n")
			continue
		} else {
			newList = append(newList, dataList[indexToPrint(run)])
		}
	}
	fmt.Println(newList)
	return newList
}

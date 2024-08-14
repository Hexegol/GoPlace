package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}


func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {
	// open src file
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, err
	}
	defer srcFile.Close()

	// open dst file
	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, nil, err
	}
	defer dstFile.Close()

	// prepare processing
	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}

		fmt.Fprintln(writer, res)
		lineIdx++
	}

	return occ, lines, nil
}

func main() {
	old := "Turks"
	new := "armenians"
	occ, lines, err := FindReplaceFile("turkk.txt", "armenians.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
		return
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==")
	fmt.Printf("Number of occurrences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")

}

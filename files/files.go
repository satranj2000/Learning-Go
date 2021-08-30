package files

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() {
	filename := ".//files//test.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic("File not found")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("String is  %v\n", t)
		fmt.Printf("Length is  %v\n", len(t))
	}

}

func SaveFile() {
	filename := ".//files//test.txt"
	f, err := os.Create(filename)
	if err != nil {
		panic("File created")
	}
	defer f.Close()
	f.WriteString("set GOENV=C:\\Users\\sathi\\AppData\\Roaming\\go\\env")

}

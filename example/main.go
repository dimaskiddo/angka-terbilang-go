package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	angka "github.com/dimaskiddo/angka-terbilang-go"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Angka Terbilang in Go")
	fmt.Println("---------------------------")

	fmt.Print("Masukan Angka: ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	output, err := angka.ToTerbilang(input)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	fmt.Printf("%v\n", output)
}

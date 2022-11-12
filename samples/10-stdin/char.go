package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	input := strings.Split(sc.Text(), "")
	for _, str := range input {
		fmt.Println(str)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hoge")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	var i interface{}
	json.Unmarshal([]byte(text), &i)
	fmt.Println(i)
}

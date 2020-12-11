package nestp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Parse() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	var i map[string]interface{}
	json.Unmarshal([]byte(text), &i)
	fmt.Println(i)
	traverse(i)
}

func traverse(node map[string]interface{}) {
	for _, v := range node {
		switch vtype := v.(type) {
		case string:
			fmt.Println(vtype)
		default:
			fmt.Println("default")
		}
	}
}

package nestp

import (
	"bufio"
	"fmt"
	"os"
)

func Parse() {
	scanner := bufio.NewScanner(os.Stdin)
	var j string
	for scanner.Scan() {
		j += scanner.Text()
	}
	s, e := Format([]byte(j))
	if e != nil {
		panic(e)
	}
	fmt.Println(string(s))
}

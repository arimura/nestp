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
	fmt.Println(j)
	s, e := Format([]byte(j))
	fmt.Println(string(s))
	fmt.Println(e)
}

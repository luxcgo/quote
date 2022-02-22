package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sb := strings.Builder{}
	for i := 2; i <= 42; i++ {
		s := fmt.Sprintf("### 第 %d 年\n\n```assembly\n```\n\n", i)
		sb.WriteString(s)
	}
	out, _ := os.Create("hrm")
	defer out.Close()
	out.Write([]byte(sb.String()))
}

package _022_generateParenthesis

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	strings := make([]string, 0)
	c(strings)
	fmt.Println(strings)
	generateParenthesis(3)
}

func c(s []string) {
	s = append(s, "1")
}

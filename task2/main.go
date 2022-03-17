package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	input = "{Пожалуйста|Просто} сделайте так, чтобы это {удивительное|крутое|простое} тестовое предложение {изменялось {быстро|мгновенно} случайным образом|менялось каждый раз}."
)

func parse(str string, m1 map[string]struct{}) {
	if strings.Index(str, "|") == -1 {
		m1[str] = struct{}{}
		return
	}
	re := regexp.MustCompile(`\{([^{]+?)\}`)
	buf := string(re.Find(re.Find([]byte(str))))
	split := strings.Split(buf[1:len(buf)-1], "|")
	for _, val := range split {
		newstr := strings.Replace(str, buf, val, 1)
		parse(newstr, m1)
	}
}

func main() {
	m1 := make(map[string]struct{})
	parse(input, m1)
	for k := range m1 {
		fmt.Println(k)
	}
}

package main

import (
	"fmt"
	"github.com/DeedsBaron/colors"
	"strconv"
	"testing"
)

var tests = []string{
	"{A|B|C} тест",
	"{A|B} {A|B}",
	"{A|D {B|C}} тест",
	"{Пожалуйста|Просто} сделайте так, чтобы это {удивительное|крутое|простое} тестовое предложение {изменялось {быстро|мгновенно} случайным образом|менялось каждый раз}.",
}

func TestParse(t *testing.T) {
	for i, str := range tests {
		m1 := map[string]struct{}{}
		fmt.Println(colors.Cyan + "Test №" + strconv.Itoa(i) + ":" + str + colors.Res)
		parse(str, m1)
		for k := range m1 {
			fmt.Println(k)
		}
	}
}

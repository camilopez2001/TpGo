package main


import (
	"fmt"
	"tudai2021.com/model"
	"strconv"
)

func Convertor(l string) model.Result{
	r := model.NewResult()
	r.Type = l[:2]
	r.Value = l[4:]
	x, _ := strconv.Atoi(l[2:4])
	r.Length = x
	return r
}

func main() {
	r := Convertor("TX02AB")
	fmt.Println(r)
}


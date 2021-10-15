package main


import (
	"fmt"
	"tudai2021.com/model"
	"strconv"
	"errors"
)



func Convertor(l string) (model.Result,error){
	r := model.NewResult()
	name := ""
    value := ""
    length := ""
	tipo := ""

    for i, pos := range l {
        len := i
        if len < 2 {
			if number(string(pos)) == true{
				return r, errors.New("Uno de los 2 primeros caracteres es un numero")
			}
            name += string(pos)
        }
        if len >= 2 && len < 4 {
			if number(string(pos)) == false{
				return r, errors.New("el 3 o 4 caracter es una letra, debe ser un valor")
			}
            length += string(pos)
        }
        if len >= 4 {
			if(len == 4){
				tipo = string(pos)
			}
            value += string(pos)
        }
        len++
    }

    length_int, _ := strconv.Atoi(length)

	lon := 0
	for i, _ := range value {
        lon = i
        lon++
    }
	if(length_int != lon){
		return r, errors.New("no coinciden el largo con la cantidad")
	}


	n := number(tipo)
	for _, pos := range value {
		if n{
			if number(string(pos)) == false {
				return r, errors.New("hay una letra")
			}
		} else {
				if number(string(pos)) == true {
					return r, errors.New("hay un numero")
				}
			}
		
	}


	r.Type = name
	r.Value = value
	r.Length = length_int
	return r, nil;
}


func number (n string) bool{
	_, err := strconv.Atoi(n);
	if err == nil{
			return true
	}else{
		return false
	}
} 

func main() {

	r, err:= Convertor("1N09ABC")
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(r)
	} 
	
	
}


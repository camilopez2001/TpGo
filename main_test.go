package main

import (
	"testing"
	"strconv"
)

func TestConversor(t *testing.T) {

	re := Convertor("aN100987654321")
	
	if !letra(re.Type[0:1]){
		t.Error("el primercaracter es un numero")
	} 
	if !letra(re.Type[1:2]){
		t.Error("el segundo caracter es un numero")
	}

	s := strconv.Itoa(re.Length)
	if !number(s[0:1]){
		t.Error("el 3 caracter es una letra, debe ser un valor")
	}
	if !number(s[1:2]){
		t.Error("el 4 caracter es una letra")
	}

	len := 0
	for i, _ := range re.Value {
        len = i
        len++
    }
	if(re.Length != len){
		t.Error("no coinciden el largo con la cantidad")
	}
	
	n := number(re.Value[0:1])
	for i, _ := range re.Value {
        if n{
			if !number(strconv.Itoa(i)) {
				t.Error("hay una letra")
			}
			
		} else {
			if !letra(strconv.Itoa(i)) {
				t.Error("hay un numero")
			}
		}
    }

}
func number (n string) bool{
	if "0" <= n || n <= "9" {
		return true
	}else{return false}
	
} 
func letra (s string) bool{
	if "a"<= s || s <="z" || s >= "A" || s <=  "Z" {
		return true
	}else{return false}
} 
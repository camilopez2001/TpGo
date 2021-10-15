package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestParser(t *testing.T) {
    var cases = []struct {
        Input   string // input string in order to be parsed
        Success bool   // paser result
        Type    string // the input type
        Value   string // the input value
        Length  int    // value length
    }{
        {"TX02AB", true, "TX", "AB", 2},
        {"NN100987654321", true, "NN", "0987654321", 10},
        {"TX06ABCDE", false, "", "", 0},
        {"NN04000A", false, "", "", 0},
    	{"NN04123A", false, "", "", 0},
        {"NN04AAA1", false, "", "", 0},
        {"1N09ABC", false, "", "", 0},
        {"1T08AAA1", false, "", "", 0},
        {"NN09000A", false, "", "", 0},
    }

    for _, testData := range cases {
        _, err := Convertor(testData.Input)
		assert.Equal(t, err , testData.Success)
    } 
}

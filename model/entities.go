package model

type Result struct {
	Type    string 
	Length  int
	Value   string

}

func NewResult() Result{
	return Result{"",0,""}
} 
 


package model

type Result struct {
	Type    string 
	Value   string
	Length  int
}

func NewResult() Result{
	return Result{"","",0}
} 
 


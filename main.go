package main 

import (
	"fmt"
	"github.com/yavuzpeksen/stringutil"
	"github.com/muratsplat/randgenerators"
)

func main() {
	
	fmt.Println(stringutil.Reverse("!gnaL oG ,olleH"))
	
	var size uint = 5
	random, err := randgenerators.GenereteRandString(size, randgenerators.UPPER_CASE_ALPHA, randgenerators.LOWER_CASE_ALPHA, randgenerators.NUMBER)
	if err == nil{
		fmt.Println("Generated string: ", random) 
	}
	
}


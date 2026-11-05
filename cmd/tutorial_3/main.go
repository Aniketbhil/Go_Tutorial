package main 
import (
	"errors";
	"fmt"
)

func main() {
	var numerator int = 12
	var denominator int = 0

	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil{
		fmt.Printf(err.Error())
	}else{
	fmt.Printf("The result is %v and remainder is %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divided by Zero")
		return 0, 0, err
	}

	var remainder = numerator%denominator
	var result int = numerator/denominator
	return result, remainder, err
}
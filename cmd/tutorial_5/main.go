package main 
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft()) }
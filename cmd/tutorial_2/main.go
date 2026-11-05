package main
import "fmt"
func main() {
	var myInt int = 12
	var myInt8 int8 = 12
	var myInt16 int16 = 12
	var myInt32 int32 = 12
	var myInt64 int64 = 12
	var myUnInt uint = 12

	fmt.Println(myInt)
	fmt.Println(myInt8)
	fmt.Println(myInt16)
	fmt.Println(myInt32)
	fmt.Println(myInt64)
	fmt.Println(myUnInt)

	var myFloat32 float32 = 23.00
	var myFloat64 float64 = 1234567888.777

	fmt.Println(myFloat32)
	fmt.Println(myFloat64)


	var string1 string = "Heello \nWorld"
	fmt.Println(string1)

	var string2 string = `Hello 
	World`
	fmt.Println(string2)

	var myRune rune  = 12
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	const myIntnew int = 99
	fmt.Println(myIntnew)
}
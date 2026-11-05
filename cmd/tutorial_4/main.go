package main
import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr1 := [...]int32{1,2,3}
	fmt.Println(intArr1)	

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"Adam":32, "Sarah": 19}
	fmt.Println(myMap2["Adam"])

	for name, age := range myMap2 {
		fmt.Printf("Name:%v, Age:%v",name, age)
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
}
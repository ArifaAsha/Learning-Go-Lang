// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// type List struct {
// 	Data int
// }

// // func (list List) SetData(val int) List {
// // 	List.Data = val
// // 	return list
// // }

// // func (list List) GetData() int {
// // 	return list.Data
// // }

// // func stringToSlice(A string) []int {
// // 	// strings := strings.Split(A, " ")
// // 	// ints := make([]int, len(strings))

// // 	// for i, s := range strings {
// // 	// 	ints[i], _ = strconv.Atoi(s)
// // 	// }
// // 	// fmt.Println(ints)

// // 	// Creates a float64 slice with length 0 and capacity len(a)
// // 	int[] y = Arrays.stream(A.split(" ")).mapToInt(Integer::parseInt).toArray();

// // 	return y
// // 	// fmt.Println(reflect.TypeOf(ints))
// // 	// fmt.Printf("%#v\n", ints)
// // }

// func stringsToIntegers(lines []string) []int {
// 	var integers []int
// 	for _, line := range lines {
// 		//fmt.Println(i, line)
// 		n, err := strconv.Atoi(line)
// 		if err != nil {
// 			// handle error
// 			fmt.Println(err)
// 			os.Exit(2)
// 			continue
// 		}
// 		integers = append(integers, n)
// 	}

// 	return integers

// }

// func stringToSlice(s []string) {
// 	fmt.Println(s[0])
// }

// // func stringToSlice(s []string) []int {
// // 	fmt.Println(s[0])
// // 	strs := strings.Split(s, " ")
// // 	// fmt.Println("Strs = ", strs)
// // 	// fmt.Println(reflect.TypeOf(strs))

// // 	// str2 := stringsToIntegers(strs)
// // 	// fmt.Println(str2)
// // 	res := make([]int, len(strs))
// // 	for i := range res {
// // 		res[i], _ = strconv.Atoi(strs[i])
// // 	}
// // 	// fmt.Println(res)
// // 	// res2 := stringSliceToIntSlice(strs)
// // 	return res
// // }

// func convertInterfaceToSlice(list interface{}) []int {
// 	// str := fmt.Sprintf("%v", x)
// 	str := fmt.Sprintf("%v", list) //making string from interface values
// 	// fmt.Println(str)
// 	// if len(strConv2) > 0 {
// 	convertedSlice := stringToSlice(str)
// 	return convertedSlice
// 	// }
// 	// return nil
// }

// func convertIntToSlice(x int) []int {
// 	slice := []int{x}
// 	// fmt.Println(slice)
// 	return slice

// }

// // func (list List) Cons(toAdd int) []int {
// // 	newList := convertInterfaceToSlice(list)
// // 	toAddSlice := convertIntToSlice(toAdd)
// // 	constructedList := append(newList, toAddSlice[:]...)
// // 	return constructedList
// // }

// // func Cons(list interface{}, toAdd int) []int {
// // 	newList := convertInterfaceToSlice(list)

// // 	toAddSlice := convertIntToSlice(toAdd)

// // 	constructedList := append(newList, toAddSlice[:]...)

// // 	return constructedList
// // }

// func Cons(list interface{}, toAdd int) []int {
// 	fmt.Println("Getting in list = ", list)
// 	newList := convertInterfaceToSlice(list)

// 	toAddSlice := convertIntToSlice(toAdd)

// 	constructedList := append(newList, toAddSlice[:]...)

// 	return constructedList
// }

// func main() {

// 	// var list []int
// 	// list := List{}

// 	// // list2 := Cons(list, 8)
// 	// // list3 := convertInterfaceToSlice(list2)
// 	// // list4 := Cons(list, 10)
// 	// list3 := Cons(list, 10)
// 	// list4 := Cons(list, 12)

// 	// fmt.Println(list3, list4)
// 	// fmt.Println(list4)

// 	list := []int{1, 2, 3}
// 	convertInterfaceToSlice(list)
// }

package main

import (
	"fmt"
	"os"
	"strconv"
)

type List struct {
	Data int
}

// func (list List) SetData(val int) List {
// 	List.Data = val
// 	return list
// }

// func (list List) GetData() int {
// 	return list.Data
// }

// func stringToSlice(A string) []int {
// 	// strings := strings.Split(A, " ")
// 	// ints := make([]int, len(strings))

// 	// for i, s := range strings {
// 	// 	ints[i], _ = strconv.Atoi(s)
// 	// }
// 	// fmt.Println(ints)

// 	// Creates a float64 slice with length 0 and capacity len(a)
// 	int[] y = Arrays.stream(A.split(" ")).mapToInt(Integer::parseInt).toArray();

// 	return y
// 	// fmt.Println(reflect.TypeOf(ints))
// 	// fmt.Printf("%#v\n", ints)
// }

func stringsToIntegers(lines []string) []int {
	var integers []int
	for _, line := range lines {
		//fmt.Println(i, line)
		n, err := strconv.Atoi(line)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
			continue
		}
		integers = append(integers, n)
	}

	return integers

}

func stringToSlice(s []string) {
	fmt.Println(s[0])
}

// func stringToSlice(s []string) []int {
// 	fmt.Println(s[0])
// 	strs := strings.Split(s, " ")
// 	// fmt.Println("Strs = ", strs)
// 	// fmt.Println(reflect.TypeOf(strs))

// 	// str2 := stringsToIntegers(strs)
// 	// fmt.Println(str2)
// 	res := make([]int, len(strs))
// 	for i := range res {
// 		res[i], _ = strconv.Atoi(strs[i])
// 	}
// 	// fmt.Println(res)
// 	// res2 := stringSliceToIntSlice(strs)
// 	return res
// }

func convertInterfaceToSlice(list interface{}) {
	// str := fmt.Sprintf("%v", x)
	str := fmt.Sprintf("%v", list) //making string from interface values

	// string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println(str, i)
	// if len(strConv2) > 0 {

}

func convertIntToSlice(x int) []int {
	slice := []int{x}
	// fmt.Println(slice)
	return slice

}

// func (list List) Cons(toAdd int) []int {
// 	newList := convertInterfaceToSlice(list)
// 	toAddSlice := convertIntToSlice(toAdd)
// 	constructedList := append(newList, toAddSlice[:]...)
// 	return constructedList
// }

// func Cons(list interface{}, toAdd int) []int {
// 	newList := convertInterfaceToSlice(list)

// 	toAddSlice := convertIntToSlice(toAdd)

// 	constructedList := append(newList, toAddSlice[:]...)

// 	return constructedList
// }

// func Cons(list interface{}, toAdd int) []int {
// 	fmt.Println("Getting in list = ", list)
// 	newList := convertInterfaceToSlice(list)

// 	toAddSlice := convertIntToSlice(toAdd)

// 	constructedList := append(newList, toAddSlice[:]...)

// 	return constructedList
// }

func main() {
	list := []int{1, 2, 3}
	convertInterfaceToSlice(list)
}

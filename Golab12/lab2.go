package main

import("fmt"
	"strconv"
//	"strings"
)

func main() {
	var IPlist1 [4]byte
	IPlist1[0] = 192
	IPlist1[1] = 168
	IPlist1[2] = 0
	IPlist1[3] = 1
	IPlist2 := [4]byte{192,128,63,63}
	printList(IPlist1)
	printList(IPlist2)
	fmt.Println(listEven(0,11))
	fmt.Println(listEven(1,0))
	fmt.Println(listEven(5,100))
	result:= []int{1,2,3,4,5}
	for i:=0; i < len(result); i++{
		fmt.Println("Before : " + strconv.Itoa(result[i]))
	}
	result=Map(result, DobleOrNothing)
	for i:=0; i < len(result); i++{
		fmt.Println("After : " + strconv.Itoa(result[i]))
	}
	// Строка 24-31
	countSymbols("ВСЕОБЩИЕ ФОРМЫ СУЩЕСТВОВАНИЯ МАТЕРИИ И ПРОСТРАНСТВО И ВРЕМЯ НЕ СУЩЕСТВУЮТ ВНЕ МАТЕРИИ И НЕЗАВИСИМО ОТ НЕЁ ТО СОБЫТЕ СУЩЕСТВУЕТ ВСЕГДА ТАМ ГДЕ ЕСТЬ ПРОСТРАНСТВО И ВРЕМЯ СОБЫТИЕ ПОЗНАЕТСЯ ЧЕЛОВЕКОМ ЧЕРЕЗ ИНФОРМАЦИЮ СОВОКУПНОСТЬ ИНФОРМАЦИИ ФОРМИРУЕТ СОБЫТИЕ ТО ЕСТЬ ЯВЛЯЕТСЯ ТОЧКОЙ В СОБЫТИИ ИСХОДЯ ИЗ ЭТОГО ИНФОРМАЦИЯ ТАКЖЕ ИМЕЕТ КАК ПРОСТРАНСТВЕННУЮ ТАК И ВРЕМЕННУЮ КООРДИНАТЫ И ТАК ЖЕ ЯВЛЯЕТСЯ МЕРОЙ СООТНОШЕНИЯ МЕЖДУ ")
}

func printList(IPlist [4]byte) {
	var IPString string
	var temp string
	for i:= 0; i < 4; i++{
		temp = strconv.Itoa(int(IPlist[i]))
		if i < 3{IPString = IPString + temp + "."} else {IPString = IPString + temp}
	}
	fmt.Println(IPString)
}

func countSymbols(text string){
	symbolMap := map[byte]int{}
	glued := text
//	broken := strings.Fields(text)
//	fmt.Println(broken)
//	glued := strings.Join(broken,"")
//	fmt.Println(glued)
	for i:=0; i < (len(glued)-1); i++{
		val, ok := symbolMap[glued[i]]
		if ok{
			symbolMap[glued[i]] = val+1
		}else{ symbolMap[glued[i]]=1}
	}
	for key, value:= range symbolMap{
		fmt.Println(string(key) + " : " + strconv.Itoa(value))
	}
}

func listEven(A int, B int) ([]int, error){
	if A%2 == 1{
		A=A+1
	}
	glue:= []int{A}
	A=A+2
	if B < A{
		return []int{}, fmt.Errorf("B should be bigger than A %s", "logical error")
	}
	i:=0
	for A <= B{
		glue = append(glue, A)
		A=A+2
		i=i+1
	}
	return glue, nil
}
func Map(n []int, f func(int) int) []int {
	result := make([]int, len(n))
	i:=0
	for i < len(n){
		result[i] = f(n[i])
		i++
	}
	return result
}

 func DobleOrNothing(x int) int{return x*x}
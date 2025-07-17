package main

import (
	"fmt"
	"goProject/internal/service"
)

/* const PassStatus = "pass"
const FailStatus = "fail" */
// Константы для статусов проверки
const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func taskNum(arr []int) bool {
	m := make(map[int]bool)
	flag := false

	for i := 0; i < len(arr); i++ {
		if m[arr[i]] {
			flag = true
		}
		m[arr[i]] = true
	}
	return flag
}

func doubleNum(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	return arr
}

type HealthCheck struct {
	Status    string
	ServiceID int
}

func NewHealthCheck(status string, id int) HealthCheck {
	return HealthCheck{
		status,
		id,
	}
}

func (h *HealthCheck) changeStatus(status string) {
	h.Status = status
}

func main() {
	var PersonService service.PersonService
	p := service.NewPerson("Harry", 22)
	init := NewHealthCheck("off", 5930)
	init.changeStatus("fail")
	PersonService = p
	//St := HealthCheck{Status: "pass", ServiceID: 1}
	//fmt.Println("id=before=", St.Status)
	//St.changeStatus("LALAL")
	//fmt.Println("id==", St.Status)
	//mapa := TestService.CreateMap();
	fmt.Println("before==", p.GetAge())
	defer fmt.Println("after DEFER==", p.GetAge())
	defer PersonService.ChangeAge(55)
	fmt.Println("after==", p.GetAge())


	//fmt.Println("init==", mapa)
	/* nums := []int{1, 2, 3, 4, 5, 6, 8}
	numsIncorrect := []int{1, 2, 3, 4, 5, 6, 8, 4, 2, 4}
	_ = numsIncorrect
	result := taskNum(nums)
	resultDouble := doubleNum(nums)
	fmt.Println("result==", result)
	fmt.Println("resultDouble==", resultDouble) */
	//strFunc()
	//numFunc()
	//massiveFunc()
	//structFunc()
	/* structArr := GenerateCheck()
	    fmt.Println("structArr==", structArr);
	    for i := 0; i < len(structArr); i++ {
	        if(structArr[i].Status == PassStatus) {
	            fmt.Println("id==", structArr[i].ServiceID);
	        }
		}
	    for _, check := range structArr {
			if check.Status == PassStatus {
				fmt.Println(check.ServiceID)
			}
		} */
	/*     myStr, myStrLength := prefixFunc("lalal");
	       fmt.Println("myStr==", myStr);
	       fmt.Println("myStrLength==", myStrLength); */
}

func deferFunc(num int) {
	fmt.Println("hello world, im defer function!", num)
}

func prefixFunc(str string) (withPrefix string, lengthStr int) {
	num := 5
	num2 := 10
	// defer срабатывает ДО return
	// порядок вызова снизу вверх
	defer deferFunc(num)  // 4
	defer deferFunc(num2) // 3
	// anonym func пример вызова сразу
	func(n int) {
		fmt.Println("anonym, num!", num) // 1
	}(num)
	num = num * 5
	fmt.Println("hello world, num!", num) // 2
	// как дефолтные значения
	// withPrefix ""
	// lengthStr 0
	fmt.Println("withPrefix before!", withPrefix)
	withPrefix = str + "_myPrefix"
	lengthStr = len(str)
	fmt.Println("withPrefix, after!", withPrefix)
	return withPrefix, lengthStr
}

// Функция GenerateCheck, которая возвращает слайс из структур HealthCheck
func GenerateCheck() []HealthCheck {
	var checks []HealthCheck
	for i := 0; i < 5; i++ {
		status := PassStatus
		if i%2 != 0 {
			status = FailStatus
		}
		checks = append(checks, HealthCheck{ServiceID: i, Status: status})
	}
	return checks
}

func structFunc() {

	var slice2 []HealthCheck
	_ = slice2
	/*  slice := func () {
	    slice2 := make([]HealthCheck, 5, 5)
	    for k, v := range m {
	        fmt.Println("value", v);
	        fmt.Println("key", k);
	    }
	    return []
	}() */

	/*  student := Sab {
	        name: "Mak",
	        age: 23,
	        hasJob: false,
	    }
	*/
	/*  student := HealthCheck{name: "Rob", age: 44, hasJob: true}
	    fmt.Println("student==", student); */
}

func mapFunc() {

	//m := make(map[string]string);
	m := map[string]string{
		"hyndai": "solaris",
		"kia":    "rio",
	}
	for k, v := range m {
		fmt.Println("value", v)
		fmt.Println("key", k)
	}
}

func massiveFunc() {
	// при создании всегда указывается длина массива
	//j := [4]int{1,2,3,4}
	//j1 := [...]int{1,2,3,4} // неизветсное колличество элементов
	//fmt.Println("v5", doubleArr(j, 2))
	// slice - последовательность переменной длины
	//slice1 := []int
	// c := cap(slice1)
	//l := len(slice1)
	slice2 := make([]int, 1, 1)
	slice2 = append(slice2, 77, 99, 22, 232)
	slice2 = append(slice2, 10)
	slice2 = append(slice2, 11)
	fmt.Printf("v5 slice2", cap(slice2))
	fmt.Printf("v5 slice2", slice2)
}

func doubleArr(slice2 []int, value int) []int {
	//var arr []int = make([]int, len(slice));
	var arr []int = slice2[:]
	for i, item := range slice2 {
		arr[i] = item * value
	}
	return arr
}

func numFunc() {
	/*
	   float32 6 знаков после , лучше использовать float64
	   float64 15 знаков после ,
	*/
	/* ignore variables */
	v1, v2 := 6.787, 7.4545
	//_ = v1
	//_ = v2
	fmt.Printf("v1", v1)
	fmt.Printf("v2", v2)
	/* блочный режим

	   var (
	       block1 = 22
	       block2 = 'ghkgk'
	   )

	   _ = block1
	   _ = block2*/
}

func strFunc() {
	/*
	   строки - последовательность неизменяемых байтов
	*/
	// v1 := "\nHello\t" // включает символы в строке
	// v2 := `\nHello\t` // исключает символы в строке
	// v3 := 'H' // байтовое представление 1 символ 72
	// v4 := v2[0] + v2[1] // 1 байт из строки  202

	v5 := "строка"
	v6 := "12345"
	v8 := "some originally"
	v8 += " ylalal"
	fmt.Println("v5", len(v5))                 // len возвращает число байтов в строке, а не символов! 12
	fmt.Println("v5", len(v6))                 // 5
	fmt.Println("вывод подстроки", v6[:2])     // 12
	fmt.Println("вывод подстроки v8", v8[16:]) // ylalal

	v7 := "alalal" + v5 // alalalстрока
	_ = v7
}

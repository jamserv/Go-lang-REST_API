package main

import (
	"fmt"
	"strconv"
)

func main() {
	title := "Mis datos"
	name := "janez"
	age := 27
	stranger := true

	fmt.Println("[" + title + "]")
	fmt.Println(name + "::age::" + strconv.Itoa(age)) //conver int to String or concatenate
	conto_int, _ := strconv.Atoi("579801")            //conver int to String or concatenate
	fmt.Println(name + "::identity::" + strconv.Itoa(conto_int))
	//fmt.Println(conto_int)
	fmt.Println(name + "::stranger::" + strconv.ParseBool(stranger))
}

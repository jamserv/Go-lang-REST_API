package main

import (
	"fmt"
	"strconv"
)

/**
if else decision bauble
 */
func ilse()  {
	limiters := 180
	diffDays := 30

	if limiters > diffDays {
		fmt.Println(strconv.Itoa(limiters))
		fmt.Println(strconv.Itoa(diffDays))
	} else if limiters < diffDays {
		fmt.Println(strconv.Itoa(limiters) + " is not major that " + strconv.Itoa(diffDays))
	} else {
		fmt.Println("are equals")
	}

}

func main() {
	ilse()
}

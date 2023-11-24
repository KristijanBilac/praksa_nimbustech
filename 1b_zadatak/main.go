package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a []int

	for {
		var input string
		fmt.Print("Unesite broj (ili 'kraj' za zavrsetak unosa): ")
		fmt.Scanln(&input)

		if input == "kraj" {
			break
		}

		list, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Unesite broj.")
			continue
		}

		a = append(a, list)
	}

	fmt.Println("Uneti brojevi:", a)

	m1 := 0

	for index, value := range a {
		if value > m1 {
			m1 = a[index]
		}
	}

	fmt.Printf("Najveci broj u listi je: %v \n", m1)

}

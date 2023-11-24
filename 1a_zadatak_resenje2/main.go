package main

import "fmt"

func Brisanjenajbr(list *[]int) {
	novalist := *list
	i1 := 0
	m1 := 0
	for index, value := range novalist {

		if value > m1 {
			m1 = novalist[index]
			i1 = index
		}
	}
	//fmt.Printf("najveci za broj za brisanje je: %v , sa indeksom: %v \n", m1, i1)					PROVERA
	*list = append(novalist[:i1], novalist[i1+1:]...)
}

func main() { //Funkcija sa pokazivacem na originalnu listu i brisanje prva 2 najveca broja iz nje preko pokazivaca u funkciji

	orglist := []int{-1, 2, 5, 6, 9, 11, 51, 201, 51}
	fmt.Printf("Lista PRE funckije brisanja izgleda ovako: %v\n", orglist)

	for i := 0; i <= 2; i++ {
		Brisanjenajbr(&orglist)
	}

	fmt.Printf("Lista Nakon funckije brisanja izgleda ovako: %v\n", orglist)

	m1 := 0

	for index, value := range orglist {
		if value > m1 {
			m1 = orglist[index]
		}
	}

	fmt.Printf("Najveci u listi: %v\n", m1)

}

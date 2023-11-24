package main

import "fmt"

func Najvecibr(a1 []int) int {
	i1 := 0
	m1 := 0
	for index, value := range a1 {

		if value > m1 {
			m1 = a1[index]
			i1 = index
		}
	}
	//fmt.Printf("najveci broj za brisanje je: %v , sa indeksom: %v \n", m1, i1)					PROVERA
	return i1
}

func main() { //Funkcija koja trazi najveci i vraca njegov index za brisanje je u mainu

	slice := []int{-1, 2, 5, 6, 9, 11, 51, 201, 51}
	fmt.Printf("Lista PRE funckije brisanja izgleda ovako: %v\n", slice)

	for f := 0; f <= 3; f++ {
		if f != 3 {
			i1 := Najvecibr(slice)
			slice = append(slice[:i1], slice[i1+1:]...)
		} else {
			Najvecibr(slice)
		}

	}
	fmt.Printf("Lista NAKON funckije brisanja izgleda ovako: %v\n", slice)

	max := slice[Najvecibr(slice)]
	fmt.Printf("Treci najveci broj iz liste: %v\n", max)

}

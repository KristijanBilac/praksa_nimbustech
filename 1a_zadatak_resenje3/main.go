package main

import "fmt"

func TreciNajveciBr(brojevi []int) int {
	var prvi, drugi, treci int

	for _, br := range brojevi {
		if br > prvi {
			treci = drugi
			drugi = prvi
			prvi = br
		} else if br > drugi {
			treci = drugi
			drugi = br
		} else if br > treci {
			treci = br
		}
	}

	return treci
}

func main() {
	brojevi := []int{-1, 2, 5, 6, 9, 11, 51, 201, 51}

	rez := TreciNajveciBr(brojevi)

	fmt.Println("Treci najveci broj:", rez)
}

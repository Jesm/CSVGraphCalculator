package main

import(
	"fmt"
)

func DisplayMatrix(arr [][]string){
	fmt.Printf("    ")
	for x:=range arr {
		fmt.Printf(" %2v ", x)
	}
	fmt.Printf("\n")

	for x, v:=range arr {
		fmt.Printf(" %2v ", x)		
		for _, v1:=range v {
			fmt.Printf(" %2v ", v1)
		}
		fmt.Printf("\n")
	}
}
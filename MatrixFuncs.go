package main

import(
	"fmt"
	"strconv"
)

func ConvertToIntMatrix(arr [][]string) [][]int {
	ret:=make([][]int, len(arr))

	for x, v:=range arr {
		ret[x]=make([]int, len(v))
		for y, v1:=range v {
			i, _:=strconv.Atoi(v1)
			ret[x][y]=i
		}
	}

	return ret
}

func DisplayMatrix(arr [][]int){
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
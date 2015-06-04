// Jonathan Szablevski
// Nº de Matrícula 0134469

package main

import(
	"fmt"
	"strconv"
)

func ConvertToIntMatrix(matriz [][]string) [][]int {
	ret:=make([][]int, len(matriz))

	for x, v:=range matriz {
		ret[x]=make([]int, len(v))
		for y, v1:=range v {
			i, _:=strconv.Atoi(v1)
			ret[x][y]=i
		}
	}

	return ret
}

func DisplayMatrix(matriz [][]int){
	fmt.Printf("    ")
	for x:=range matriz {
		fmt.Printf(" %2v ", x)
	}
	fmt.Printf("\n")

	for x, v:=range matriz {
		fmt.Printf(" %2v ", x)		
		for _, v1:=range v {
			fmt.Printf(" %2v ", v1)
		}
		fmt.Printf("\n")
	}
}

func GenerateDistanceMatrix(matriz [][]int, include_in bool) [][]int {	
	ret:=make([][]int, len(matriz))

	for x, _:=range matriz {
		ret[x]=CalculateMatrixDistancesOf(matriz, x, include_in)
	}

	return ret
}

func CalculateMatrixDistancesOf(matriz [][]int, index int, include_in bool) []int {
	size:=len(matriz[index])
	ret:=make([]int, size)
	visited:=make([]bool, size)

	curr:=index
	visited[curr]=true
	for {
		// Atualiza as distancias
		for x:=0; x<size; x++ {
			if visited[x] {
				continue
			}

			dist:=matriz[curr][x]
			if include_in && (dist==0 || (matriz[x][curr]!=0 && dist-matriz[x][curr]>0) ) {
				dist=matriz[x][curr]
			}

			if dist==0 {
				continue
			}

			dist+=ret[curr]
			if ret[x]==0||dist<ret[x] {
				ret[x]=dist
			}
		}

		// Define o nodo a ser verificado na proxima iteracao
		curr=-1

		for x:=0; x<size; x++ {
			if visited[x]||ret[x]==0 {
				continue
			}

			if curr==-1||ret[x]<ret[curr] {
				curr=x
			}
		}

		if curr!=-1 {
			visited[curr]=true
		} else {
			break
		}
	}

	return ret
}

func proximityOfCentralityOf(matriz [][]int, index int) float64 {
	size:=len(matriz)
	if size<2 {
		return 0
	}

	sum:=0
	for x, v:=range matriz[index] {
		if x!=index {
			sum+=v
		}
	}

	if sum==0 {
		return 0
	}

	return float64(size-1)/float64(sum)
}
package main

import(
	"fmt"
	// "io"
	"encoding/csv"
	"os"
)

func main() {
	arquivo, erro:=os.Open(os.Args[1])
	if erro!=nil {
		fmt.Printf("%v \n", erro)
		return
	}

	leitor:=csv.NewReader(arquivo)
	valores, erro:=leitor.ReadAll()
	if erro!=nil {
		fmt.Printf("%v \n", erro)
		return
	}

	matriz:=ConvertToIntMatrix(valores)

	fmt.Printf("Matriz lida:\n")
	DisplayMatrix(matriz)

	fmt.Printf("\nMatriz gerada:\n")
	DisplayMatrix(GenerateDistanceMatrix(matriz))
}
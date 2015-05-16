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
	
	for x:=0; x<len(valores); x++ {
		for y:=0; y<len(valores[x]); y++ {
			fmt.Printf(" %v ", valores[x][y])
		}
		fmt.Printf("\n")
	}
}
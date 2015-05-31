// Jonathan Szablevski
// Nº de Matrícula 0134469

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

	matrizDistancias:=GenerateDistanceMatrix(matriz)

	fmt.Printf("\nMatriz gerada:\n")
	DisplayMatrix(matrizDistancias)

	proximities:=make([]float64, len(matrizDistancias))
	for x, v:=range matrizDistancias {
		proximities[x]=proximityOfCentralityOf(v, x)
	}
	// sort.Sort(ByProximity{proximities})

	fmt.Printf("\nProximidade da centralidade de cada nodo:\n")
	for x, v:=range proximities {
		fmt.Printf("%v = %.2v%v\n", x, v*100, "%")
	}
}

// type Nodo struct {
// 	index int
// 	proximity float
// }

// func (s ByProximity) Less(i, j int) bool {
// 	return s.Organs[i].Name < s.Organs[j].Name
// }
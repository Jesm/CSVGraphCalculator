// Jonathan Szablevski
// Nº de Matrícula 0134469

package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"sort"
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

	proximities:=make(Nodes, len(matrizDistancias))
	for x, v:=range matrizDistancias {
		proximities[x]=Node{x, proximityOfCentralityOf(v, x)}
	}
	sort.Sort(proximities)

	fmt.Printf("\nProximidade da centralidade de cada nodo:\n")
	for _, v:=range proximities {
		fmt.Printf("%v = %.2v%v\n", v.index, v.proximity*100, "%")
	}
}
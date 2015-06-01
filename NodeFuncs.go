// Jonathan Szablevski
// Nº de Matrícula 0134469

package main

type Node struct {
	index int
	proximity float64
}

type Nodes []Node

func (arr Nodes) Len() int {
    return len(arr)
}

func (arr Nodes) Less(i, j int) bool {
    return arr[i].proximity > arr[j].proximity;
}

func (arr Nodes) Swap(i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}
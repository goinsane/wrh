package wrh

import (
	"math"

	"github.com/spaolacci/murmur3"
)

// Node is a container of Seed, Weight, Data to find responsible nodes. Data is
// a custom data to use free.
type Node struct {
	Seed   uint32
	Weight float64
	Data   interface{}
	score  float64
}

// Score calculates score by given key.
func (nd *Node) Score(key []byte) float64 {
	_, h2 := murmur3.Sum128WithSeed(key, nd.Seed)
	hf := uint64ToFloat64(h2)
	x := 1.0 / (-math.Log(hf))
	return nd.Weight * x
}

// Nodes is a Node slice and it implements sort.Interface to use sort package.
type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].score >= n[j].score
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

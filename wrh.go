package wrh

// Package wrh is implementation of weighted rendezvous hash algorithm.

func uint64ToFloat64(v uint64) float64 {
	ones := uint64(^uint64(0) >> (64 - 53))
	zeros := float64(1 << 53)
	return float64(v&ones) / zeros
}

// ResponsibleNodes calculates all scores from nodes and puts responsible nodes
// into respNodes.
func ResponsibleNodes(nodes Nodes, key []byte, respNodes Nodes) {
	respNodesLen := len(respNodes)
	if respNodesLen <= 0 {
		return
	}
	nodesLen := len(nodes)
	for i := range respNodes {
		respNodes[i] = Node{}
	}
	for i := 0; i < nodesLen; i++ {
		sc := nodes[i].Score(key)
		k := -1
		for j := 0; j < respNodesLen; j++ {
			if sc > respNodes[j].score {
				if k < 0 || (respNodes[k].score > respNodes[j].score) {
					k = j
				}
			}
		}
		if k >= 0 {
			respNodes[k] = nodes[i]
			respNodes[k].score = sc
		}
	}
}

// ResponsibleNodes2 calculates all scores from nodes and returns responsible nodes.
func ResponsibleNodes2(nodes Nodes, key []byte, count int) Nodes {
	if count <= 0 {
		return nil
	}
	respNodes := make(Nodes, count)
	ResponsibleNodes(nodes, key, respNodes)
	return respNodes
}

// FindBySeed finds a node by seed and returns its index. If seed is not exists,
// it returns -1.
func FindBySeed(nodes Nodes, seed uint32) int {
	for i, j := 0, len(nodes); i < j; i++ {
		if nodes[i].Seed == seed {
			return i
		}
	}
	return -1
}

// FindByMaxScore finds a node which has maximum score and returns its index.
// If nodes has no node, it returns -1.
func FindByMaxScore(nodes Nodes) int {
	result := -1
	var max float64
	for i, j := 0, len(nodes); i < j; i++ {
		if nodes[i].score > max {
			result = i
			max = nodes[i].score
		}
	}
	return result
}

// MergeNodes merges nodes1 and nodes2, appends mergedNodesIn, returns
// mergedNodes which has all merged nodes. mergedNodesIn can be nil.
func MergeNodes(nodes1, nodes2 Nodes, mergedNodesIn Nodes) (mergedNodes Nodes) {
	mergedNodes = mergedNodesIn
	for i := range nodes1 {
		if FindBySeed(mergedNodes, nodes1[i].Seed) >= 0 {
			continue
		}
		mergedNodes = append(mergedNodes, nodes1[i])
	}
	for i := range nodes2 {
		if FindBySeed(mergedNodes, nodes2[i].Seed) >= 0 {
			continue
		}
		mergedNodes = append(mergedNodes, nodes2[i])
	}
	return
}

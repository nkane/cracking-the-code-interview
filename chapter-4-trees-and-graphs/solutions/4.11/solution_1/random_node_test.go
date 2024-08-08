package random_node

import (
	"testing"
)

func TestTreeNode_GetRandomNode(t *testing.T) {
	root := &TreeNode{Data: 10}

	// Insert nodes
	root.InsertInOrder(5)
	root.InsertInOrder(15)
	root.InsertInOrder(3)
	root.InsertInOrder(7)
	root.InsertInOrder(13)
	root.InsertInOrder(17)

	nodeCount := make(map[int]int)
	for i := 0; i < 10000; i++ {
		node := root.GetRandomNode()
		nodeCount[node.Data]++
	}

	// Output the distribution for manual inspection
	for k, v := range nodeCount {
		t.Logf("Node %d: %d times", k, v)
	}

	// Simple check: Ensure that every node has been hit at least once
	expectedNodes := []int{10, 5, 15, 3, 7, 13, 17}
	for _, n := range expectedNodes {
		if nodeCount[n] == 0 {
			t.Errorf("Node %d was never selected", n)
		}
	}

	// Check for a reasonably even distribution (tolerance is arbitrary here, adjust as needed)
	for _, n := range expectedNodes {
		if nodeCount[n] < 1200 || nodeCount[n] > 1800 {
			t.Errorf("Node %d was selected an unusual number of times: %d", n, nodeCount[n])
		}
	}
}

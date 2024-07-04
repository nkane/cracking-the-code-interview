package queue_via_stack

import "testing"

func TestQueue(t *testing.T) {
	queue := CreateQueue[string]()
	queue.Add("A")
	queue.Add("B")
	queue.Add("C")

	v := queue.Peek()
	if v == nil {
		t.Fatalf("peek failed, nil results\n")
	}
	if *v != "A" {
		t.Fatalf("value expected A, got: %s", *v)
	}

	queue.Add("E")
	queue.Add("F")

	v = queue.Remove()
	if v == nil {
		t.Fatalf("peek failed, nil results\n")
	}
	if *v != "A" {
		t.Fatalf("value expected A, got: %s", *v)
	}
}

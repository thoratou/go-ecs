package ecs

import (
	"testing"
)

func TestIdStack(t *testing.T) {
	stack := newEntityIdStack()

	if stack.Len() != 0 {
		t.Fatal("the stack should be empty")
	}

	if !stack.Empty() {
		t.Fatal("the stack should be empty")
	}

	stack.Put(5)

	if stack.Len() != 1 {
		t.Fatal("the stack should contain one item")
	}

	if stack.Empty() {
		t.Fatal("the stack should not be empty")
	}

	if stack.Peek() != 5 {
		t.Fatal("peek value should be 5")
	}

	if stack.Peek() != 5 {
		t.Fatal("peek value should still be 5")
	}

	if stack.Pop() != 5 {
		t.Fatal("pop value should be 5")
	}

	if !stack.Empty() {
		t.Fatal("the stack should be empty")
	}

	stack.Put(1)
	stack.Put(2)
	stack.Put(3)
	stack.Put(4)

	if stack.Len() != 4 {
		t.Fatal("the stack should contain four items")
	}

	if stack.Pop() != 4 {
		t.Fatal("pop value should be 5")
	}

	if stack.Pop() != 3 {
		t.Fatal("pop value should be 5")
	}

	stack.Put(9)

	if stack.Pop() != 9 {
		t.Fatal("pop value should be 5")
	}

	if stack.Pop() != 2 {
		t.Fatal("pop value should be 5")
	}

	if stack.Pop() != 1 {
		t.Fatal("pop value should be 5")
	}
}

package data_structures_test

import (
	"fmt"
	"testing"

	"github.com/kalogs-c/all-go-rythm/data_structures"
)

func TestLinkedList(t *testing.T) {
	lk := data_structures.NewLinkedList[int]()

	lk.Append(1)
	lk.Append(2)
	lk.Prepend(3)
	lk.InsertAt(2, 4)
	lk.InsertAt(3, 6)
	lk.InsertAt(3, 5)

	slc := lk.ToSlice()

	if len(slc) == 0 {
		t.Errorf("Error len == 0 ")
	}

	for _, v := range slc {
		fmt.Printf("%d - ", v)
	}
	fmt.Println("FIM")
}

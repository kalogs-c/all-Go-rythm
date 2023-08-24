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

	err := lk.Remove(func(curr int) bool {
		return curr == 3
	})
	if err != nil {
		t.Error(err)
	}

	lk.ForEach(func(curr int) int {
		return curr * 20
	})

	slc := lk.ToSlice()

	if len(slc) == 0 {
		t.Errorf("Error len == 0")
	}

	for _, v := range slc {
		fmt.Printf("%d - ", v)
	}
	fmt.Println("FIM")
}

func TestLinkedListRemoveAt(t *testing.T) {
	lk := data_structures.NewLinkedList[int]()

	lk.Append(1)
	lk.Append(2)
	lk.Prepend(3)
	lk.InsertAt(2, 4)
	lk.InsertAt(3, 5)
	lk.InsertAt(2, 6)
	slc := lk.ToSlice()

	for _, v := range slc {
		fmt.Printf("%d - ", v)
	}
	fmt.Println("FIM")

	err := lk.RemoveAt(5)
	if err != nil {
		t.Error(err)
	}

	slc = lk.ToSlice()

	if len(slc) == 0 {
		t.Errorf("Error len == 0")
	}

	for _, v := range slc {
		fmt.Printf("%d - ", v)
	}
	fmt.Println("FIM TEST")
}

func TestLinkedListGet(t *testing.T) {
	lk := data_structures.NewLinkedList[int]()
	lk.Append(1)
	lk.Append(2)
	lk.Prepend(3)
	lk.InsertAt(2, 4)
	lk.InsertAt(3, 5)
	lk.InsertAt(2, 6)

	v, err := lk.Get(4)
	t.Log(v)
	if err != nil {
		t.Error(err)
	}

	if v != 6 {
		t.Error("Error")
	}
}

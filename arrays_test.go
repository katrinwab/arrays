package arrays

import (
	"testing"
)

func TestIsEqual(t *testing.T) {
	if !IsEqual([]int{}, []int{}) {
		t.Fail()
	}
	if !IsEqual([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
		t.Fail()
	}
	if IsEqual([]int{}, []int{1}) {
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	if !IsEqual(Intersect([]int{}, []int{}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Intersect([]int{1}, []int{}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Intersect([]int{1}, []int{2}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Intersect([]int{1, 2}, []int{2, 3}), []int{2}) {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	if !IsEqual(Union([]int{}, []int{}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Union([]int{1}, []int{}), []int{1}) {
		t.Fail()
	}
	if !IsEqual(Union([]int{2}, []int{1}), []int{1, 2}) {
		t.Fail()
	}
}

func TestDistinct(t *testing.T) {
	if !IsEqual(Distinct([]int{}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Distinct([]int{1, 1}), []int{1}) {
		t.Fail()
	}
	if !IsEqual(Distinct([]int{1, 2, 2, 3, 4, 4, 4}), []int{1, 2, 3, 4}) {
		t.Fail()
	}
}

func TestExcept(t *testing.T) {
	if !IsEqual(Except([]int{}, []int{}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Except([]int{}, []int{1, 2}), []int{}) {
		t.Fail()
	}
	if !IsEqual(Except([]int{1, 2}, []int{}), []int{1, 2}) {
		t.Fail()
	}
	if !IsEqual(Except([]int{1, 2}, []int{3, 4}), []int{1, 2}) {
		t.Fail()
	}
	if !IsEqual(Except([]int{1, 2, 3, 4, 5}, []int{2, 5}), []int{1, 3, 4}) {
		t.Fail()
	}
}

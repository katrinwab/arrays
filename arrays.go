package arrays

func IsEqual(a, b []int) bool {
	size := len(a)
	if size != len(b) {
		return false
	}
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func IntersectEx(r, a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	for i < asize && j < bsize {
		if a[i] < b[j] {
			i++
			continue
		}
		if a[i] > b[j] {
			j++
			continue
		}
		r = append(r, a[i])
		i++
		j++
	}
	return r
}

func Intersect(a, b []int) []int {
	size := len(a)
	if size > len(b) {
		size = len(b)
	}
	r := make([]int, 0, size)
	return IntersectEx(r, a, b)
}

func unionEx(r, a, b []int, all bool) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	for i < asize && j < bsize {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
			continue
		}
		if a[i] > b[j] {
			r = append(r, b[j])
			j++
			continue
		}
		if all {
			r = append(r, a[i], b[j])
		} else {
			r = append(r, a[i])
		}
		i++
		j++
	}
	for i < asize {
		r = append(r, a[i])
		i++
	}
	for j < bsize {
		r = append(r, b[j])
		j++
	}
	return r

}

func UnionAllEx(r, a, b []int) []int {
	return unionEx(r, a, b, true)
}

func UnionAll(a, b []int) []int {
	size := len(a) + len(b)
	r := make([]int, 0, size)
	return UnionAllEx(r, a, b)
}

func UnionEx(r, a, b []int) []int {
	return unionEx(r, a, b, false)
}

func Union(a, b []int) []int {
	size := len(a) + len(b)
	r := make([]int, 0, size)
	return UnionEx(r, a, b)
}

func Distinct(a []int) []int {
	i := 0
	j := 0
	size := len(a)
	for i < size {
		n := i + 1
		if n < size {
			if a[i] == a[n] {
				i++
				continue
			}
		}
		a[j] = a[i]
		i++
		j++
	}
	return a[:j]
}

func Except(a, b []int) []int {
	i := 0
	j := 0
	k := 0
	asize := len(a)
	bsize := len(b)
	for i < asize {
		uniq := true
		for j < bsize {
			if b[j] < a[i] {
				j++
				continue
			}
			if b[j] == a[i] {
				uniq = false
			}
			break
		}
		a[k] = a[i]
		if uniq {
			k++
		}
		i++
	}
	return a[:k]
}

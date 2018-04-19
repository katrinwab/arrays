package arrays

func IntersectEx(a, b, c []int) []int {
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
		c = append(c, a[i])
		i++
		j++
	}
	return c
}

func Intersect(a, b []int) []int {
	size := len(a)
	if size > len(b) {
		size = len(b)
	}
	c := make([]int, 0, size)
	return IntersectEx(a, b, c)
}

func UnionEx(a, b, c []int) []int {
	asize := len(a)
	bsize := len(b)
	i := 0
	j := 0
	for i < asize && j < bsize {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
			continue
		}
		if a[i] > b[j] {
			c = append(c, b[j])
			j++
			continue
		}
		c = append(c, a[i])
		c = append(c, b[j])
		i++
		j++
	}
	for i < asize {
		c = append(c, a[i])
		i++
	}
	for j < bsize {
		c = append(c, b[j])
		j++
	}
	return c
}

func Union(a, b []int) []int {
	size := len(a) + len(b)
	c := make([]int, 0, size)
	return UnionEx(a, b, c)
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
		valid := true
		for j < bsize {
			if b[j] < a[i] {
				j++
				continue
			}
			if b[j] == a[i] {
				valid = false
			}
			break
		}
		a[k] = a[i]
		if valid {
			k++
		}
		i++
	}
	return a[:k]
}

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

func Combine(a [][]int, f func([]int, []int) []int) []int {
	size := len(a)
	if size == 0 {
		a = [][]int{[]int{}, []int{}}
	}
	if size == 1 {
		a = append(a, []int{})
	}
	c := f(a[0], a[1])
	for i := 2; i < size; i++ {
		c = f(c, a[i])
	}
	return c
}

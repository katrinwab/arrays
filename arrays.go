package arrays

func Intersect(a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	csize := asize
	if csize > bsize {
		csize = bsize
	}
	c := make([]int, 0, csize)
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

func Union(a, b []int) []int {
	asize := len(a)
	bsize := len(b)
	c := make([]int, 0, (asize + bsize))
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

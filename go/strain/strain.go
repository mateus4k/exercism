package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(pred func(int) bool) (res Ints) {
	for _, x := range i {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

func (i Ints) Discard(pred func(int) bool) Ints {
	return i.Keep(func(n int) bool {
		return !pred(n)
	})
}

func (l Lists) Keep(pred func([]int) bool) (res Lists) {
	for _, x := range l {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

func (s Strings) Keep(pred func(string) bool) (res Strings) {
	for _, x := range s {
		if pred(x) {
			res = append(res, x)
		}
	}
	return
}

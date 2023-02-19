package generics

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  *T
}

func (lst *element[T]) Push(v T) {
	if lst.val == nil {
		lst.val = &v
	} else {
		p := lst
		for {
			if p.next == nil {
				p.next = &element[T]{val: &v}
				break
			}
			p = p.next
		}
	}
}

func (lst *element[T]) GetAll() []T {
	if lst.val == nil {
		return nil
	}

	elems := []T{*lst.val}

	for e := lst.next; e != nil; e = e.next {
		elems = append(elems, *e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := element[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

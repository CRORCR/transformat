package demo2_sort

type Sortinterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Bubble(s Sortinterface) {
	for i := s.Len() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s.Less(j+1, j) {
				s.Swap(j, j+1)
			}
		}
	}
}

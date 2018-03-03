package demo2_sort

type Sortinterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//这是个冒泡排序,之前是只能对int类型排序,现在可以把它改成对任意类型进行排序
func Bubble(s Sortinterface) {
	for i := s.Len() - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s.Less(j+1, j) { //控制顺序,逆序
				s.Swap(j, j+1) //控制顺序,逆序
			}
		}
	}
}

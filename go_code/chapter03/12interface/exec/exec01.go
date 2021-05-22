package exec

// 实现对结构体切片的排序

// 1. 声明hero结构体

type Hero struct {
	Name string
	Age  int
}

// 2. 声明hero结构体类型的切片类型

type HeroSlice []Hero

// 3. 实现Integer接口

func (hs HeroSlice) Len() int {
	return len(hs)
}

// 4. less方法决定使用什么标准进行排序

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age // 升序排序
}

// 5.Swap用来交换两个元素

func (hs HeroSlice) Swap(i, j int) {
	tmp := hs[i]
	hs[i] = hs[j]
	hs[j] = tmp
}

package ganzhiwuxin

var diZhiNumToName = [...]string{"寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}

type DiZhi struct {
	name  string
	num   int
	wuXin WuXin
}

func (d *DiZhi) Init(n interface{}) bool {
	for index, value := range diZhiNumToName {
		switch n := n.(type) {
		case int:
			if n == index+1 {
				d.name = value
				d.num = index + 1
				var w WuXin
				if (index+1)%3 == 0 {
					w.Init(3)
				} else {
					w.Init((index+1)/3 + (index+1)/7 + 1)
				}
				d.wuXin = w
				return true
			}
		case string:
			if n == value {
				d.name = value
				d.num = index + 1
				var w WuXin
				if (index+1)%3 == 0 {
					w.Init(3)
				} else {
					w.Init((index+1)/3 + (index+1)/7 + 1)
				}
				d.wuXin = w
				return true
			}
		}

	}
	return false

}

func (t DiZhi) GetName() string {
	return t.name
}

func (t DiZhi) GetNum() int {
	return t.num
}

func (d DiZhi) Is阳() bool {
	return d.num%2 != 0
}

func (d DiZhi) GetWuXin() WuXin {
	return d.wuXin
}

func (d DiZhi) String() string {
	return d.name
}

func (d DiZhi) Sheng(t2 GanZhi) bool {

	return d.wuXin.Sheng(t2.GetWuXin())

}

func (t DiZhi) Ke(t2 GanZhi) bool {
	return t.wuXin.Ke(t2.GetWuXin())
}

func (t DiZhi) Eq(t2 GanZhi) bool {
	switch t2 := t2.(type) {
	case DiZhi:
		return t.num == t2.num
	default:
		return false
	}
}

func (t DiZhi) Add(n int) DiZhi {
	var t1 DiZhi
	tmp := (t.num + n + 12) % 12
	if tmp == 0 {
		tmp = 12
	}
	t1.Init(tmp)
	return t1
}

func (t DiZhi) Sub(n DiZhi) int {
	return t.num - n.num
}

//地支三合
func (a DiZhi) Is三合(b, c DiZhi) bool {
	if a.Eq(b) || a.Eq(c) || b.Eq(c) {
		return false
	}
	return (a.num-b.num)%4 == 0 && (b.num-c.num)%4 == 0
}

//得到某地支的三合地支
/*
卯 + 4 == 未
未 + 4 == 亥
*/
func (a DiZhi) Get三合() []DiZhi {
	b := a.Add(4)
	c := b.Add(4)
	return []DiZhi{b, c}
}

//判断两个地支是否六合
func (a DiZhi) Is六合(b DiZhi) bool {
	return (a.num+b.num-11)%12 == 0
}

//判断两个地支是否六冲
func (a DiZhi) Is六冲(b DiZhi) bool {
	if a.Eq(b) {
		return false
	}
	return (a.num-b.num)%6 == 0
}

//得到某地支的三合地支
func (a DiZhi) Get六冲() (c DiZhi) {
	for i := 1; i < 13; i++ {
		var b DiZhi
		b.Init(i)
		if a.Is六冲(b) {
			return b
		}
	}
	return
}

//判断两个地支a是否刑b
func (a DiZhi) Is刑(b DiZhi) bool {

	xinGroup := make([][2]int, 0)

	//寅 巳 申
	xinGroup = append(xinGroup, [2]int{1, 4})
	xinGroup = append(xinGroup, [2]int{4, 7})
	xinGroup = append(xinGroup, [2]int{7, 1})
	//未 丑 戌
	xinGroup = append(xinGroup, [2]int{6, 12})
	xinGroup = append(xinGroup, [2]int{12, 9})
	xinGroup = append(xinGroup, [2]int{9, 6})

	//子 卯
	xinGroup = append(xinGroup, [2]int{2, 11})
	xinGroup = append(xinGroup, [2]int{11, 2})

	//辰 午 酉 亥
	xinGroup = append(xinGroup, [2]int{3, 3})
	xinGroup = append(xinGroup, [2]int{5, 5})
	xinGroup = append(xinGroup, [2]int{8, 8})
	xinGroup = append(xinGroup, [2]int{10, 10})
	for i := 0; i < len(xinGroup); i++ {
		if a.num == xinGroup[i][0] && b.num == xinGroup[i][1] {
			return true
		}
	}
	return false
}

//得到某地支所刑的地支
func (a DiZhi) Get刑() (c DiZhi) {
	for i := 1; i < 13; i++ {
		var b DiZhi
		b.Init(i)
		if a.Is刑(b) {
			return b
		}
	}
	return
}

package ganzhiwuxin

type GanZhi interface {
	GetName() string
	GetNum() int
	Is阳() bool
	GetWuXin() WuXin
	Sheng(t2 GanZhi) bool
	Ke(t2 GanZhi) bool
	Eq(t2 GanZhi) bool //干与干、支与支是否相等
	//Add(n int) GanZhi  将干或支前移n位
	//Sub(n GanZhi) int  两干或两支相隔多少位
}

func New(s string, n interface{}) (interface{}, bool) {
	switch s {
	case "WuXin":
		var w WuXin
		if w.Init(n) {
			return w, true
		}
	case "TianGan":
		var t TianGan
		if t.Init(n) {
			return t, true
		}
	case "DiZhi":
		var d DiZhi
		if d.Init(n) {
			return d, true
		}
	}
	return nil, false
}

var tianGanNumToName = [...]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

type TianGan struct {
	name  string
	num   int
	wuXin WuXin
}

func (t *TianGan) Init(n interface{}) bool {
	for index, value := range tianGanNumToName {
		switch n := n.(type) {
		case int:
			if n == index+1 {
				t.name = value
				t.num = index + 1
				var w WuXin
				w.Init((index + 2) / 2)
				t.wuXin = w
				return true
			}
		case string:
			if n == value {
				t.name = value
				t.num = index + 1
				var w WuXin
				w.Init((index + 2) / 2)
				t.wuXin = w
				return true
			}
		}

	}
	return false

}

func (t TianGan) GetName() string {
	return t.name
}

func (t TianGan) GetNum() int {
	return t.num
}

func (t TianGan) Is阳() bool {
	return t.num%2 != 0
}

//返回干的五行
func (t TianGan) GetWuXin() WuXin {
	return t.wuXin
}

//干生干、支
func (t TianGan) Sheng(t2 GanZhi) bool {

	return t.wuXin.Sheng(t2.GetWuXin())

}

//干与干、支相克
func (t TianGan) Ke(t2 GanZhi) bool {
	return t.wuXin.Ke(t2.GetWuXin())
}

func (t TianGan) Eq(t2 GanZhi) bool {
	switch t2 := t2.(type) {
	case TianGan:
		return t.num == t2.num
	default:
		return false
	}
}

func (t TianGan) Add(n int) TianGan {
	var t1 TianGan
	tmp := (t.num + n + 10) % 10
	if tmp == 0 {
		tmp = 10
	}
	t1.Init(tmp)
	return t1
}

func (t TianGan) Sub(n TianGan) int {
	return t.num - n.num
}

//天干五合
func (t TianGan) Is五合(t2 TianGan) bool {
	if t.Eq(t2) {
		return false
	}
	return (t.num-t2.num)%5 == 0
}

func (t TianGan) Get五合() (t1 TianGan) {
	for i := 1; i < 11; i++ {
		var b TianGan
		b.Init(i)
		if t.Is五合(b) {
			return b
		}
	}
	return
}

func (t TianGan) String() string {
	return t.name
}

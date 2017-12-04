package ganzhiwuxin

var wuXinNumToName = [...]string{"木", "火", "土", "金", "水"}

type WuXin struct {
	name string
	num  int
}

func (w *WuXin) Init(n interface{}) bool {
	for index, value := range wuXinNumToName {
		switch n := n.(type) {
		case int:
			if n == index+1 {
				w.name = value
				w.num = index + 1
				return true
			}
		case string:
			if n == value {
				w.name = value
				w.num = index + 1
				return true
			}
		}

	}
	return false
}

func (w WuXin) GetName() string {
	return w.name
}

func (w WuXin) Sheng(w2 WuXin) bool {
	return (w.num-w2.num-4)%5 == 0
}

func (w WuXin) Ke(w2 WuXin) bool {
	return (w.num-w2.num-3)%5 == 0
}

func (w WuXin) String() string {
	return w.name
}

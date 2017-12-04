package models

import (
	"ganzhiwuxin"
	"shipan"
)

type PaiPanJson struct {
	YueJiang string
	ZhanShi  string
	ZhanRi   string
	ZhouZhan string
}

type ShiPanJson struct {
	TianPan           []string
	SiKe              []string
	SanChuan          []string
	DunGan            []string
	TianJiang         []string
	SanChuanTianJiang []string
	LuQing            []string
}

func (this *PaiPanJson) PaiPan() ShiPanJson {
	var shiPanJson ShiPanJson

	var yueJiang, zhanShi ganzhiwuxin.DiZhi
	yueJiang.Init(this.YueJiang)
	zhanShi.Init(this.ZhanShi)

	// 求天盘
	var tp shipan.TianPan
	tp.Init(yueJiang, zhanShi)
	for i := 1; i < 13; i++ {
		var diZhi ganzhiwuxin.DiZhi
		diZhi.Init(i)
		shiPanJson.TianPan = append(shiPanJson.TianPan, tp.GetTianSheng(diZhi).GetName())
	}

	//求四课
	var riJiaZi ganzhiwuxin.JiaZi
	var gan ganzhiwuxin.TianGan
	var zhi ganzhiwuxin.DiZhi
	zhanRi := []rune(this.ZhanRi)
	gan.Init(string(zhanRi[0]))
	zhi.Init(string(zhanRi[1]))
	riJiaZi.Init(gan, zhi)

	var siKe shipan.SiKe
	siKe.Init(tp, riJiaZi)

	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetGan().GetName())
	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetGanYang().GetName())
	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetGanYing().GetName())

	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetZhi().GetName())
	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetZhiYang().GetName())
	shiPanJson.SiKe = append(shiPanJson.SiKe, siKe.GetZhiYing().GetName())

	//求三传
	var sanChuan shipan.SanChuan
	sanChuan.Init(siKe)
	shiPanJson.SanChuan = append(shiPanJson.SanChuan, sanChuan.Get初传().GetName())
	shiPanJson.SanChuan = append(shiPanJson.SanChuan, sanChuan.Get中传().GetName())
	shiPanJson.SanChuan = append(shiPanJson.SanChuan, sanChuan.Get末传().GetName())

	//求遁干
	for _, v := range shiPanJson.SanChuan {
		var d ganzhiwuxin.DiZhi
		d.Init(v)
		if duGan, ok := siKe.Get遁干(d); ok {
			shiPanJson.DunGan = append(shiPanJson.DunGan, duGan.GetName())
		} else {
			shiPanJson.DunGan = append(shiPanJson.DunGan, "")
		}
	}

	//求天将
	var tianJiang shipan.TianJiang
	var zhouZhan bool
	if this.ZhouZhan == "yes" {
		zhouZhan = true
	}
	tianJiang.Init(siKe, zhouZhan)
	for i := 1; i < 13; i++ {
		var a ganzhiwuxin.DiZhi
		a.Init(i)
		shiPanJson.TianJiang = append(shiPanJson.TianJiang, tianJiang.Get天将(tp.GetTianSheng(a)))
	}

	//三传天将
	for _, v := range shiPanJson.SanChuan {
		var d ganzhiwuxin.DiZhi
		d.Init(v)
		shiPanJson.SanChuanTianJiang = append(shiPanJson.SanChuanTianJiang, tianJiang.Get天将(d))
	}
	//六亲
	for _, v := range shiPanJson.SanChuan {
		var d ganzhiwuxin.DiZhi
		d.Init(v)
		g := siKe.GetGan()
		if g.Ke(d) {
			shiPanJson.LuQing = append(shiPanJson.LuQing, "财")
		} else if g.Sheng(d) {
			shiPanJson.LuQing = append(shiPanJson.LuQing, "子")
		} else if d.Sheng(g) {
			shiPanJson.LuQing = append(shiPanJson.LuQing, "父")
		} else if d.Ke(g) {
			shiPanJson.LuQing = append(shiPanJson.LuQing, "官")
		} else {
			shiPanJson.LuQing = append(shiPanJson.LuQing, "比")
		}
	}
	return shiPanJson
}

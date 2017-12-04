package ganzhiwuxin

type JiaZi struct {
	t TianGan
	d DiZhi
}

func (jz *JiaZi) Init(t TianGan, d DiZhi) bool {
	if (t.Is阳() && d.Is阳()) || (!t.Is阳() && !d.Is阳()) {
		jz.t = t
		jz.d = d
		return true
	}
	return false
}

func (jz JiaZi) GetGan() TianGan {
	return jz.t
}

func (jz JiaZi) GetDiZhi() DiZhi {
	return jz.d
}

func (jz JiaZi) Eq(h JiaZi) bool {
	return jz.t.Eq(h.t) && jz.d.Eq(h.d)
}

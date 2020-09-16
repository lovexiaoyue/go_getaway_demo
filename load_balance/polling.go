package load_balance

import "errors"

type Polling struct{
	curIndex int
	rss      []string
}

// 添加
func (r *Polling)Add(params ...string)error{
	if len(params) == 0{
		return errors.New("params len 1 at last")
	}
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *Polling)Next()string{
	if len(r.rss) == 0 {
		return ""
	}
	lens := len(r.rss)
	if r.curIndex >= lens{
		r.curIndex = 0
	}
	curAdd := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1)%lens
	return curAdd
}
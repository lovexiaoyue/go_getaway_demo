package load_balance

import (
	"errors"
	"strconv"
)

type WeightRoundBinBalance struct {
	currentIndex    int
	res             []*WeightNode
}

type WeightNode struct {
	Addr            string
	Weight          int  // 权重值
	CurrentWeight   int  // 节点当前权重
	EffectiveWeight int  // 有效权重
}

func (r *WeightRoundBinBalance)Add(params ...string)error  {

	if len(params) != 2{
		return errors.New("param len need 2")
	}

	parInt,err := strconv.ParseInt(params[1],10,64)
	if err != nil {
		return err
	}

	node := &WeightNode{Addr: params[0],Weight: int(parInt)}
	node.EffectiveWeight = node.Weight
	r.res = append(r.res, node)
	return nil
}

func (r *WeightRoundBinBalance)Next()string {
	total := 0
	var best *WeightNode
	for i := 0; i<len(r.res) ; i++{
		w := r.res[i]
		// step 1 统计所有有效权重之和
		total += w.EffectiveWeight

		// step 2 变更节点临时权重为节点临时权重+节点有效权重
		w.CurrentWeight += w.EffectiveWeight

		// step 3 有效权重默认与权重相同，通讯异常时 -1 ，通讯成功 +1，知道恢复到weight大小
		if w.EffectiveWeight < w.Weight{
			w.EffectiveWeight ++
		}

		// step 4 选择最大临时权重节点
		if best == nil || w.CurrentWeight > best.CurrentWeight{
			best = w
		}
	}

	if best == nil {
		return ""
	}
	// step 5 变更临时权重为，临时权重-有效权重之和
	best.CurrentWeight -= total
	return best.Addr
}
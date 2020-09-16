package load_balance

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32

type UInt32Slice []uint32

func (s UInt32Slice)Len() int {
	return len(s)
}

func (s UInt32Slice)Less(i , j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice)Swap(i , j int) {
	s[i] , s[j] = s[j] , s[i]
}

type ConsistentHashBalance struct {
	mux         sync.RWMutex
	hash        Hash
	replicas    int                 // 复制因子
	keys        UInt32Slice         // 已排序的节点hash切片
	hasMap      map[uint32]string   // 节点hash和key的map，键是哈希值，值是节点key
}

func NewConsistentHashBalance(replicas int , fun Hash) *ConsistentHashBalance {
	m := &ConsistentHashBalance{
		replicas: replicas,
		hash: fun,
		hasMap: make(map[uint32]string),
	}

	if m.hash == nil {
		// 最多32位，保证是一个2^32-1环
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

func (c *ConsistentHashBalance) IsEmpty() bool {
	return len(c.keys) == 0
}

func(c *ConsistentHashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at last")
	}

	addr := params[0]
	c.mux.Lock()
	defer c.mux.Unlock()
	// 结合复制因子计算所有的虚拟节点的hash值，并存入m.key中，同时在m.haspMap中保存哈希值和key的映射
	for i:=0 ; i < c.replicas ; i++ {
		hash := c.hash([]byte(strconv.Itoa(i) + addr))
		c.keys = append(c.keys, hash)
		c.hasMap[hash] = addr
	}
	// 对所有虚拟节点的hash值进行排序，方便之后进行二分查找
	sort.Sort(c.keys)
	return nil
}

func (c *ConsistentHashBalance) Get(key string) (string , error) {
	if c.IsEmpty() {
		return "", errors.New("node is empty")
	}
	hash := c.hash([]byte(key))
	// 通过二分查找获取最有节点，第一个“服务器hash”值大于“数据hash”值的就是最优“服务节点”
	idx := sort.Search(len(c.keys), func(i int) bool{ return c.keys[i] >= hash})
	// 如果查找结果 大于 服务节点哈希数组的最大索引，表示此该对象哈希值位于最后一个节点之后，那么放入第一个节点中。
	if idx == len(c.keys) {
		idx = 0
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.hasMap[c.keys[idx]] , nil
}

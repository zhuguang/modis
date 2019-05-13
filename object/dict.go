package object

var (
	HTSIZE_INIT int64 = 4
)

type dictEntry struct {
	key   interface{}
	value interface{}
	next  *dictEntry
}

type dictHashTable struct {
	table    []*dictEntry
	used     int64
	size     int64
	sizeMask int64
}

type dictType interface {
	hashFunc(k interface{}) int64
	keyCompare(key1 interface{}, key2 interface{}) int
}

//如果没有在重哈希，rehashIdx = -1
type dict struct {
	ht        [2]*dictHashTable
	rehashIdx int64
	dictType  *dictType
}

//不用指针传参
func dictCreate(dictType dictType) *dict {
	ht := &dictHashTable{make([]*dictEntry, HTSIZE_INIT), 0, HTSIZE_INIT, HTSIZE_INIT - 1}
	dict := &dict{[2]*dictHashTable{ht, nil}, -1, &dictType}
	return dict
}

//如果key已存在，不增加元素，返回 false，如果不存在，增加成功返回 true
func (d *dict) dictAdd(key interface{}, value interface{}) (bool, error) {
	var entryP *dictEntry = nil
	idx := d.dictKeyIndex(key, &entryP)
	if idx == -1 {
		return false, nil
	}
	ht := d.ht[0]
	if d.isRehashing() {
		ht = d.ht[1]
	}
	entry := ht.table[idx]
	// c语言中可以通过malloc返回一个给entry建的内存指针，然后把这个指针传入dictSetHashKey dictSetHashVal函数设置kv值，不需要担心entry为空
	// 但go语言只能通过先创建一个真实对象，然后把这个指针传入函数设置kv，但这样需要判断这个entry是否为新建的，逻辑不够简洁，所以传入index参数
	// 在函数中处理table[idx] == nil的问题
	if ht.table[idx] == nil {
		ht.table[idx] = &dictEntry{key, value, nil}
	} else {
		entry = &dictEntry{key, value, entry}
	}
	ht.used ++
	return true, nil
}

//给某个dictEntry赋值kv，绑定在哈希表上
func (entry *dictEntry) dictSetEntry(key interface{}, value interface{}) {
	entry.key = key
	entry.value = value
}

func (d *dict) dictReplace(key interface{}, value interface{}) (bool, error) {
	var entryP *dictEntry = nil
	index := d.dictKeyIndex(key, &entryP)
	if index == -1 {
		entryP.dictSetEntry(key, value)
		return true, nil
	}
	//todo 重复调用 dictKeyIndex
	if ok, err := d.dictAdd(key, value); ok {
		return ok, err
	} else if err != nil {
		return false, nil
	}
	return true, nil
}

//返回 key 被分配的index如果key已存在，返回-1
//用位运算计算模: a % (2^n) == a & (2^n -1)
func (d *dict) dictKeyIndex(key interface{}, entryP **dictEntry) int64 {
	hash := (*d.dictType).hashFunc(key)
	var entry *dictEntry
	var idx int64
	for _, ht := range d.ht {
		idx = hash & ht.sizeMask
		entry = ht.table[idx]
		for entry != nil {
			if (*d.dictType).keyCompare(entry.key, key) == 0 {
				*entryP = entry
				return -1
			}
			entry = entry.next
		}
		//如果没在重哈希，不用检查ht[1]
		if !d.isRehashing() {
			break
		}
	}
	return idx
}

//todo
func (d *dict) isRehashing() bool {
	return d.rehashIdx != -1
}

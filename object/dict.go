package object

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
	ht := &dictHashTable{make([]*dictEntry, 4), 0, 0, 0}
	dict := &dict{[2]*dictHashTable{ht, nil}, -1, &dictType}
	return dict
}

func (d *dict) dictAdd(key interface{}, value interface{}) {
	idx := d.dictKeyIndex(key)
	if idx == -1 {
		return
	}
	ht := d.ht[0]
	if d.isRehashing() {
		ht = d.ht[1]
	}
	entry := ht.table[idx]
	if entry == nil {
		ht.table[idx] = &dictEntry{key, value, nil}
	} else {
		entry = &dictEntry{key, value, entry}
	}
	ht.size ++
	ht.sizeMask ++
}

//返回 key 被分配的index如果key已存在，返回-1
//用位运算计算模: a % (2^n) == a & (2^n -1)
func (d *dict) dictKeyIndex(key interface{}) int64 {
	hash := (*d.dictType).hashFunc(key)
	var entry *dictEntry
	var idx int64
	for _, ht := range d.ht {
		idx = hash & ht.sizeMask
		entry = ht.table[idx]
		for entry != nil {
			if (*d.dictType).keyCompare(entry.key, key) == 0 {
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

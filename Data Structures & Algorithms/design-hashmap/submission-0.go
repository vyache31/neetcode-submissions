type Item struct {
	key   int
	value int
}

type MyHashMap struct {
	slots [][]Item
	size  int
	cap   int
}

func Constructor() MyHashMap {
	size := 0
	cap := 1
	slots := make([][]Item, cap)
	return MyHashMap{
		slots: slots,
		size:  size,
		cap:   cap,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	if this.loadFactor() > 0.6 {
		this.resize()
	}
	i := key % this.cap
	if this.slots[i] == nil || len(this.slots[i]) == 0 {
		this.size++
		this.slots[i] = []Item{
			{key, value},
		}
	} else {
		for idx, item := range this.slots[i] {
			if item.key == key {
				this.slots[i][idx] = Item{key, value}
				return
			}
		}
		this.slots[i] = append(this.slots[i], Item{key, value})
	}
}

func (this *MyHashMap) Get(key int) int {
	i := key % this.cap
	if this.slots[i] == nil {
		return -1
	}
	for _, item := range this.slots[i] {
		if item.key == key {
			return item.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	i := key % this.cap
	if this.slots[i] == nil || len(this.slots[i]) == 0 {
		return
	}
	itemIdx := -1
	for idx, item := range this.slots[i] {
		if item.key == key {
			itemIdx = idx
			this.size--
			break
		}
	}
	if itemIdx == -1 {
		return
	}
	bucket := this.slots[i]
	this.slots[i] = append(bucket[:itemIdx], bucket[itemIdx+1:]...)
}

func (this *MyHashMap) resize() {
	var newCap int
	if this.cap < 256 {
		newCap = this.cap * 2
	} else {
		newCap = int(float64(this.cap) * 1.25)
	}
	oldSlots := this.slots
	this.slots = make([][]Item, newCap)
	this.cap = newCap
	this.size = 0
	for _, bckt := range oldSlots {
		for _, item := range bckt {
			i := item.key % newCap
			if this.slots[i] == nil {
				this.slots[i] = []Item{item}
				this.size++
			} else {
				this.slots[i] = append(this.slots[i], item)
			}
		}
	}
}

func (this *MyHashMap) loadFactor() float64 {
	return float64(this.size) / float64(this.cap)
}



/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
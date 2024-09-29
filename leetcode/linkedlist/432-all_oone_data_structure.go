package linkedlist

type CountNode struct {
	value         int
	keys          map[string]bool
	before, after *CountNode
}

type AllOne struct {
	keyToCnt   map[string]*CountNode
	head, tail *CountNode
}

func Constructor() AllOne {
	return AllOne{
		keyToCnt: make(map[string]*CountNode),
	}
}

func getNewNode(value int) *CountNode {
	return &CountNode{value: value, keys: make(map[string]bool)}
}

func (this *AllOne) Inc(key string) {
	if this.IsEmpty() {
		this.head = getNewNode(1)
		this.head.keys[key] = true
		this.tail = this.head
		this.keyToCnt[key] = this.head
		return
	}

	curNode := this.keyToCnt[key]
	if curNode == nil {
		// if there's node for cnt 1 already, just add key
		if this.tail.value == 1 {
			this.tail.keys[key] = true
		} else {
			// add new tail
			newTail := getNewNode(1)
			newTail.keys[key] = true
			this.tail.after = newTail
			newTail.before = this.tail
			this.tail = newTail
		}
		this.keyToCnt[key] = this.tail
	} else {
		beforeNode := curNode.before
		if beforeNode == nil {
			// create new head
			beforeNode = getNewNode(curNode.value + 1)
			this.head = beforeNode
		} else if beforeNode.value != curNode.value+1 {
			// if beforeNode is not consecutive one, create a new one
			betterBeforeNode := getNewNode(curNode.value + 1)
			beforeNode.after = betterBeforeNode
			betterBeforeNode.before = beforeNode
			beforeNode = betterBeforeNode
		}
		beforeNode.keys[key] = true

		// delete key in old node
		delete(curNode.keys, key)

		// remove nodes if empty
		if len(curNode.keys) == 0 {
			beforeNode.after = curNode.after
			if curNode.after != nil {
				curNode.after.before = beforeNode
			} else {
				this.tail = beforeNode
			}
		} else {
			beforeNode.after = curNode
			curNode.before = beforeNode
		}

		this.keyToCnt[key] = beforeNode
	}
	return
}

func (this *AllOne) Dec(key string) {
	if this.IsEmpty() {
		return
	}

	curNode := this.keyToCnt[key]
	if curNode == nil {
		return
	}

	if curNode.value == 1 {
		delete(curNode.keys, key)

		if len(curNode.keys) != 0 {
			this.keyToCnt[key] = nil
			return
		}

		if curNode.before != nil {
			curNode.before.after = nil
			this.tail = curNode.before
		} else {
			this.head = nil
			this.tail = nil
		}
		this.keyToCnt[key] = nil
		return
	}

	afterNode := curNode.after
	if afterNode == nil {
		// create new tail
		afterNode = getNewNode(curNode.value - 1)
		this.tail = afterNode
	} else if afterNode.value != curNode.value-1 {
		// if beforeNode is not consecutive one, create a new one
		betterAfterNode := getNewNode(curNode.value - 1)
		afterNode.before = betterAfterNode
		betterAfterNode.after = afterNode
		afterNode = betterAfterNode
	}
	afterNode.keys[key] = true

	// delete key in old node
	delete(curNode.keys, key)

	// remove nodes if empty
	if len(curNode.keys) == 0 {
		afterNode.before = curNode.before
		if curNode.before != nil {
			curNode.before.after = afterNode
		} else {
			this.head = afterNode
		}
	} else {
		afterNode.before = curNode
		curNode.after = afterNode
	}

	this.keyToCnt[key] = afterNode
	return
}

func (this *AllOne) GetMaxKey() string {
	if this.IsEmpty() {
		return ""
	}
	for key, _ := range this.head.keys {
		return key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.IsEmpty() {
		return ""
	}
	for key, _ := range this.tail.keys {
		return key
	}
	return ""
}

func (this *AllOne) IsEmpty() bool {
	return this.head == nil
}

package warehouse

func (w *Warehouse) AddItem(item string, quantity int) {
	w.rwMutex.Lock()
	defer w.rwMutex.Unlock()
	w.items[item] += quantity
}

func (w *Warehouse) RemoveItem(item string, quantity int) {
	w.rwMutex.Lock()
	defer w.rwMutex.Unlock()
	w.items[item] -= quantity
}

func (w *Warehouse) GetQuantity(item string) int {
	w.rwMutex.RLock()
	defer w.rwMutex.RUnlock()
	return w.items[item]
}

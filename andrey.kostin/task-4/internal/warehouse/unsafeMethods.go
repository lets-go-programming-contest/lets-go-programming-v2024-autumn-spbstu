package warehouse

func (w *Warehouse) AddItemUnsafe(item string, quantity int) {
	w.items[item] += quantity
}

func (w *Warehouse) RemoveItemUnsafe(item string, quantity int) {
	w.items[item] -= quantity
}

func (w *Warehouse) GetQuantityUnsafe(item string) int {
	return w.items[item]
}

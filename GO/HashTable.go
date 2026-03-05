package main

type Pair struct {
	key   rune
	value int
}

type OpenAddressingHashTable struct {
	table    []Pair
	occupied []bool
	capacity int
	size     int
}

func NewOpenAddressingHashTable(capacity int) *OpenAddressingHashTable {
	if capacity == 0 {
		capacity = 256
	}
	return &OpenAddressingHashTable{
		table:    make([]Pair, capacity),
		occupied: make([]bool, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (h *OpenAddressingHashTable) hash(key rune, attempt int) int {
	return (int(key) + attempt) % h.capacity
}

func (h *OpenAddressingHashTable) Insert(key rune, value int) {
	// Проверка на переполнение
	if h.size >= h.capacity {
		return
	}
	
	// Линейное пробирование
	index := int(key) % h.capacity
	startIndex := index
	
	for {
		if !h.occupied[index] {
			h.table[index] = Pair{key, value}
			h.occupied[index] = true
			h.size++
			return
		}
		
		if h.occupied[index] && h.table[index].key == key {
			h.table[index].value = value
			return
		}
		
		// Следующий индекс
		index = (index + 1) % h.capacity
		
		// Если обошли всю таблицу и вернулись к началу
		if index == startIndex {
			break
		}
	}
}

func (h *OpenAddressingHashTable) Search(key rune) (int, bool) {
	for attempt := 0; attempt < h.capacity; attempt++ {
		index := h.hash(key, attempt)
		
		if h.occupied[index] && h.table[index].key == key {
			return h.table[index].value, true
		}
		
		if !h.occupied[index] {
			break
		}
	}
	return 0, false
}

func (h *OpenAddressingHashTable) Remove(key rune) bool {
	for attempt := 0; attempt < h.capacity; attempt++ {
		index := h.hash(key, attempt)
		
		if h.occupied[index] && h.table[index].key == key {
			h.occupied[index] = false
			h.size--
			return true
		}
		
		if !h.occupied[index] {
			break
		}
	}
	return false
}

func (h *OpenAddressingHashTable) Clear() {
	for i := 0; i < h.capacity; i++ {
		h.occupied[i] = false
	}
	h.size = 0
}

func (h *OpenAddressingHashTable) IsEmpty() bool {
	return h.size == 0
}

func (h *OpenAddressingHashTable) GetSize() int {
	return h.size
}

func (h *OpenAddressingHashTable) ToVector() []Pair {
	result := []Pair{}
	for i := 0; i < h.capacity; i++ {
		if h.occupied[i] {
			result = append(result, h.table[i])
		}
	}
	return result
}

func (h *OpenAddressingHashTable) FromVector(elements []Pair) {
	h.Clear()
	for _, element := range elements {
		h.Insert(element.key, element.value)
	}
}

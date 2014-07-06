/*
A very simply sorted map.
*/
package index

type key string
type val float64

type SortedMap struct {
	data  map[key]val
	order []key
	ord   map[key]int
}

func NewSortedMap() *SortedMap {
	m := new(SortedMap)
	m.data = make(map[key]val)
	m.ord = make(map[key]int)
	return m
}

func (m *SortedMap) Put(item key, value val) {
	if _, ok := m.data[item]; ok == false {
		m.order = append(m.order, item)
		m.ord[item] = len(m.ord)
	}
	m.data[item] = value
	m.fixOrder(m.ord[item])
}

func (m *SortedMap) Remove(item key) {
	pos := m.ord[item]
	delete(m.ord, item)
	for _, v := range m.order[pos+1:] {
		m.ord[v] = m.ord[v] - 1
	}
	m.order = append(m.order[:pos], m.order[pos+1:]...)
	delete(m.data, item)
}

func (m *SortedMap) fixOrder(position int) {
	currentPos := position

	if prevPos := currentPos - 1; prevPos >= 0 {
		curr := m.order[currentPos]
		prev := m.order[prevPos]
		if m.data[curr] < m.data[prev] {
			m.order[currentPos], m.order[prevPos] = m.order[prevPos], m.order[currentPos]
			m.ord[m.order[currentPos]] = currentPos
			m.ord[m.order[prevPos]] = prevPos
			m.fixOrder(prevPos)
		}
	}

	if nextPos := currentPos + 1; nextPos < len(m.ord) {
		curr := m.order[currentPos]
		next := m.order[nextPos]
		if m.data[curr] > m.data[next] {
			m.order[currentPos], m.order[nextPos] = m.order[nextPos], m.order[currentPos]
			m.ord[m.order[currentPos]] = currentPos
			m.ord[m.order[nextPos]] = nextPos
			m.fixOrder(nextPos)
		}
	}
}

func (m *SortedMap) Dec(item key, amount val) (newValue val) {
	if _, ok := m.data[item]; ok == false {
		m.Put(item, 0-amount)
		return amount
	}
	m.data[item] = m.data[item] - amount
	m.fixOrder(m.ord[item])
	return m.data[item]
}

func (m *SortedMap) Inc(item key, amount val) (newValue val) {
	if _, ok := m.data[item]; ok == false {
		m.Put(item, amount)
		return amount
	}
	m.data[item] = m.data[item] + amount
	m.fixOrder(m.ord[item])
	return m.data[item]
}

func (m *SortedMap) Get(item key) val {
	return m.data[item]
}

func (m *SortedMap) OrderedKeys() []key {
	return m.order
}

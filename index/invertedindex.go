package index

import (
	"container/list"
	"math"
)

type Doc interface {
	GetText() string
	TF() map[string]int
}

type Inverted struct {
	index        map[string]*list.List
	count        map[string]uint64
	ngrams       map[string]string
	NumDocuments uint64
	NumTokens    uint64
}

func (i *Inverted) Init() *Inverted {
	i.index = make(map[string]*list.List)
	i.count = make(map[string]uint64)
	return i
}

func NewInvertedIndex() *Inverted {
	return new(Inverted).Init()
}

func (i *Inverted) GetDocs(term string) []Doc {
	docs := []Doc{}

	l := i.index[term]
	if l == nil {
		return nil
	}

	for e := l.Front(); e != nil; e = e.Next() {
		docs = append(docs, e.Value.(Doc))
	}

	return docs
}

func (i *Inverted) TF(n string) uint64 {
	return i.count[n]
}

func (i *Inverted) IDF(n string) float64 {
	docs := i.index[n]
	l := 0
	if docs != nil {
		l = docs.Len()
	}
	return math.Log2(float64(i.NumDocuments+1) / float64(l+1))
}

func (i *Inverted) StringsToGrams(t map[string]uint) *map[string]uint {
	nt := make(map[string]uint)
	for k, v := range t {
		nt[k] = v
	}
	return &nt
}

func (i *Inverted) AddDocument(d Doc) {
	doc := d.TF()
	for k, v := range doc {
		i.NumTokens += uint64(v)
		i.count[k] += uint64(v)
		if _, ok := i.index[k]; ok == false {
			i.index[k] = list.New()
		}
		i.index[k].PushBack(d)
	}
	i.NumDocuments++
}

// TODO: Update word count to reflect removed documents?
func (i *Inverted) RemoveDocument(d Doc) {
	for k, _ := range d.TF() {
		for e := i.index[k].Front(); e != nil; e = e.Next() {
			if e.Value == d {
				i.index[k].Remove(e)
				if i.index[k].Len() == 0 {
					delete(i.index, k)
				}
				break
			}
		}
	}

	i.NumDocuments--
}

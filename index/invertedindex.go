package index

import (
	"container/list"
	"math"
	"strings"
)

type Doc interface {
	Text() string
	Tokens() map[string]int
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

func shouldRemove(c byte) bool {
	switch {
	case 'a' <= c && c <= 'z':
		return false
	case 'A' <= c && c <= 'Z':
		return false
	case '0' <= c && c <= '9':
		return false
	case c == ' ' || c == '&':
		return false
	}
	return true
}

func clean(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if shouldRemove(s[i]) == false {
			r += string(s[i])
		}
	}
	return strings.ToLower(strings.Trim(r, " "))
}

func (i *Inverted) GetDocs(gram string) []Doc {
	docs := []Doc{}

	l := i.index[gram]
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
	return math.Log(float64(i.NumDocuments) / float64(i.count[n]+1))
}

func (i *Inverted) StringsToGrams(t map[string]uint) *map[string]uint {
	nt := make(map[string]uint)
	for k, v := range t {
		nt[k] = v
	}
	return &nt
}

func (i *Inverted) AddDocument(t Doc) {
	doc := t.Tokens()
	for k, v := range doc {
		i.NumTokens += uint64(v)
		i.count[k] += uint64(v)
		if _, ok := i.index[k]; ok == false {
			i.index[k] = list.New()
		}
		i.index[k].PushBack(t)
	}
	i.NumDocuments++
}

// TODO: Update word count to reflect removed documents?
func (i *Inverted) RemoveDocument(t Doc) {
	for k, _ := range t.Tokens() {
		for e := i.index[k].Front(); e != nil; e = e.Next() {
			if e.Value == t {
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

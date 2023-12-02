package utils

type Trie[V any] struct {
  HasVal   bool
  Val      V
  children map[rune]*Trie[V]
}

func NewTrie[V any]() *Trie[V] {
  return &Trie[V]{
    children: map[rune]*Trie[V]{},
  }
}

func (t *Trie[V]) Add(keys string, ki int, val V) {
  if ki == len(keys) {
    t.Val = val
    t.HasVal = true
    return
  }

  c, ok := t.children[rune(keys[ki])]
  if !ok {
    c = NewTrie[V]()
    t.children[rune(keys[ki])] = c
  }
  c.Add(keys, ki+1, val)
}

func (t *Trie[V]) Get(word string, wi int) (*Trie[V], bool) {
  if wi == len(word) || t.HasVal {
    return t, t.HasVal
  }

  c, ok := t.children[rune(word[wi])]
  if !ok {
    return nil, false
  }
  return c.Get(word, wi+1)
}

func (t *Trie[V]) GetReverse(word string, wi int) (*Trie[V], bool) {
  if wi == -1 || t.HasVal {
    return t, t.HasVal
  }

  c, ok := t.children[rune(word[wi])]
  if !ok {
    return nil, false
  }
  return c.GetReverse(word, wi-1)
}

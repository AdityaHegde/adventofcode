package utils

import "fmt"

type LinkedList[Value interface{}] struct {
  Root *LinkedListNode[Value]
  Tail *LinkedListNode[Value]
}

func NewLinkedList[Value interface{}]() *LinkedList[Value] {
  return &LinkedList[Value]{}
}

type LinkedListNode[Value interface{}] struct {
  Next  *LinkedListNode[Value]
  Prev  *LinkedListNode[Value]
  Value Value
}

func (ll *LinkedList[Value]) Push(value Value) {
  node := &LinkedListNode[Value]{
    Value: value,
  }
  ll.connectTail(node)
}

func (ll *LinkedList[Value]) Pop() *LinkedListNode[Value] {
  if ll.Tail == nil {
    return nil
  }
  node := ll.Tail
  ll.clipTail(ll.Tail)
  node.Prev = nil
  return node
}

func (ll *LinkedList[Value]) Unshift(value Value) *LinkedListNode[Value] {
  node := &LinkedListNode[Value]{
    Value: value,
  }
  if ll.Root == nil {
    ll.Root = node
    ll.Tail = node
  } else {
    node.Next = ll.Root
    ll.Root.Prev = node
    ll.Root = node
  }
  return node
}

func (ll *LinkedList[Value]) Shift() *LinkedListNode[Value] {
  if ll.Root == nil {
    return nil
  }
  node := ll.Root
  ll.Root = ll.Root.Next
  node.Next = nil
  if ll.Root != nil {
    ll.Root.Prev = nil
  } else {
    ll.Tail = nil
  }
  return node
}

func (ll *LinkedList[Value]) ShiftTail(toLinkedList *LinkedList[Value], count int) {
  if ll.Tail == nil {
    return
  }

  node := ll.Tail
  for i := 1; i < count; i++ {
    if node.Prev == nil {
      break
    }
    node = node.Prev
  }
  tail := ll.Tail
  ll.clipTail(node)
  toLinkedList.connectTail(node)
  toLinkedList.Tail = tail
}

func (ll *LinkedList[Value]) connectTail(node *LinkedListNode[Value]) {
  node.Prev = ll.Tail
  if ll.Root == nil {
    ll.Root = node
    ll.Tail = node
  } else {
    ll.Tail.Next = node
    ll.Tail = node
  }
}

func (ll *LinkedList[Value]) clipTail(node *LinkedListNode[Value]) {
  ll.Tail = node.Prev
  if ll.Tail != nil {
    ll.Tail.Next = nil
  } else {
    ll.Root = nil
  }
}

func (ll *LinkedList[Value]) Print() {
  for c := ll.Root; c != nil; c = c.Next {
    fmt.Print(c.Value)
    fmt.Print(" ")
  }
  fmt.Println()
}

func (ll *LinkedList[Value]) Empty() bool {
  return ll.Root == nil
}

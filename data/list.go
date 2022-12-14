package data

type (
	ListNode ListNodeOf[any]

	ListNodeOf[T any] struct {
		Previous, Next *ListNodeOf[T]

		Value T
	}

	List ListOf[any]

	ListOf[T any] struct {
		Head, Tail *ListNodeOf[T]
	}
)

func NewList() (l *List) {
	l = (*List)(NewListOf[any]())
	return
}

func NewListOf[T any]() (l *ListOf[T]) {
	l = &ListOf[T]{}
	return
}

func PushHead(l *List, object any) {
	PushHeadOf((*ListOf[any])(l), object)
}

func PushHeadOf[T any](l *ListOf[T], object T) {
	head := l.Head

	node := &ListNodeOf[T]{
		Next:  head,
		Value: object,
	}

	if head != nil {
		head.Previous = node
	} else {
		l.Tail = node
	}

	l.Head = node
}

func PushTail(l *List, object any) {
	PushTailOf((*ListOf[any])(l), object)
}

func PushTailOf[T any](l *ListOf[T], object T) {
	tail := l.Tail

	node := &ListNodeOf[T]{
		Previous: tail,
		Value:    object,
	}

	if tail != nil {
		tail.Next = node
	} else {
		l.Head = node
	}

	l.Tail = node
}

func PopHead(l *List) (node *ListNode, ok bool) {
	n, ok := PopHeadOf((*ListOf[any])(l))
	node = (*ListNode)(n)
	return
}

func PopHeadOf[T any](l *ListOf[T]) (node *ListNodeOf[T], ok bool) {
	node = l.Head

	if node == nil {
		return
	}

	l.Head = node.Next
	ok = true
	return
}

func PopTail(l *List) (node *ListNode, ok bool) {
	n, ok := PopTailOf((*ListOf[any])(l))
	node = (*ListNode)(n)
	return
}

func PopTailOf[T any](l *ListOf[T]) (node *ListNodeOf[T], ok bool) {
	node = l.Tail

	if node == nil {
		return
	}

	l.Tail = node.Previous
	ok = true
	return
}

func Remove(l *List, node *ListNode) {
	RemoveOf((*ListOf[any])(l), (*ListNodeOf[any])(node))
}

func RemoveOf[T any](l *ListOf[T], node *ListNodeOf[T]) {
	if node.Next != nil {
		node.Next.Previous = node.Previous
	} else {
		l.Tail = node.Previous
	}

	if node.Previous != nil {
		node.Previous.Next = node.Next
	} else {
		l.Head = node.Next
	}
}

func Each(l *List, action func(object any)) {
	EachOf((*ListOf[any])(l), action)
}

func EachOf[T any](l *ListOf[T], action func(object T)) {
	node := l.Head

	for node != nil {
		action(node.Value)

		node = node.Next
	}
}

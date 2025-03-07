package structs

func insert[T comparable](orderedHead *Node[T], currentNode *Node[T], less func(contentA T, contentB T) bool) *Node[T] {
	var orderedPrevNode *Node[T] = nil
	orderedCurrentNode := orderedHead
	for orderedCurrentNode != nil {
		if !less(orderedCurrentNode.content, currentNode.content) {
			if orderedPrevNode == nil {
				currentNode.next = orderedHead
				orderedHead = currentNode
			} else {
				orderedPrevNode.next = currentNode
				currentNode.next = orderedCurrentNode
			}
			return orderedHead
		} else {
			orderedPrevNode = orderedCurrentNode
			orderedCurrentNode = orderedCurrentNode.next
		}
	}

	orderedPrevNode.next = currentNode
	currentNode.next = nil
	return orderedHead
}

func insertionSort[T comparable](head *Node[T], less func(contentA T, contentB T) bool) *Node[T] {
	orderedHead := head
	if head != nil {
		currentNode := orderedHead.next
		orderedHead.next = nil

		for currentNode != nil {
			tmpNode := currentNode
			currentNode = currentNode.next
			orderedHead = insert[T](orderedHead, tmpNode, less)
		}
	}
	return orderedHead
}

func merge[T comparable](left *Node[T], right *Node[T], less func(contentA T, contentB T) bool) *Node[T] {
	var head, tail *Node[T]
	for left != nil && right != nil {
		if less(left.content, right.content) {
			if head == nil {
				head = left
				tail = left
			} else {
				tail.next = left
				tail = left
			}
			left = left.next
		} else {
			if head == nil {
				head = right
				tail = right
			} else {
				tail.next = right
				tail = right
			}
			right = right.next
		}
	}

	if left != nil {
		tail.next = left
	} else if right != nil {
		tail.next = right
	}

	return head
}

func mergeSort[T comparable](head *Node[T], len uint32, less func(contentA T, contentB T) bool, insertion bool) *Node[T] {
	if head == nil || head.next == nil {
		return head
	}

	middle := head
	i := uint32(1)
	for ; i < len/2; i++ {
		middle = middle.next
	}

	secondHalf := middle.next
	middle.next = nil

	if len < 100 && insertion {
		return insertionSort(head, less)
	}

	left := mergeSort(head, i, less, insertion)
	right := mergeSort(secondHalf, len-i, less, insertion)

	return merge(left, right, less)
}

func makeMergeSortWorker[T comparable](head *Node[T], len uint32, less func(contentA T, contentB T) bool, insertion bool) chan *Node[T] {
	ch := make(chan *Node[T])
	go func() {
		ch <- mergeSortParallel(head, len, less, insertion)
	}()
	return ch
}

func mergeSortParallel[T comparable](head *Node[T], len uint32, less func(contentA T, contentB T) bool, insertion bool) *Node[T] {
	if head == nil || head.next == nil {
		return head
	}

	middle := head
	i := uint32(1)
	for ; i < len/2; i++ {
		middle = middle.next
	}

	secondHalf := middle.next
	middle.next = nil

	var left, right *Node[T]
	if len > 2000 {
		rightCh := makeMergeSortWorker(secondHalf, len-i, less, insertion)
		left = mergeSortParallel(head, i, less, insertion)
		right = <-rightCh
	} else if len < 100 && insertion {
		return insertionSort(head, less)
	} else {
		left = mergeSort(head, i, less, insertion)
		right = mergeSort(secondHalf, len-i, less, insertion)
	}

	return merge(left, right, less)
}

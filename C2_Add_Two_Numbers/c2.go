package c2

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToList(nums []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = nil

	for _, num := range nums {
		newNode := &ListNode{Val: num}
		if head == nil {
			head = newNode
			curr = newNode
		} else {
			curr.Next = newNode
			curr = newNode
		}
	}
	return head
}

// func printList(head *ListNode) {
// 	if head == nil {
// 		fmt.Println("nil")
// 		return
// 	}
// 	for head != nil {
// 		fmt.Printf("%d", head.Val)
// 		head = head.Next
// 		if head != nil {
// 			fmt.Print(" -> ")
// 		} else {
// 			fmt.Println()
// 		}
// 	}
// }

func listToArray(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	sum := 0
	for node := dummy; l1 != nil || l2 != nil || sum > 0; node = node.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummy.Next
}

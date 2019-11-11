package myLeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var firstNode *ListNode = new(ListNode)
	var currentNode *ListNode = firstNode

	// 1.只疊加兩串數字
	for {
		if l1 != nil && l2 != nil {
			//兩個都不為空

			//新增節點並跳過去
			currentNode.Next = new(ListNode)
			currentNode = currentNode.Next

			currentNode.Val += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			//兩個都為空
			break
		} else {
			//只有一個空
			//選出不為空的那一個點
			//接上現成且剩餘的link
			if l1 != nil {
				currentNode.Next = l1
			} else {
				currentNode.Next = l2
			}

			break
		}
	}

	// 2.單獨處理進位
	currentNode = firstNode.Next
	for nil != currentNode {
		//走訪所有點
		if currentNode.Val >= 10 {
			currentNode.Val = currentNode.Val - 10
			if nil == currentNode.Next {
				currentNode.Next = new(ListNode)
			}
			currentNode.Next.Val++
		}
		currentNode = currentNode.Next
	}

	return firstNode.Next
}

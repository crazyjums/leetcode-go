package main

type ListNode struct {
	Var  int
	Next *ListNode
}

func GetInterSectionNode (headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	//var pA *ListNode
	//var pB *ListNode
	//pA = headA
	//pB = headB
	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
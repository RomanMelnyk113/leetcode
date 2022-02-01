package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeLinked(head *ListNode) bool {
	cur := head

	original := []int{}
	for {
		original = append(original, cur.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}

	for i, j := 0, len(original)-1; i < len(original); i, j = i+1, j-1 {
		if original[i] != original[j] {
			return false
		}
	}
	return true
}

//func main() {
//p4 := ListNode{1, nil}
//p3 := ListNode{3, &p4}
//p2 := ListNode{2, &p3}
//p1 := ListNode{1, &p2}
//isP := isPalindromeLinked(&p1)
//fmt.Println(isP)
//}

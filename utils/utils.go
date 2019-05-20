package utils

type ListNode struct {
	Var  int
	Next *ListNode
}

func GenerateLinkList(arr []int) *ListNode {
	var link *ListNode
	var head *ListNode
	for v := range arr {
		l := new(ListNode)
		l.Var = arr[v]
		if link == nil {
			link = l
			head = link
		} else {
			link.Next = l
			link = l
		}
	}
	return head
}

func EqualStrings(l, r []string) bool {
	if len(l) != len(r) {
		return false
	}
	i := 0
	for i < len(l) {
		if l[i] != r[i] {
			return false
		}
		i++
	}
	return true
}

func CheckEqual(a,b [][]int)bool{
	if len(a) != len(b){
		return false
	}
	for i,v := range a{
		if checkEqualArr(v,b[i]) == false{
			return false
		}
	}
	return true
}

func checkEqualArr(a,b []int)bool{
	if len(a) != len(b){
		return false
	}
	for i,v := range a{
		if b[i] != v{
			return false
		}
	}
	return true
}
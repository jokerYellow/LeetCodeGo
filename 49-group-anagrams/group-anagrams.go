package leetcode
import "sort"

type sortBytes []byte 

func (s sortBytes) Less(i,j int)bool{
	return s[i] < s[j]
}

func (s sortBytes)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func (s sortBytes) Len() int{
	return len(s)
} 

func groupAnagrams(strs []string)[][]string{
	m := map[string][]string{}
	for _,item := range strs{
		b := []byte(item)
		sort.Sort(sortBytes(b))
		s := string(b)
		if v,ok := m[s];ok{
			m[s] = append(v,item)
		}else{
			m[s] = []string{item}
		}
	}
	var values [][]string
	for _,v := range m {
		values = append(values,v)
	}
	return values
}
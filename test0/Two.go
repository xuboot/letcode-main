package main

func twosum(nums []int,target int) []int  {
	m := make(map[int]int)
	for i:=0; i < len(nums) ;i ++ {
		tmp := target - nums[i]
		if _,ok := m[tmp];ok{
			return []int{m[tmp],i}
		}
		m[nums[i]] = i
	}
	return nil
}


func main()  {
	twosum([]int{1, 2, 3,4,5 ,6,7,8},3)
}

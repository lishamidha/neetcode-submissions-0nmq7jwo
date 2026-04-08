func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for i,v:= range nums{
        find:=target-v
        if idx,ok:=m[find];ok{
            return []int{idx,i}
        }
        m[v]=i
    }
    return nil
}

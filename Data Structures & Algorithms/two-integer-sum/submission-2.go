func twoSum(nums []int, target int) []int {
    map1:=make(map[int]int)
    for i,v:= range nums{
        find:=target-v
        if idx, found:=map1[find]; found {
            return []int{idx,i}
        }
        map1[v]=i
    }
    return nil
}

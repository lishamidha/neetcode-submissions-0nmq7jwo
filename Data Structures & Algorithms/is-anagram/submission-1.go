func isAnagram(s string, t string) bool {
    freq:=make(map[rune]int)
    if len(s)!=len(t){
        return false
    }
    for _,v:=range s{
        freq[v]++
    }
    for _,v:=range t{
        freq[v]--
        if freq[v]<0{
            return false
        }
    }
    return true
}

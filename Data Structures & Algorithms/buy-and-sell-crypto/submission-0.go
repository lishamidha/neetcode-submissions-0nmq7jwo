func maxProfit(prices []int) int {
   // map1:=make(map[int]int)
    l:=len(prices)
    profit:=0
    for i:=0;i<l;i++{
        buy:=prices[i]
        for j:=i+1;j<l;j++{
            sell:=prices[j]
            profit=max(profit,sell-buy)
        }
    }   
    return profit
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

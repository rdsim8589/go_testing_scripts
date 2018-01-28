package bestSellBuyStock

func MaxProfit(prices []int) int {
	profit := 0
	ownStock := 0
	for i := 0; i < len(prices)-1; i++ {
		if ownStock == 0 {
			if prices[i] < prices[i+1] {
				profit -= prices[i]
				ownStock = 1
			}
		}
		if ownStock == 1 {
			if prices[i] > prices[i+1] {
				profit += prices[i]
				ownStock = 0
			}
		}
	}
	if ownStock == 1 {
		profit += prices[len(prices)-1]
		ownStock = 0
	}
	return profit
}

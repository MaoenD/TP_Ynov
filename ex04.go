package tp

func Ft_profit(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        }
        profit := price - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}

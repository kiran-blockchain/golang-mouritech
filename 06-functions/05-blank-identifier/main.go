package main

import (
	"fmt"
	"math"
)

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64,error) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange,nil
}

func main() {
	prevStockPrice := 80000.0
	currentStockPrice := 120000.0

	change, _ ,err:= getStockPriceChange(prevStockPrice, currentStockPrice)
	if err!=nil{
		fmt.Println(err);
	}

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}

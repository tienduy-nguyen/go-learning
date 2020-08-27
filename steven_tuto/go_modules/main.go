package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"go-modules/c4f"
	"math/big"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))                       // "$123,456,789.21"
	fmt.Println(ac.FormatMoney(12345678))                               // "$12,345,678.00"
	fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))                // "$25,925,925.67"
	fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))               // "-$25,925,925.67"
	fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123))) // "$123,456,789.21"

	ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	fmt.Println(ac.FormatMoney(4999.99)) // "€4.999,99"
	c4f.Print("This is red color")
}

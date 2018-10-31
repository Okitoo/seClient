package main

import (
	"fmt"
	"github.com/Okitoo/seClient"
)

var bot *seClient.Client

func main() {
	//create a new bot
	bot = seClient.New("https://start-ex.com/api", "", "")

	//get all markets
	markets, err := bot.Public_ListMarkets()
	if err != nil {
		println("Error listing markets: ", err.Error())
		return
	}
	println("Total Markets: ", len(markets))

	// list all prices in the market
	for _, market := range markets {
		orders, _ := bot.Public_GetMarketOrders(market.ID, 10)
		if len(orders.Bids) > 0 {
			fmt.Printf("Buy Price: %s, %0.8f, %s\n", market.AltCoin, float64(orders.Bids[0].Price)/100000000, market.BaseCoin)
		}
		if len(orders.Asks) > 0 {
			fmt.Printf("Sell Price: %s, %0.8f, %s\n", market.AltCoin, float64(orders.Asks[0].Price)/100000000, market.BaseCoin)
		}
	}

	/*
		A customer walks in and wants to buy X42, he has bitcoins
		You want to charge him a 1% conversion fee from your website
	*/

	customer_x42_withdraw_address_x42 := "customer_address"
	your_1_percent_profit_address_x42 := "your_x42_address"

	//tell instaex to send 99% of converted coins to the customer and 1% of the converted coins to you
	instaexWithdraw := customer_x42_withdraw_address_x42 + " 99\n" + your_1_percent_profit_address_x42 + " 1\n"

	//create the insta-ex
	ix_info, err := bot.Public_CreateInstaex("BTC", "X42", instaexWithdraw)
	if err != nil {
		println("Error creating instaex markets: ", err.Error())
		return
	}

	fmt.Printf("Deposit minimum of %0.8f BTC to %s", float64(ix_info.Info.MinDeposit)/100000000, ix_info.Data.DepositAddress)

}

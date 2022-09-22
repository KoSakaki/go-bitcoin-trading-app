package main

import (
	"gotrading/app/controllers"
	_ "gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)

	// fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
	// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)

	// order := &bitflyer.Order{
	// 	ProductCode:     config.Config.ProductCode,
	// 	ChildOrderType:  "MARKET",
	// 	Side:            "BUY",
	// 	Size:            0.001,
	// 	MinuteToExpires: 1,
	// 	TimeInForce:     "GTC",
	// }
	// res, err := apiClient.SendOrder(order)
	// if err != nil {
	// 	log.Println("Error: ", err)
	// }
	// if res.Status != 200 {
	// 	log.Printf("status : %v, Error: %s", res.Status, res.ErrorMessage)

	// }

	// i := res.ChildOrderAcceptanceID
	// params := map[string]string{
	// 	"product_code":              config.Config.ProductCode,
	// 	"child_order_acceptance_id": i,
	// }
	// r, _ := apiClient.ListOrder(params)
	// fmt.Println(r)

}

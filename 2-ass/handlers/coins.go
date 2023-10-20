package handlers

import (
	"2-ass/database"
	"2-ass/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCoins(ctx *gin.Context) {
	var coins []models.Coin
	res := database.GetDB().Find(&coins)
	if res.Error != nil {
		panic(res.Error)
	}
	ctx.IndentedJSON(http.StatusOK, coins)
}
func GetCoin(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	var item models.Coin
	res := database.GetDB().First(&item, "id = ?", id)
	//res := database.GetDB().First(&item, id)
	if res.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, item)
}

func InsertData() {
	data := GetData()
	for _, item := range data {
		result := database.GetDB().Exec(`
			INSERT INTO coins (id, symbol, name, current_price)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO UPDATE
			SET symbol = EXCLUDED.symbol, name = EXCLUDED.name, current_price = EXCLUDED.current_price
		`, item.ID, item.Symbol, item.Name, item.CurrentPrice)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}

func GetData() []models.Coin {
	apiURL := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	response, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var data []models.Coin

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return data
}

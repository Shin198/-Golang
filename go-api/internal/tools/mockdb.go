package tools

import (
	"time"
)

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"Thanh Tu": {
		AuthToken: "123qwe",
		Username:  "ThanhTu",
	},
	"Thanh Xinh": {
		AuthToken: "061201",
		Username:  "Thanh Xinh",
	},
	"Shin": {
		AuthToken: "Shin98",
		Username:  "Shin",
	},
}
var mockCoinDetails = map[string]CoinDetails{
	"Thanh Tu": {
		Coins:    150,
		Username: "ThanhTu",
	},
	"Thanh Xinh": {
		Coins:    350,
		Username: "Thanh Xinh",
	},
	"Shin": {
		Coins:    950,
		Username: "Shin",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

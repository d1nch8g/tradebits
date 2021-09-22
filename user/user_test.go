package user

import (
	"reflect"
	"sync_tree/calc"
	"sync_tree/data"
	"sync_tree/trade"
	"testing"
	"time"
)

var dummyMesKey = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}
var dummyName = "Name"

func TestCreateUser(t *testing.T) {
	var adress = calc.Rand()
	err := Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	if err != nil {
		t.Error("attemt to create new user failed")
	}
	data.TestRM(adress)
}

func TestCreateExisting(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	err := Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	if err == nil {
		t.Error("attemt to create existing user succeded, that is bad error")
	}
	data.TestRM(adress)
}

func TestGetFreeUser(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	freeUser := Get(adress)
	freeUser.Save()
	if freeUser.PublicName != "Name" {
		t.Error("get free user error")
	}
	data.TestRM(adress)
}

var usr2 *user

func getBusyUser(adress []byte) {
	usr2 = Get(adress)
}

func TestGetBusyUser(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr1 := Get(adress)
	go getBusyUser(adress)
	time.Sleep(time.Second)
	usr1.Save()
	time.Sleep(time.Second)
	if !reflect.DeepEqual(usr2.adress, adress) {
		t.Error("adress of second user should be the same")
	}
	data.TestRM(adress)
}

func TestUserLook(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Look(adress)
	if len(usr.adress) != 0 {
		t.Error("user adress should be empty")
	}
	if usr.PublicName != dummyName {
		t.Error("user info is incorrect")
	}
	data.TestRM(adress)
}

func TestPutUserMessage(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.PutUserMessage([]byte{1, 2, 3}, "message")
	usr.Save()
	mes := Look(adress).GetMessages([]byte{1, 2, 3})[0]
	if mes != "umessage" {
		t.Error("the message should be 'message' - " + mes)
	}
	data.TestRM(adress)
}

func TestPutMarketMessage(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.PutMarketMessage([]byte{1, 2, 3}, "message")
	usr.Save()
	mes := Look(adress).GetMessages([]byte{1, 2, 3})[0]
	if mes != "mmessage" {
		t.Error("the message should be 'message' - " + mes)
	}
	data.TestRM(adress)
}

func TestNewUserNonNullableMessageMap(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	if usr.Messages == nil {
		t.Error("user messages should never be null")
	}
	usr.Save()
	data.TestRM(adress)
}

func TestAttachToLookedUser(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Look(adress)
	buy := trade.Buy{}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("buy should not be attached to user, that can never be saved")
	}
	sell := trade.Sell{}
	sellAttached := usr.AttachSell(&sell, []byte{})
	if sellAttached {
		t.Error("sell trade should not be attached, cuz user can never be saved")
	}
	data.TestRM(adress)
}

func TestAttachTradesWithZeroOffer(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   0,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("this buy should not be attached cuz hase zero offer")
	}
	sell := trade.Sell{
		Offer:   0,
		Recieve: 100,
	}
	sellAttached := usr.AttachSell(&sell, []byte{})
	if sellAttached {
		t.Error("this sell should never be attached cuz 0 offer")
	}
	data.TestRM(adress)
}

func TestAttachTradeWithBigBalance(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if buyAttached {
		t.Error("this buy should not be attached cuz its over users balance")
	}
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	usr.Balances["x"] = 0
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if sellAttached {
		t.Error("this sell should not be attached cuz its over users balance")
	}
	data.TestRM(adress)
}

func TestAttachNormalTrades(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.Balance = 1000
	buy := trade.Buy{
		Offer:   1000,
		Recieve: 1000,
	}
	buyAttached := usr.AttachBuy(&buy)
	if !buyAttached {
		t.Error("this buy should be attached normally")
	}
	if !reflect.DeepEqual(buy.Adress, usr.adress) {
		t.Error("buy adress after bounding should be equal to users")
	}
	usr.Balances["x"] = 1000
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if !sellAttached {
		t.Error("this sell should be attached normally")
	}
	if !reflect.DeepEqual(sell.Adress, usr.adress) {
		t.Error("sell adress after bounding should be equal to users")
	}
	data.TestRM(adress)
}

func TestAttachSellNonExistingMarket(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	sell := trade.Sell{
		Offer:   1000,
		Recieve: 1000,
	}
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if sellAttached {
		t.Error("this sell should not be attached, cuz user dont have such market")
	}
	data.TestRM(adress)
}

func TestAttchBoundedTrades(t *testing.T) {
	var adress = calc.Rand()
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	usr := Get(adress)
	usr.Balance = 100
	usr.Balances["x"] = 100
	buy := trade.Buy{
		Offer:   10,
		Recieve: 10,
		Adress:  []byte{0},
	}
	sell := trade.Sell{
		Offer:   10,
		Recieve: 10,
		Adress:  []byte{0},
	}
	buyAttached := usr.AttachBuy(&buy)
	sellAttached := usr.AttachSell(&sell, []byte("x"))
	if buyAttached || sellAttached {
		t.Error("those trades are already bounded and should not be attached")
	}
	data.TestRM(adress)
}

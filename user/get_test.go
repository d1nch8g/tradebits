package user

import (
	"reflect"
	"sync_tree/calc"
	"sync_tree/data"
	"testing"
	"time"
)

func TestGetFreeUser(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
	Create(
		adress,
		dummyMesKey,
		dummyName,
	)
	freeUser := Get(adress)
	freeUser.Save()
	if reflect.DeepEqual(adress, freeUser.adress) {
		t.Error("get free user error")
	}
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}

var usr2 *user

func getBusyUser(adress []byte) {
	usr2 = Get(adress)
}

func TestGetBusyUser(t *testing.T) {
	var adress = calc.Rand()
	dummyName := string(calc.Rand()[0:8])
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
	data.TestRM([]byte(dummyName))
	data.TestRM(adress)
}
package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"reflect"
	"sync_tree/data"
	"sync_tree/lock"
	"sync_tree/trade"
	"time"
)

type user struct {
	adress     []byte
	Balance    uint64
	MesKey     []byte
	PublicName string
	Markets    map[string]uint64
	Mes        map[string]string
	Arch       map[string]string
}

/*
Create new user, in case there is already user with same adress
the error will be logged
*/
func Create(adress []byte, MesKey []byte, PublicName string) error {
	if data.Check(adress) {
		return errors.New("possibly user already exists")
	}
	u := user{
		Balance:    0,
		MesKey:     MesKey,
		PublicName: PublicName,
		Markets:    make(map[string]uint64),
		Mes:        make(map[string]string),
		Arch:       map[string]string{},
	}
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	data.Put(adress, cache.Bytes())
	return nil
}

/*
Get existing user from database, by getting user his ID is gonna be
locked, so another of that user are not gonna appear
*/
func Get(adress []byte) *user {
	lockErr := lock.Lock(adress)
	if lockErr != nil {
		time.Sleep(time.Millisecond * 89)
		return Get(adress)
	}
	u := user{adress: adress}
	userBytes := data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}

/*
Non blocking function to look for user contents, it's impossible to save
instance of that user to database.
*/
func Look(adress []byte) *user {
	u := user{}
	userBytes := data.Get(adress)
	cache := bytes.NewBuffer(userBytes)
	gob.NewDecoder(cache).Decode(&u)
	return &u
}

/*
This function is used to save user, after that function his state is
gonna be fixed in database, adress will be unlocked, and adress will be
set to nil, so other changes won't be saved (user will have to be recreated)
*/
func (u *user) Save() {
	saveAdress := u.adress
	u.adress = nil
	cache := new(bytes.Buffer)
	gob.NewEncoder(cache).Encode(u)
	data.Change(saveAdress, cache.Bytes())
	lock.Unlock(saveAdress)
}

/*
Function to add message from some adress to concrete user
*/
func (u *user) PutMessage(userAdress []byte, mes string) {
	strAdr := string(userAdress)
	u.Mes[strAdr] = u.Mes[strAdr] + "|" + mes
}

/*
This function is made to get all new messages and to put all current messages
to archieve
*/
func (u *user) GetAllMessages() map[string]string {
	messages := u.Mes
	for sender, message := range u.Mes {
		u.Arch[sender] = u.Arch[sender] + "|" + message
	}
	u.Mes = make(map[string]string)
	return messages
}

// this funciton checks wether it is possible to generate some trade for user,
// if it is possible, this lowers users balance, and returns nil
func (u *user) AttachBuy(trade *trade.Buy) error {
	if u.adress == nil {
		return errors.New("this user could never be saved, get user with get function instead of look")
	}
	if trade.Offer == 0 || trade.Recieve == 0 {
		return errors.New("offer/recieve should not be equal to zero")
	}
	if trade.Offer > u.Balance {
		return errors.New("trade offer is bigger than users balance, this trade could not be created")
	}
	trade.Attached = true
	u.Balance = u.Balance - trade.Offer
	return nil
}

func (u *user) AttachSell(trade *trade.Sell) error {
	if u.adress == nil {
		return errors.New("this user could never be saved, get user with get function instead of look")
	}
	if trade.Offer == 0 || trade.Recieve == 0 {
		return errors.New("offer/recieve should not be equal to zero")
	}
	if val, ok := u.Markets[string(trade.Adress)]; ok {
		if trade.Offer > val {
			return errors.New("trade offer is bigger than users balance, this trade could not be created")
		}
		trade.Attached = true
		u.Markets[string(trade.Adress)] = val - trade.Offer
		return nil
	}
	return errors.New("this market does not exist for specific user")
}

func (u *user) UnboundBuy(trade *trade.Buy) error {
	if !reflect.DeepEqual(trade.Adress, u.adress) {
		return errors.New("adress of user and trade are not matching")
	}
	u.Balance = u.Balance + trade.Offer
	return nil
}

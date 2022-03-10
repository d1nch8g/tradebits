package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Key      string   `bson:"ukey"`
	Balance  int      `bson:"balance"`
	Messages []string `bson:"messages"`
}

type CreateRequest struct {
	Hkey string `json:"hkey"`
	Ukey string `json:"ukey"`
	Sign string `json:"sign"`
}

func UserCreatePut(w http.ResponseWriter, r *http.Request) {
	req := CreateRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	if req.Hkey == "" || req.Ukey == "" || req.Sign == "" {
		w.WriteHeader(406)
		return
	}
	if req.Hkey != crypt.Pub() {
		w.WriteHeader(421)
		return
	}
	verfied := crypt.Verify(req.Hkey+req.Ukey, req.Ukey, req.Sign)
	if !verfied {
		w.WriteHeader(401)
		return
	}
	exists := mongo.Check("user", "ukey", req.Ukey)
	if exists {
		w.WriteHeader(403)
		return
	}
	mongo.Put("user", User{
		Key:      req.Ukey,
		Balance:  0,
		Messages: []string{},
	})
	w.WriteHeader(201)
}

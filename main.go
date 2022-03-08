package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tradebits/api"
	"tradebits/crypter"
	"tradebits/graylog"
	"tradebits/mongoer"
	"tradebits/rediser"

	"github.com/joho/godotenv"
)

func readConfigField(field string) string {
	envar, found := os.LookupEnv(field)
	if !found {
		log.Fatal("problem loading .ENV field: ", envar)
	}
	return envar
}

func init() {
	godotenv.Load()
	gerr := graylog.Setup(readConfigField("GRAYLOG_API"))
	if gerr != nil {
		log.Fatal("Unable to setup graylog", gerr)
	}
	crypt, err1 := crypter.Get(readConfigField("MARKET_PRIVATEKEY"))
	redis, err2 := rediser.Get(readConfigField("REDIS_HOST"))
	mongo, err3 := mongoer.Get(readConfigField("MONGO_HOST"))
	m := map[string]string{
		"name":      readConfigField("MARKET_NAME"),
		"mkey":      crypt.Pub(),
		"descr":     readConfigField("MARKET_DESCR"),
		"img":       readConfigField("MARKET_IMG"),
		"worktime":  readConfigField("MARKET_WORKTIME"),
		"fee":       readConfigField("MARKET_FEE"),
		"delimiter": readConfigField("MARKET_DELIMITER"),
	}
	respbytes, err4 := json.Marshal(m)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatal("1) Setup err: ", err1, err2, err3, err4)
	}
	colErr1 := mongo.CreateCollection("user")
	colErr2 := mongo.CreateCollection("net")
	colErr3 := mongo.CreateCollection("trades")
	if colErr1 != nil || colErr2 != nil || colErr3 != nil {
		log.Fatal("2) Setup err: ", colErr1, colErr2, colErr3)
	}
	idxErr1 := mongo.CreateIndex("user", "key", "hashed")
	idxErr2 := mongo.CreateIndex("trades", "ukey", "hashed")
	idxErr3 := mongo.CreateIndex("trades", "mkey", "hashed")
	if idxErr1 != nil || idxErr2 != nil || idxErr3 != nil {
		log.Fatal("3) Setup err: ", idxErr1, idxErr2, idxErr3)
	}
	api.Setup(respbytes, mongo, crypt, redis)
	log.Println("Setup sucess...")
}

func main() {
	log.Printf("Server starting...")
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(readConfigField("MARKET_PORT"), router))
}

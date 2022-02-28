/*
 * TradeBits
 *
 * In this API description you can find description of ways markets communicate with each other and users.   Each sign of message user send to market is additionally reinforced by market name as indtent to ensure markets will not try to proceed same operations with users on other markets. That is not mention in API, but implemented on server and client sides.   Each procedure on specific market need to be reinforced by market name in senders message, to prevent mimic behavior by markets collecting data. That will guarantee, that only owner of the private key is able to process transactions on specific market.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
)

func MarketCloseDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func MarketCreatePut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func MarketDecreasePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func MarketRelatedGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

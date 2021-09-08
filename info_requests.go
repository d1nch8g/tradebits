package main

import (
	"context"
	"errors"
	"fmt"
	"sync_tree/market"
	"sync_tree/search"
	"sync_tree/user"

	pb "sync_tree/api"
)

func (s *server) InfoHasTrades(
	ctx context.Context,
	in *pb.InfoHasTradesRequest,
) (*pb.Response, error) {
	user := user.Look(in.UserAdress)
	if user != nil {
		market := market.Look(in.MarketAdress)
		if market != nil {
			has := market.HasTrades(in.UserAdress)
			return &pb.Response{Passed: has}, nil
		}
	}
	return &pb.Response{Passed: false}, nil
}

func (s *server) InfoMarket(
	ctx context.Context,
	in *pb.InfoMarketRequest,
) (*pb.InfoMarketResponse, error) {
	m := market.Look(in.Adress)
	if m != nil {
		fmt.Println(m.Name)
		for idx, buy := range m.Pool.Buys {
			fmt.Println("buy", idx, "offer", buy.Offer, "recieve", buy.Recieve)
		}
		for idx, sell := range m.Pool.Sells {
			fmt.Println("sell", idx, "offer", sell.Offer, "recieve", sell.Recieve)
		}
		return &pb.InfoMarketResponse{
			MesKey:  m.MesKey,
			Name:    m.Name,
			Img:     m.Img,
			Descr:   m.Descr,
			OpCount: m.OpCount,
		}, nil
	}
	return &pb.InfoMarketResponse{}, errors.New("market not found")
}

func (s *server) InfoSearch(
	ctx context.Context,
	in *pb.InfoSearchRequest,
) (*pb.InfoSearchResponse, error) {
	//fmt.Println("user made a search request on: ", in.Info)
	results := search.Search(in.Info)
	return &pb.InfoSearchResponse{ConcMarkets: results}, nil
}

func (s *server) InfoUser(
	ctx context.Context,
	in *pb.InfoUserRequest,
) (*pb.InfoUserResponse, error) {
	//fmt.Println("giving information about", in.Adress)
	user := user.Look(in.Adress)
	fmt.Println(user)
	if user == nil {
		return &pb.InfoUserResponse{
			PublicName: "====",
			Balance:    0,
		}, nil
	}
	adressesSlice := [][]byte{}
	balancesSlice := []uint64{}
	for strAdr, bal := range user.Balances {
		adressesSlice = append(adressesSlice, []byte(strAdr))
		balancesSlice = append(balancesSlice, bal)
	}
	return &pb.InfoUserResponse{
		PublicName:     user.PublicName,
		Balance:        user.Balance,
		MesKey:         user.MesKey,
		MarketAdresses: adressesSlice,
		MarketBalances: balancesSlice,
	}, nil
}
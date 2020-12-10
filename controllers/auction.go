package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"lelangbackend/helper"
	"lelangbackend/models"
	"log"
	"strconv"
	"time"
)

type AuctionController struct {
	beego.Controller
}

func (a *AuctionController) Add() {
	var auctionItem models.AuctionItem
	json.Unmarshal(a.Ctx.Input.RequestBody, &auctionItem)

	a.Ctx.Input.Bind(&auctionItem.Picture, "picture")
	a.Ctx.Input.Bind(&auctionItem.BuyNow, "buy_now")
	a.Ctx.Input.Bind(&auctionItem.BidMultiple, "bid_multiple")
	a.Ctx.Input.Bind(&auctionItem.StockCondition, "stock_condition")
	a.Ctx.Input.Bind(&auctionItem.StartBid, "start_bid")
	//a.Ctx.Input.Bind(&auctionItem.CurrentBid, "current_bid")
	a.Ctx.Input.Bind(&auctionItem.Description, "description")
	a.Ctx.Input.Bind(&auctionItem.IdAuctioneer, "id_auctioneer")
	finishDate, err := time.Parse(time.RFC3339, a.GetString("finish_date"))
	if err != nil {
		log.Print(err)
	}
	auctionItem.FinishDate = finishDate

	publishDate, er := time.Parse(time.RFC3339, a.GetString("publish_date"))
	if er != nil {
		log.Print(er)
	}
	auctionItem.PublishDate = publishDate

	rc, msg := models.AddAuction(auctionItem)
	helper.Response(rc, msg, nil, a.Controller)

	a.ServeJSON()
}

// @Title getAllAuction
// @Summary getAllAuction
// @Description get all auction
// @Success 200 {object} models.AuctionItem
// @Failure 403 body is empty
// @router /api/lelang [get]
func (a *AuctionController) Get() {
	rc, msg, data := models.GetAllAuction()

	helper.Response(rc, msg, data, a.Controller)
}

func (a *AuctionController) Bid() {
	var bid models.Bid
	json.Unmarshal(a.Ctx.Input.RequestBody, &bid)

	a.Ctx.Input.Bind(&bid.IdBidder, "bidder_id")
	a.Ctx.Input.Bind(&bid.IdAuctionItem, "auction_id")
	a.Ctx.Input.Bind(&bid.OfferBid, "offer_bid")

	bid.BidDate = time.Now()

	rc, msg := models.BidAuction(bid)
	helper.Response(rc, msg, nil, a.Controller)
}

func (a *AuctionController) GetAllBidder() {
	auctionID, err := strconv.Atoi(a.GetString(":auctionid"))
	if err != nil {
		log.Print("AUCTION ID ERROR " + err.Error())
	}

	rc, msg, data := models.GetAllBidder(auctionID)
	helper.Response(rc, msg, data, a.Controller)
}

func (a *AuctionController) ChooseBidder() {

}

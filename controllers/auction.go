package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"lelangbackend/helper"
	"lelangbackend/models"
	"log"
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
	helper.Response(rc, msg, a.Controller)

	a.ServeJSON()
}

func (a *AuctionController) Get() {
	rc, data := models.GetAllAuction()
	helper.Response(rc, data, a.Controller)

	a.ServeJSON()
}

func (a *AuctionController) Bid() {
	var bid models.Bid
	json.Unmarshal(a.Ctx.Input.RequestBody, &bid)

	a.Ctx.Input.Bind(&bid.IdBidder, "bidder_id")
	a.Ctx.Input.Bind(&bid.IdAuctionItem, "auction_id")
	a.Ctx.Input.Bind(&bid.OfferBid, "offer_bid")

	bid.BidDate = time.Now()

	rc, msg := models.BidAuction(bid)
	helper.Response(rc, msg, a.Controller)

	a.ServeJSON()
}

// @router /api/lelang/bid/:id [get]
func (a *AuctionController) GetAllBidder() {
	actionID := a.Ctx.Input.Param(":action_id")
	log.Print("ID : " + actionID)

	rc, data := models.GetAllBidder(actionID)
	helper.Response(rc, data, a.Controller)

	a.ServeJSON()
}

func (a *AuctionController) ChooseBidder() {

}

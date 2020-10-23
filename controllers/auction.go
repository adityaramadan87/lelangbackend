package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"lelangbackend/helper"
	"lelangbackend/models"
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
	a.Ctx.Input.Bind(&auctionItem.FinishDate, "finish_date")
	a.Ctx.Input.Bind(&auctionItem.PublishDate, "publish_date")

	rc, msg := models.AddAuction(auctionItem)
	helper.Response(rc, msg, a.Controller)

	a.ServeJSON()
}

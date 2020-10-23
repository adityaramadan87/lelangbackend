package models

import (
	"github.com/astaxie/beego/orm"
)

type AuctionItem struct {
	Id              int
	Picture         string
	BuyNow          int
	BidMultiple     int
	StockCondition  string
	StartBid        int
	CurrentBid      int
	Description     string
	IdAuctioneer    int
	FinishDate      string
	PublishDate     string
	IdChoosenBidder int
}

type Bid struct {
	Id            int
	IdAuctionItem int
	IdBidder      int
	OfferBid      int
	BidDate       string
}

func init() {
	orm.RegisterModel(new(AuctionItem), new(Bid))
}

func AddAuction(a AuctionItem) (rc int, msg string) {
	o := orm.NewOrm()
	o.Using("default")

	if a != (AuctionItem{}) {

		a.CurrentBid = 0

		var user S_user
		o.Raw("SELECT * FROM s_user WHERE id = ?", a.IdAuctioneer).QueryRow(&user)
		if user == (S_user{}) {
			return 1, "Auctioneer ID not found"
		}

		_, err := o.Insert(&a)
		if err != nil {
			return 1, "Error when insert data " + err.Error()
		}

		return 0, "Insert data success"

	}

	return 1, "Auction Item null"
}

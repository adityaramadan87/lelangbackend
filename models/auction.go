package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"time"
)

type AuctionItem struct {
	Id              int       `json:"id"`
	Picture         string    `json:"picture"`
	BuyNow          int       `json:"buy_now"`
	BidMultiple     int       `json:"bid_multiple"`
	StockCondition  string    `json:"stock_condition"`
	StartBid        int       `json:"start_bid"`
	CurrentBid      int       `json:"current_bid"`
	Description     string    `json:"description"`
	IdAuctioneer    int       `json:"id_auctioneer"`
	FinishDate      time.Time `json:"finish_date"`
	PublishDate     time.Time `json:"publish_date"`
	IdChoosenBidder int       `json:"id_choosen_bidder"`
}

type Bid struct {
	Id            int       `json:"id"`
	IdAuctionItem int       `json:"id_auction_item"`
	IdBidder      int       `json:"id_bidder"`
	OfferBid      int       `json:"offer_bid"`
	BidDate       time.Time `json:"bid_date"`
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

func GetAllAuction() (rc int, msg string, data []AuctionItem) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Raw(
		"SELECT * FROM auction_item WHERE id_choosen_bidder = 0 AND publish_date <= ? AND finish_date >= ?",
		time.Now(), time.Now()).QueryRows(&data)

	if err != nil {
		log.Print(err.Error())
		return 1, err.Error(), nil
	}

	return 0, "Success", data

}

func BidAuction(b Bid) (rc int, msg string) {
	o := orm.NewOrm()
	o.Using("default")

	var auctionItem AuctionItem

	if b != (Bid{}) {
		o.Raw("SELECT * FROM auction_item WHERE id = ?", b.IdAuctionItem).QueryRow(&auctionItem)
		if auctionItem == (AuctionItem{}) {
			return 1, "Auction item not found"
		}

		if auctionItem.CurrentBid == 0 {
			if auctionItem.StartBid != 0 {
				if b.OfferBid < auctionItem.StartBid {
					return 1, "Offer Bid must be greater or equal than Start Bid"
				}
				if b.OfferBid < (auctionItem.StartBid + auctionItem.BidMultiple) {
					return 1, "Offer Bid minimum is " + strconv.Itoa(auctionItem.StartBid+auctionItem.BidMultiple)
				}
			}
		} else {
			if b.OfferBid < auctionItem.CurrentBid {
				return 1, "Offer Bid must be greater or equal than Current Bid"
			}

			if b.OfferBid < (auctionItem.CurrentBid + auctionItem.BidMultiple) {
				return 1, "Offer Bid minimum is " + strconv.Itoa(auctionItem.CurrentBid+auctionItem.BidMultiple)
			}
		}

		auctionItem.CurrentBid = b.OfferBid

		if _, err := o.Update(&auctionItem, "current_bid"); err != nil {
			return 1, "Error while update " + err.Error()
		}

		if _, err := o.Insert(&b); err != nil {
			return 1, "Error while insert " + err.Error()
		}

		return 0, "Bid Success"

	}

	return 1, "Bid Data Null"
}

func GetAllBidder(actionID int) (rc int, msg string, data []Bid) {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Raw("SELECT * FROM bid WHERE id_auction_item = ? ORDER BY offer_bid DESC", actionID).QueryRows(&data)

	if err != nil {
		log.Print(err.Error())
		return 1, err.Error(), nil
	}

	return 0, "Success", data
}

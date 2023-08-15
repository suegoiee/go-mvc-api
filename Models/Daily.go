package Models

type Daily struct {
	Title string `json:"title" bson:"title"`
	Time  string `json:"time" bson:"time"`
	// OpeningPrice float32 `json:"oPrice" bson:"oPrice"`
	// TradeInfo    []Info  `json:"info" bson:"info"`
}

// type Info struct {
// 	Price     float32   `json:"price" bson:"price"`
// 	Amount    int32     `json:"amount" bson:"amount"`
// 	TradeTime time.Time `json:"time" bson:"time"`
// }

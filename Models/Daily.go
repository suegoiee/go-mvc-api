package Models

type Daily struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
	// OpeningPrice float32 `json:"oPrice" bson:"oPrice"`
	// TradeInfo    []Info  `json:"info" bson:"info"`
}

// type Info struct {
// 	Price     float32   `json:"price" bson:"price"`
// 	Amount    int32     `json:"amount" bson:"amount"`
// 	TradeTime time.Time `json:"time" bson:"time"`
// }

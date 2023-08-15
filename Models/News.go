package Models

type News struct {
	Title   string  `json:"title" bson:"title"`
	Time    string  `json:"time" bson:"time"`
	Content string  `json:"content" bson:"content"`
	Source  string  `json:"source" bson:"source"`
	Images  []Image `json:"image" bson:"image"`
	Link    string  `json:"link" bson:"link"`
}

type Image struct {
	Link        string `json:"link" bson:"link"`
	Description string `json:"description" bson:"description"`
}

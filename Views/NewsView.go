package Views

import (
	"ginPrac/Models"
)

type NewsView []News

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

func RenderNews(newsList []*Models.News) NewsView {
	view := NewsView{}
	for i := 0; i < len(newsList); i++ {
		news := News{
			Title:   newsList[i].Title,
			Time:    newsList[i].Time,
			Content: newsList[i].Content,
			Source:  newsList[i].Source,
			Link:    newsList[i].Link,
		}

		images := []Image{}
		for j := 0; j < len(newsList[i].Images); j++ {
			tmpImage := Image{
				Link:        newsList[i].Images[j].Link,
				Description: newsList[i].Images[j].Description,
			}
			images = append(images, tmpImage)
		}

		news.Images = images

		view = append(view, news)
	}
	return view
}

package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
)

func main(){

	c := colly.NewCollector()

	detailCollector := c.Clone()

	c.OnHTML("li[id] a",func(e *colly.HTMLElement){

		link := e.Attr("href")
		if !(strings.HasSuffix(link,"company/")||strings.HasSuffix(link, "code-for-a-living/")||
			strings.HasSuffix(link, "insights/")||strings.HasSuffix(link, "engineering/")||
			strings.HasSuffix(link, "podcasts/")||strings.HasSuffix(link, "nav-developer-hiring-blog")){
			return
		}

		c.Visit(link)

	})

	c.OnHTML("article div.m-post-card__content-column",func(e *colly.HTMLElement){


		link,_ := e.DOM.Find("h2.m-post-card__title").Find("a").Attr("href")


		detailCollector.Visit(link)


	})

	detailCollector.OnHTML("div.column.is-two-thirds",func(e *colly.HTMLElement){
		title := strings.TrimSpace(e.DOM.Find("h1.section-title").Text())
		content := e.DOM.Find("div.m-post-content").Text()
		fmt.Printf("\n <<%s>> \n content: %s \n ----------------------------------",title,content)
	})

	c.OnRequest(func(r *colly.Request){
		fmt.Println("visiting ",r.URL.String())
	})

	c.Visit("https://stackoverflow.blog")


}
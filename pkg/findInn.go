package pkg

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func FindInn(companyName string) (Inn []string) {
	c := colly.NewCollector()

	lookingFor := "ИНН"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		text := e.Text
		if strings.Contains(e.Text, lookingFor) {
			indexOfLookingFor := strings.Index(text, lookingFor)
			text = text[indexOfLookingFor:]
			text = text[strings.Index(text, "Н ")+3:]
			re := regexp.MustCompile("[0-9]+")

			Inn = append(Inn, re.FindString(text))

		}
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println(companyName+"error: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
	})

	c.Visit("https://www.google.com/search?q=" + companyName + "+ИНН")
	unique(&Inn)
	fmt.Println("Finded inn of "+companyName+" = ", Inn)

	return
}

func unique(strokeSlice *[]string) {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range *strokeSlice {

		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	*strokeSlice = list
}

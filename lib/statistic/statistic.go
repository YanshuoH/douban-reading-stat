package statistic

import (
  "encoding/json"
  "time"
  "strconv"
  "regexp"

  "github.com/YanshuoH/douban-reading-stat/models"
)

const (
  DATE_LAYOUT = "2006-01-02 15:04:05"
  PRICE_REGEXP = "[0-9.]+"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func initializeYearEntity(year string) *models.StatEntity {
  entity := new(models.StatEntity)

  entity.Count = 0
  entity.Price = 0

  ratingMap := make(map[string]int)
  for i := 1; i <= 5; i++ {
    ratingMap[strconv.Itoa(i)] = 0
  }
  entity.Rating = ratingMap

  monthMap := make(map[string]int)
  for i := 1; i <= 12; i++ {
    monthMap[strconv.Itoa(i)] = 0
  }
  entity.Month = monthMap

  entity.Author = make(map[string]int)
  entity.Translator = make(map[string]int)
  entity.Publisher = make(map[string]int)
  entity.Posters = make([]map[string]string, 0)
  entity.Tags = make(map[string]int)

  return entity
}

func Stat(user *models.User) map[string]*models.StatEntity {
  result := make(map[string]*models.StatEntity)
  priceRe := regexp.MustCompile(PRICE_REGEXP)

  for _, post := range user.Books {
    p := post.(map[string]interface{})

    // Only calculate read state
    if p["status"].(string) != "read" {
      continue
    }

    // Year parsing
    updated, err := time.Parse(DATE_LAYOUT, p["updated"].(string))
    checkErr(err)
    year := strconv.Itoa(updated.Year())

    // Initial year key
    if _, ok := result[year]; !ok {
      result[year] = initializeYearEntity(year)
      y, _ := strconv.Atoi(year)
      result[year].Year = y
    }
    result[year].Count ++

    // Sum rating
    if _, ok := p["rating"]; ok {
      userRatingEntity := p["rating"].(map[string]interface{})
      userRating := userRatingEntity["value"].(string)

      result[year].Rating[userRating] ++
    }

    // Sum month
    markedMonth := int(updated.Month())
    result[year].Month[strconv.Itoa(markedMonth)] ++

    // Enter book block in json
    book := models.BookEntity{}
    res, err := json.Marshal(p["book"])
    checkErr(err)
    json.Unmarshal(res, &book)

    // Sum price
    if len(book.Price) > 0 {
      priceString := priceRe.FindString(book.Price)
      if priceString != "" {
        price, err := strconv.ParseFloat(priceString, 64)
        checkErr(err)
        result[year].Price += price
      }
    }

    // Sum author
    if len(book.Author) > 0 {
      for _, value := range book.Author {
        // Check if author in result list
        if _, ok := result[year].Author[value]; !ok {
          result[year].Author[value] = 0
        }
        result[year].Author[value] ++
      }

    }

    // Sum translator
    if len(book.Translator) > 0 {
      for _, value := range book.Translator {
        // Check if translator in result array
        if _, ok := result[year].Translator[value]; !ok {
          result[year].Translator[value] = 0
        }
        result[year].Translator[value] ++
      }
    }

    // Sum publisher
    if len(book.Publiser) > 0 {
      publisherString := book.Publiser
      if _, ok := result[year].Publisher[publisherString]; !ok {
        result[year].Publisher[publisherString] = 0
      }
      result[year].Publisher[publisherString] ++
    }

    // Sum Tags
    if len(book.Tags) > 0 {
      for _, value := range book.Tags {
        // Check if tags in result array
        tagname := value["name"].(string)
        if _, ok := result[year].Tags[tagname]; !ok {
          result[year].Tags[tagname] = 0
        }
        result[year].Tags[tagname] ++
      }
    }

    // Posters
    title := book.Title
    link := book.Images["large"]
    poster := map[string]string{
      "title": title,
      "link": link,
    }
    result[year].Posters = append(result[year].Posters, poster)
  } // End for

  return result
}

package api

import (
	"math/rand"
	"shortener/internal/model"
	"sort"
	"strings"
)

const (
	alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlLength = 8
)

func generateShortURL() string {
	alphabetLength := len(alphabet)
	shortURL := make([]byte, urlLength)
	for i := range shortURL {
		shortURL[i] = alphabet[rand.Intn(alphabetLength)]
	}
	return string(shortURL)
}

func sortUrls(orderType string) []*model.URLShortener {
	var urlSlices []*model.URLShortener

	if orderType == "" {
		orderType = "asc"
	}

	for _, v := range urls {
		urlSlices = append(urlSlices, &model.URLShortener{
			Destination:  v.Destination,
			ShortURL:     v.ShortURL,
			ClickCounter: v.ClickCounter,
			ExpireAt:     v.ExpireAt,
		})
	}
	if strings.ToLower(orderType) == "asc" {
		sort.Slice(urlSlices, func(i, j int) bool {
			return urlSlices[i].ClickCounter < urlSlices[j].ClickCounter
		})
	} else {
		sort.Slice(urlSlices, func(i, j int) bool {
			return urlSlices[i].ClickCounter > urlSlices[j].ClickCounter
		})
	}

	return urlSlices
}

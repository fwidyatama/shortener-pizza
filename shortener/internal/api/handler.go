package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"shortener/internal/model"
	"shortener/internal/response"
	"shortener/internal/transport"
	"strings"
	"sync"
	"time"
)

var (
	urlMutex sync.Mutex
	urls     = make(map[string]*model.URLShortener)
	validate = validator.New()
)

type apiHandler struct{}

type ApiHandler interface {
	ShortURL(w http.ResponseWriter, r *http.Request)
	RedirectURL(w http.ResponseWriter, r *http.Request)
	GetListURL(w http.ResponseWriter, r *http.Request)
}

func NewHandler() ApiHandler {
	return &apiHandler{}
}

func (a apiHandler) ShortURL(w http.ResponseWriter, r *http.Request) {
	hLog := log.WithField("function", "ShortURL")
	req := transport.ShortURLReq{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hLog.Println("error", err)

		errMsg := "Error when processing request"

		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	err = validate.Struct(req)
	if err != nil {
		hLog.Println("error", err)

		errMsg := "Error when processing request"

		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	clickCount := 0

	expireAt, _ := time.Parse(time.DateTime, req.ExpireAt)
	url := &model.URLShortener{
		Destination:  req.Destination,
		ShortURL:     fmt.Sprintf("%s/%s", r.Host, shortURL),
		ClickCounter: clickCount,
		ExpireAt:     expireAt,
	}

	urlMutex.Lock()
	urls[shortURL] = url
	urlMutex.Unlock()

	response.SuccessResponse(w, url)

}

func (a apiHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	hLog := log.WithField("function", "RedirectURL")
	vars := mux.Vars(r)
	shortURL := vars["url"]
	if shortURL == "" {
		errMsg := "Short URL cannot be empty"
		hLog.Println("error", errMsg)

		response.ErrorResponse(w, errMsg, http.StatusBadRequest)

		return
	}

	urlMutex.Lock()
	defer urlMutex.Unlock()

	url, isFound := urls[shortURL]
	if !isFound {
		errMsg := "URL not found"
		hLog.Println("error", errMsg)

		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	if time.Now().After(url.ExpireAt) {
		errMsg := "URL is expired"
		hLog.Println("error", errMsg)

		response.ErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	url.ClickCounter++

	redirectURL := ""

	// if contain https, remove it
	if strings.Contains(url.Destination, "https://") {
		url.Destination = strings.Replace(url.Destination, "https://", "", 1)
	}

	redirectURL = fmt.Sprintf("%s%s", "https://", url.Destination)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (a apiHandler) GetListURL(w http.ResponseWriter, r *http.Request) {
	urlMutex.Lock()
	defer urlMutex.Unlock()

	orderParams := r.URL.Query().Get("order")
	sortedURL := make([]*model.URLShortener, 0)
	for _, url := range urls {
		sortedURL = append(sortedURL, url)
	}

	result := sortUrls(orderParams)
	response.SuccessResponse(w, result)
}

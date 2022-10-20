package controllers

import (
	"chi-project/interfaces"
	"encoding/json"
	"net/http"
)

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var articles []interfaces.Article

func GetArticle(w http.ResponseWriter, r *http.Request) {
	if len(articles) < 1 {
		articles = []interfaces.Article{
			{
				Id:      "1",
				Title:   "Article 1",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			},
		}
	}

	response := response{
		StatusCode: 200,
		Message:    "Success Get Article",
		Data:       articles,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

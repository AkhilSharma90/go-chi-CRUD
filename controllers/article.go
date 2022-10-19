package controllers

import (
	"chi-project/interfaces"
)

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var articles []interfaces.Article

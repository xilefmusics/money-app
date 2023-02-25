package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"xilefmusics.de/money-app/helper"
	"xilefmusics.de/money-app/transaction"
)

func Lint(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	lint := transaction.LintTransactions(transactions)

	gc.IndentedJSON(http.StatusOK, lint)
}

func Reindex(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	globalData.Reindex(user)

	gc.IndentedJSON(http.StatusOK, "")
}

func Undo(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	globalData.Undo(user)

	// TODO return applied event
	gc.IndentedJSON(http.StatusOK, "")
}

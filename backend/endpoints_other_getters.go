package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"xilefmusics.de/money-app/helper"
	"xilefmusics.de/money-app/history"
	"xilefmusics.de/money-app/transaction"
)

func GetPods(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	pods := transaction.GetPods(transactions)

	gc.IndentedJSON(http.StatusOK, pods)
}

func GetDebts(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	debts := transaction.GetDebts(transactions)

	gc.IndentedJSON(http.StatusOK, debts)
}

func GetBudgets(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	budgets := transaction.GetBudgets(transactions)

	gc.IndentedJSON(http.StatusOK, budgets)
}

func GetInbudgets(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR in GetTransactions: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	inbudget := transaction.GetInbudgets(transactions)

	gc.IndentedJSON(http.StatusOK, inbudget)
}

func GetTags(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR in GetTransactions: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	transactions := globalData.GetTransactions(user)
	tags := transaction.GetTags(transactions)

	gc.IndentedJSON(http.StatusOK, tags)
}

func GetHistory(gc *gin.Context) {
	user, err := helper.GC2User(gc)
	if err != nil {
		log.Printf("ERROR in GetTransactions: %s\n", err.Error())
		gc.String(http.StatusInternalServerError, "501 Internal Server Error")
		return
	}

	year := helper.ParseNumber(gc.DefaultQuery("year", "0"), 0, true)
	month := helper.ParseNumber(gc.DefaultQuery("month", "0"), 0, true)
	day := helper.ParseNumber(gc.DefaultQuery("day", "0"), 0, true)
	len := helper.ParseNumber(gc.DefaultQuery("len", "1"), 1, true)

	if day == 0 && month == 0 && year == 0 {
		year = 1
	}

	transactions := globalData.GetTransactions(user)

	switch gc.Param("kind") {
	case "debt":
		gc.IndentedJSON(http.StatusOK, history.History(int(year), int(month), int(day), int(len), history.Debt, transactions))
		break
	case "budget":
		gc.IndentedJSON(http.StatusOK, history.History(int(year), int(month), int(day), int(len), history.Budget, transactions))
		break
	case "inbudget":
		gc.IndentedJSON(http.StatusOK, history.History(int(year), int(month), int(day), int(len), history.Inbudget, transactions))
		break
	case "pod":
		gc.IndentedJSON(http.StatusOK, history.History(int(year), int(month), int(day), int(len), history.Pod, transactions))
		break
	case "wealth":
		gc.IndentedJSON(http.StatusOK, history.History(int(year), int(month), int(day), int(len), history.Wealth, transactions))
		break
	default:
		gc.IndentedJSON(http.StatusNotFound, "404 page not found")
	}
}
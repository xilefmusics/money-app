package data

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"xilefmusics.de/money-app/event"
	"xilefmusics.de/money-app/transaction"
)

type Data map[string]*UserData

func New(dataPath string) (Data, error) {
	transactionsPath := filepath.Join(dataPath, "transactions")
	if _, err := os.Stat(transactionsPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(transactionsPath, os.ModePerm); err != nil {
				return Data{}, err
			}
		}
	}

	eventsPath := filepath.Join(dataPath, "events")
	if _, err := os.Stat(eventsPath); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(eventsPath, os.ModePerm); err != nil {
				return Data{}, err
			}
		}
	}

	transactions := make(map[string][]transaction.Transaction)
	transactionsMutex := make(map[string]*sync.Mutex)

	err := filepath.Walk(transactionsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(info.Name()) == ".json" {
			user := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			transactions[user], _ = transaction.Load(path)
			transactionsMutex[user] = &sync.Mutex{}
			log.Printf("INFO: Load transactions of user %s from path %s", user, path)
		}

		return nil
	})

	events := make(map[string][]event.Event)
	eventsMutex := make(map[string]*sync.Mutex)

	err = filepath.Walk(eventsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(info.Name()) == ".json" {
			user := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			events[user], _ = event.Load(path)
			eventsMutex[user] = &sync.Mutex{}
			log.Printf("INFO: Load events of user %s from path %s", user, path)
		}

		return nil
	})

	return Data{dataPath, transactionsPath, eventsPath, transactions, transactionsMutex, events, eventsMutex}, err
}

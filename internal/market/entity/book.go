package entity

import "sync"

type Book struct {
	Order []*Order
	Transaction []*Transaction
	OrdersChan chan *Order
	OrdersChanOut chan *Order
	wg *sync.WaitGroup
}

func NewBook()
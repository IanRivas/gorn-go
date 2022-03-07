package main

import (
	"github.com/IanRivas/gorn-go/model"
	"github.com/IanRivas/gorn-go/storage"
)

func main() {
	storage.NewPostgresDB()

	invoice := model.InvoiceHeader{
		Client: "Riki Fort",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)

}

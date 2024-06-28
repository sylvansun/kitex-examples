package main

import (
	"log"

	item "example_shop/kitex_gen/example/shop/item/itemservice"
)

func main() {
	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient("0.0.0.0:8890")
	if err != nil {
		log.Fatal(err)
	}
	itemServiceImpl.stockCli = stockCli

	svr := item.NewServer(itemServiceImpl)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

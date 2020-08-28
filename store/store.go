package store

import (
	"aavaz/store/adaptee/inmemory"
	"aavaz/store/adapter"
	"fmt"
	"log"
)

// Store global store connection interface
var Store adapter.Store

// Init loads the sample data and prepares the store layer
func Init() {
	// store inmemory adapter ...
	Store = inmemory.NewAdapter()
	if Store == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}
	fmt.Println("Inited Store...👍")
}

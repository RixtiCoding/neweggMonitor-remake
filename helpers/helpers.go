package helpers

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"sync"
)

type Product struct {
	MainItem MainItem `json:"MainItem"`
}

type MainItem struct {
	Description Description `json:"Description"`
	Image       Image       `json:"Image"`
	InStock     bool        `json:"InStock"`
	Price       float64     `json:"FinalPrice"`
}

type Normal struct {
	ImageName string `json:"ImageName"`
}

type Description struct {
	Title       string `json:"Title"`
	UrlKeywords string `json:"UrlKeywords"`
}

type Image struct {
	Normal Normal `json:"Normal"`
}

var (
	mu sync.Mutex
)

func getSkus() []string {
	mu.Lock()
	file, err := os.Open("skus.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	var skus []string
	for scanner.Scan() {
		skus = append(skus, scanner.Text())
	}
	_ = file.Close()

	if len(skus) == 0 {
		panic("Please add a sku!")
	}
	mu.Unlock()
	return skus
}

func getProxy() string {
	mu.Lock()
	file, err := os.Open("proxies.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	_ = file.Close()

	if len(txtlines) == 0 {
		panic("Please add proxies to proxies.txt")
	}

	index := rand.Intn(len(txtlines))
	mu.Unlock()
	return txtlines[index]
}

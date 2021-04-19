package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var headers = map[string]string{
	"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
	"Content-Type":  "application/x-www-form-urlencoded; charset=UTF-8",
	"cache-control": "no-store,no-cache,must-revalidate,proxy-revalidate,max-age=0",
	"pragma":        "no-cache",
}

func Monitor() {
	discordWebhook := "https://discord.com/api/webhooks/831545615530852383/ouOFr_LvWO78JZYJtsXCkggn3wT-8MRN5qDmlkxAR5hd0mgy7E-JTKiIDN6ge9rZGbPc"
	for {
		/*proxyUrl, err := url.Parse(getProxy())
		if err != nil {
			fmt.Println("Proxy Parsing failed!")
			continue
		}
		*/
		client := &http.Client{}
		for _, v := range getSkus() {
			req, err := client.Get(fmt.Sprintf("https://www.newegg.com/product/api/ProductRealtime?ItemNumber=%s", v))
			if err != nil {
				fmt.Println(err)
				continue
			}
			for k, v := range headers {
				req.Header.Set(k, v)
			}
			resp, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_ = req.Body.Close()
			var product Product
			err = json.Unmarshal(resp, &product)
			if err != nil {
				fmt.Println(err)
				continue
			}
			IsInStock := product.MainItem.InStock
			sku := v
			if IsInStock {
				sendWebhook(product, discordWebhook, sku)

			}
		}
	}

}

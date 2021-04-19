package helpers

import (
	"fmt"

	"github.com/aiomonitors/godiscord"
)

func sendWebhook(product Product, discordWebhook string, sku string) {
	imageUrl := fmt.Sprintf("https://c1.neweggimages.com/ProductImageCompressAll1280/%s", product.MainItem.Image.Normal.ImageName)
	productUrl := fmt.Sprintf("https://www.newegg.com/%s/p/%s", product.MainItem.Description.UrlKeywords, sku)
	emb := godiscord.NewEmbed(product.MainItem.Description.Title, "", productUrl)

	emb.SetThumbnail(imageUrl)
	emb.SetColor("#e7a348")
	emb.SetFooter("@RixtiRobotics", "")
	emb.AddField("Price", fmt.Sprintf("$%.2f", product.MainItem.Price), true)
	emb.AddField("SKU", sku, true)

	emb.Username = "NewEgg Monitor"

	emb.SendToWebhook(discordWebhook)
}

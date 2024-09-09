package crypto

import (
	"fmt"
	"time"
)

func CryptoTracker(currency string, interval int) {

	fmt.Printf("Videoni kuzatish: %s\n", currency)
	fmt.Printf("Ma'lumotlar %d sekund interval bilan yangilanadi...\n", interval)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C: // Har safar interval o'tganda
			price, err := GetCryptoPrice(currency)
			if err != nil {
				fmt.Printf("Xatolik: %v\n", err)
			} else {
				fmt.Printf("%s uchun joriy narx: %.2f USD\n", currency, price.Price)
			}
		}
	}
}

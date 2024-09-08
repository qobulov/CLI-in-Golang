package Crypto_tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCryptoPrice API'dan kriptovalyuta narxini olish
func GetCryptoPrice(currency string) (float64, error) {
	apiUrl := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", currency)

	// So'rovni yuborish uchun HTTP clientini tayyorlash
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return 0, fmt.Errorf("so'rov tayyorlashda xatolik: %v", err)
	}

	// User-Agent qo'shish
	req.Header.Set("User-Agent", "crypto-tracker-app")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("API so'rovida xatolik: %v", err)
	}
	defer resp.Body.Close()

	// HTTP status kodini tekshirish
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API dan noto'g'ri javob: %d - %s", resp.StatusCode, resp.Status)
	}

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("JSON decodingda xatolik: %v", err)
	}

	// Ma'lumot mavjudligini tekshirish
	if price, ok := result[currency]["usd"]; ok {
		return price, nil
	}

	return 0, fmt.Errorf("narx topilmadi")
}

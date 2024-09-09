package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Crypto struct, kriptovalyuta ma'lumotlarini saqlash uchun
type Crypto struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

// GetCryptoPrice funksiyasi, kriptovalyutaning joriy narxini API orqali oladi
func GetCryptoPrice(cryptoName string) (Crypto, error) {
	// CoinGecko talab qilgan formatda kriptovalyuta nomini kichik harflarga o'zgartiramiz
	cryptoName = strings.ToLower(cryptoName)
	apiURL := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", cryptoName)

	// API'ga so'rov yuboramiz
	res, err := http.Get(apiURL)
	if err != nil {
		return Crypto{}, err
	}
	defer res.Body.Close()

	// Agar API noto'g'ri javob qaytarsa
	if res.StatusCode != 200 {
		return Crypto{}, fmt.Errorf("crypto API'dan noto'g'ri javob: %d", res.StatusCode)
	}

	// JSON ma'lumotlarni o'qiymiz
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Crypto{}, err
	}

	// JSONni deserializatsiya qilish uchun ajratamiz
	var cryptoData map[string]map[string]float64
	err = json.Unmarshal(body, &cryptoData)
	if err != nil {
		return Crypto{}, err
	}

	// Narxni olish
	price, ok := cryptoData[cryptoName]["usd"]
	if !ok {
		return Crypto{}, fmt.Errorf("kripto narxi topilmadi")
	}

	// Natijani qaytarish
	return Crypto{
		Symbol: cryptoName,
		Price:  price,
	}, nil
}

package main

import (
    "fmt"
    "os"
    "time"
    "crypto-tracker/tracker"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Foydalanish: crypto-tracker <kriptovalyuta nomi> [interval sekunda]")
        return
    }

    currency := os.Args[1]
    interval := 10 // default interval 10 sekund
    if len(os.Args) > 2 {
        fmt.Sscanf(os.Args[2], "%d", &interval)
    }

    fmt.Printf("Videoni kuzatish: %s\n", currency)
    fmt.Printf("Ma'lumotlar %d sekund interval bilan yangilanadi...\n", interval)

    ticker := time.NewTicker(time.Duration(interval) * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            price, err := tracker.GetCryptoPrice(currency)
            if err != nil {
                fmt.Printf("Xatolik: %v\n", err)
            } else {
                fmt.Printf("%s uchun joriy narx: %.2f USD\n", currency, price)
            }
        }
    }
}

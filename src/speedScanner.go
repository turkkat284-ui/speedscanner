package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var ip string
	fmt.Print("Enter an IP address: ")
	fmt.Scan(&ip)

	var wg sync.WaitGroup

	fmt.Printf("\nScanning %s...\n", ip)

	for i := 1; i <= 1024; i++ {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", ip, port)

			// DialTimeout kullanarak yanıt vermeyen portlarda takılmayı önleriz
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err != nil {
				return
			}
			conn.Close()

			fmt.Printf("[+] Port %d is open\n", port)
		}(i) // i değişkenini parametre olarak geçmek önemlidir
	}

	wg.Wait() // Tüm tarama bittikten sonra bekle
	fmt.Println("Scan completed.")
}

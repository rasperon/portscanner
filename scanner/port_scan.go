package scanner

import (
	"fmt"
	"net"
	"sync"
)

func ScanPort(host string, port int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	conn.Close()
	results <- fmt.Sprintf("Port %d is open", port)
}

func PortScanMenu() (string, int) {
	var host string
	var portRange int

	fmt.Print("Enter target host (e.g., scanme.nmap.org): ")
	fmt.Scan(&host)

	fmt.Print("Enter port range to scan (e.g., 1024): ")
	fmt.Scan(&portRange)

	return host, portRange
}

package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/rasperon/portscanner/api"
	"github.com/rasperon/portscanner/scanner"
	"github.com/rasperon/portscanner/banner"
	"github.com/rasperon/portscanner/utils"
)

func main() {
	utils.ClearScreen()
	banner.PrintBanner()

	fmt.Println("Select an option:")
	fmt.Println("1. Scan ports")
	fmt.Println("2. Exit")

	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		host, portRange := scanner.PortScanMenu()
		results := make(chan string, portRange)
		var wg sync.WaitGroup

		for port := 1; port <= portRange; port++ {
			wg.Add(1)
			go scanner.ScanPort(host, port, results, &wg)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		opens := "Open Ports:\n"
		for result := range results {
			fmt.Println(result)
			opens += result + "\n"
		}

		fmt.Print("Enter language for analysis (e.g., en, tr, es): ")
		var language string
		fmt.Scan(&language)

		api.AnalyzePorts(opens, language)

	case 2:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice! Exiting...")
		os.Exit(1)
	}
}

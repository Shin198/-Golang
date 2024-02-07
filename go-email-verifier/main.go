package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain :: HasMX :: HasSPF :: SPFRecord :: HasDMARC :: DMARCRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: Can't read from input : %v\n", err)
	}
}

func checkDomain(domain string) {
	var HasMX, HasSPF, HasDMARC bool
	var SPFRecord, DMARCRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		HasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error :%v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			HasSPF = true
			SPFRecord = record
			break
		}
	}

	DMARCRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range DMARCRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			HasDMARC = true
			DMARCRecord = record
			break
		}
	}

	fmt.Printf("%v:: %v :: %v :: %v :: %v :: %v", domain, HasMX, HasSPF, SPFRecord, HasDMARC, DMARCRecord)
}

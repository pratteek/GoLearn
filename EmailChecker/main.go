package main

import (
	//buffer package to parse input
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string){
	var hasMX,hasSPF,hasDMARC bool
	var spfRecord, dmarcRecord string

	//using net package to find the records
	mxRecords, err:= net.LookupMX(domain)

	if err!=nil{
		log.Printf("Error: %v",err)
	}
	if(len(mxRecords)>0){
		hasMX= true
	}
	
	txtRecords, err := net.LookupTXT(domain)
	
	if err!=nil{
		log.Printf("Error: %v",err)
	}
	for _, record := range txtRecords{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF=true
			spfRecord=record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc."+domain)
	
	if err!=nil{
		log.Printf("Error: %v",err)
	}
	for _, record := range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC=true
			dmarcRecord=record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v",domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord:\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err!=nil{
		log.Fatal("Error: Could not read from input: ",err)
	}
}


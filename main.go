package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Set log level to debug
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	respBytes, err := hex.DecodeString(text[:len(text)-1])
	failOnError(err)

	// newDecodedTx := DecodeProcessedTransaction(respBytes)
	newDecodedTx := &ParsedProcessedTransaction{}
	err = newDecodedTx.DecodeProcessedTransaction(respBytes)
	failOnError(err)

	//MarshalIndent
	newDecodedTxJSON, err := json.MarshalIndent(newDecodedTx, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("DecodedProcessedTransaction: %s\n", string(newDecodedTxJSON))
	failOnError(err)

}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

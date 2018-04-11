package main

import (
	"github.com/jonSteck/adx"
	"log"
	"fmt"
	"strings"
	"os"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"github.com/gogo/protobuf/proto"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter bid request file: ")
	bidRequestFile, _ := reader.ReadString('\n')
	bidRequestFile = strings.Replace(bidRequestFile, "\n", "", -1)

	raw, err := ioutil.ReadFile(bidRequestFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var adxStruct adx.BidRequest
	json.Unmarshal(raw, &adxStruct)
	out, err := proto.Marshal(&adxStruct)
	if err != nil {
		log.Fatalln("Failed to encode bid request:", err)
	}
	if err := ioutil.WriteFile("proto.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write protobuf bid request to file:", err)
	}
}

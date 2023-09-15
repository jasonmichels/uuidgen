package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	uuid, err := NewUUIDv4()
	if err != nil {
		fmt.Printf("Error generating UUID: %v\n", err)
		return
	}
	fmt.Printf("UUIDV4: %s \n", uuid)

	unix := UnixTime()
	fmt.Printf("Unixtimestamp: %v \n", unix)
}

func UnixTime() int64 {
	return time.Now().Unix()
}

func NewUUIDv4() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Set version 4 bits
	bytes[6] = (bytes[6] & 0x0F) | 0x40
	// Set variant bits as specified in RFC4122
	bytes[8] = (bytes[8] & 0x3F) | 0x80

	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		bytes[0], bytes[1], bytes[2], bytes[3],
		bytes[4], bytes[5],
		bytes[6], bytes[7],
		bytes[8], bytes[9],
		bytes[10], bytes[11], bytes[12], bytes[13], bytes[14], bytes[15]), nil
}

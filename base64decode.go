package main
import b64 "encoding/base64"
import "fmt"
import "encoding/hex"

func main() {
	data := "Af//q6u8vJKSAAA="
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data, hex.Dump(sDec))
}


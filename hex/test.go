package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const PREVIOUS_SEED_SHA = "6f0b86a1066a151e20a4576f8faa76e95f007b2e82df55ae4eb973004150911e"
const PREVIOUS_SEED = "364fb971ae36ad3b3766b1acda67ef394a2cbdd492f6d76d66737d04c9b0f1d9"

func main() {

	//originalNumber := "24565754629508838410361435551397037334894993521323729218026582611797304472025"
	number := 0
	for {
		number += 1
		paddedNumber := padLeft(strconv.Itoa(number), "0", 64)
		hexNumber, _ := decimalToHex(paddedNumber)
		result := strings.ToLower(hexNumber)
		value, _ := stringToSHA256(result)
		fmt.Println(number, "value, ", value)

		if number == 5000000 {
			break
		}
	}
}

func stringToSHA256(input string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)
	hashInHex := hex.EncodeToString(hashInBytes)
	return hashInHex, nil
}

func padLeft(str, pad string, length int) string {
	count := length - len(str)
	if count <= 0 {
		return str
	}
	return strings.Repeat(pad, count) + str
}

func decimalToHex(decimalStr string) (string, error) {
	decimal := new(big.Int)
	decimal.SetString(decimalStr, 10)
	return fmt.Sprintf("%X", decimal), nil
}

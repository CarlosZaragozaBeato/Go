package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const NEXT_SEED = "56b7fb69db511157c1ba691bcacda94c1111f1f3b7c446634922d9b6e3098aa2"
const (
	PREVIOUS_SEED_SHA = "6f0b86a1066a151e20a4576f8faa76e95f007b2e82df55ae4eb973004150911e"
	PREVIOUS_SEED     = "364fb971ae36ad3b3766b1acda67ef394a2cbdd492f6d76d66737d04c9b0f1d9"
)

func main() {
	// number := 24565754629508838410361435551397037334894993521323729218026582611797304472025
	number := 0
	number_goal := number + 10000000

	proceso(number, number_goal)
}

// FUNCION PRINCIPAL

func proceso(number, number_goal int) {
	for {
		paddedNumber := padLeft(strconv.Itoa(number))
		hexNumber, _ := decimalToHex(paddedNumber)
		result := strings.ToLower(hexNumber)
		value, _ := stringToSHA256(result)

		if value == NEXT_SEED {
			fmt.Println("ENCONTRADO")
			fmt.Println("NUMERO HEXADECIMAL: ", hexNumber)
			fmt.Println("NUMERO DECIMAL: ", number)
			fmt.Println("NUMERO sha256: ", value)
			break
		}

		if number == number_goal {
			break
		}
		number += 1
	}
}

// FUNCIONES COMPLEMENTARIAS
func stringToSHA256(input string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)
	hashInHex := hex.EncodeToString(hashInBytes)
	return hashInHex, nil
}

func padLeft(str string) string {
	count := 64 - len(str)
	if count <= 0 {
		return str
	}
	return strings.Repeat("0", count) + str
}

func decimalToHex(decimalStr string) (string, error) {
	decimal := new(big.Int)
	decimal.SetString(decimalStr, 10)
	return fmt.Sprintf("%X", decimal), nil
}

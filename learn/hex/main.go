package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	Inicio int `json:"inicio"`
}

const NEXT_SEED = "c246fb1b6dfb974b2f4185de1425d27901103ba3b75e5de955aabc067552b969"

func main() {

	for i := 0; i < 5; i++ {
		bytesJSON, _ := os.ReadFile("/home/carlos/Go/hex/inicio.json")
		var config Config
		json.Unmarshal(bytesJSON, &config)
		number := config.Inicio

		var wg sync.WaitGroup

		resultados := make(chan *Resultado, 4)

		wg.Add(1)
		go buscarNumeroConcurrente(number, number+100000, &wg, resultados)
		wg.Add(1)
		go buscarNumeroConcurrente(number+100000, number+200000, &wg, resultados)
		wg.Add(1)
		go buscarNumeroConcurrente(number+200000, number+300000, &wg, resultados)
		wg.Add(1)
		go buscarNumeroConcurrente(number+300000, number+400000, &wg, resultados)

		go func() {
			wg.Wait()
			close(resultados)
		}()

		for resultado := range resultados {
			if resultado.encontrado {
				fmt.Println("ENCONTRADO")
				fmt.Println("NUMERO HEXADECIMAL: ", resultado.hexNumber)
				fmt.Println("NUMERO DECIMAL: ", resultado.decimalNumber)
				fmt.Println("NUMERO sha256: ", resultado.sha256Value)
			}
		}

		config.Inicio += 40000000
		nuevoBytesJSON, _ := json.MarshalIndent(config, "", "  ")
		os.WriteFile("/home/carlos/Go/hex/inicio.json", nuevoBytesJSON, 0644)
		fmt.Println("Todas las goroutines han terminado. index -> ", i)
	}
}

type Resultado struct {
	decimalNumber int
	hexNumber     string
	sha256Value   string
	encontrado    bool
}

func buscarNumeroConcurrente(inicio, fin int, wg *sync.WaitGroup, resultados chan<- *Resultado) {
	defer wg.Done()

	for i := inicio; i <= fin; i++ {
		paddedNumber := padLeft(strconv.Itoa(i))
		hexNumber, _ := decimalToHex(paddedNumber)
		result := strings.ToLower(hexNumber)
		value, _ := stringToSHA256(result)
		fmt.Println("VALOR, %q", value)

		if value == NEXT_SEED {
			resultados <- &Resultado{
				decimalNumber: i,
				hexNumber:     hexNumber,
				sha256Value:   value,
				encontrado:    true,
			}
			return
		}

	}

	resultados <- &Resultado{encontrado: false}
}

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

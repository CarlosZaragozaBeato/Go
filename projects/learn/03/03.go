package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"time"
)

type Options struct {
	Length       int
	UseUpperCase bool
	UseNumbers   bool
	UseSymbols   bool
}

func generatePassword(options Options) string {
	var chars string

	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:'<>,.?/~"

	chars += lowercaseLetters
	if options.UseUpperCase {
		chars += uppercaseLetters
	}
	if options.UseNumbers {
		chars += numbers
	}
	if options.UseSymbols {
		chars += symbols
	}
	if len(chars) == 0 {
		fmt.Println("Error: Debes seleccionar al menos un tipo de caracteres.")
		return ""
	}

	secureSeed, err := rand.Int(rand.Reader, big.NewInt(int64(time.Now().UnixNano())))
	if err != nil {
		fmt.Println("Error  al generar una semilla segura: ", err)
		return ""
	}

	rand.Seed(secureSeed.int64())

	password := make([]byte, options.Length)
	for i := 0; i < options.Length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}

func main() {
	var options Options
	flag.IntVar(&options.Length, "length", 12, "Longitud de la contraseña")
	flag.BoolVar(&options.UseUpperCase, "uppercase", true, "Incluir letras mayúsculas")
	flag.BoolVar(&options.UseNumbers, "numbers", true, "Incluir números")
	flag.BoolVar(&options.UseSymbols, "symbols", true, "Incluir símbolos")
	flag.Parse()

	password := generatePassword(options)

	if password != "" {
		fmt.Println("Contraseña Generada: ", password)
	}
}

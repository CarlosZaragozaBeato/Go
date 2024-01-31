package main

/* COSAS
HACER PROGRAMA CON GOROUTINES?
HACER PROGRAMA CON TIRADAS EN MINIBATCHES COMO CON STOCHASTIC GRADIENT ?
FUNCIONES ANONIMAS O NO?

*/

/* ENUNCIADO
La idea es la siguiente
Un fichero que lee el numero en el que empieza, convertirlo a su hexadecimal de 64 bits ,
hallar su Sha256 y compararlo con otro sha256 que estará en otro fichero , si coinciden avisar
y devolver el hexadecimal , si no restarle uno y seguir
*/

/* PASO 1: COMPLETADO
   -> COMPROBAR EL PASO DE STRING A Sha256
*/
/* PASO 2:
-> GENERACION DE NUMERO A HEX DE 64
*/

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

const NEXT_SEED_SHA = "f076e89b3c65076aa9331ee5a10b07eb4c3a01dba7296223bef7bfd735c82ecc"

const (
	PREVIOUS_SEED_SHA = "6f0b86a1066a151e20a4576f8faa76e95f007b2e82df55ae4eb973004150911e"
	PREVIOUS_SEED     = "364fb971ae36ad3b3766b1acda67ef394a2cbdd492f6d76d66737d04c9b0f1d9"
)

func main() {
	// temp_seed := "364fb971ae36ad3b3766b1acda67ef394a2cbdd492f6d76d66737d04c9b0f1d9"
	// num_temp := 28200103311205173000000000000000000
	numero := "17625064569503233297192216803934208"

	num_temp := new(big.Int)
	num_temp.SetString(numero, 10)

	// Convertir el número a cadena
	numeroComoCadena := num_temp.String()
	fmt.Println(numeroComoCadena)

	// Añadir ceros a la izquierda hasta alcanzar una longitud de 64
	numero_64 := fmt.Sprintf("%064s", numeroComoCadena)
	fmt.Println(numero_64)

	// Procesar el hash SHA (reemplace esto con su implementación real)
	read_sha(numero_64)
}

func read_sha(seed string) string {
	h := sha256.New()
	h.Write([]byte(seed))
	temp_seed_sha := hex.EncodeToString(h.Sum(nil))

	if temp_seed_sha != PREVIOUS_SEED_SHA {
		fmt.Println("CORRECT Sha256: %d", PREVIOUS_SEED_SHA)
		fmt.Println("INCORRECT Sha256: %d", temp_seed_sha)
	} else {
		fmt.Println("CORRECT : ")
		fmt.Println(PREVIOUS_SEED_SHA)
		fmt.Println(temp_seed_sha)
	}
	return temp_seed_sha
}

/*
* package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Número en formato hexadecimal
	numeroHex := "1a"

	// Convertir a decimal
	numeroDecimal, err := strconv.ParseInt(numeroHex, 16, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Imprimir el resultado
	fmt.Println("Número en decimal:", numeroDecimal)
}
* */

package fatorial

// Fatorial retorna o fatorial recursivo de um numero
func Fatorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * Fatorial(value-1)
	}
}

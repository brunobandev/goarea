package goarea

import "math"

// Pi é uma proporção numerica definida pela relação entre
// o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ é responsavel por calcular área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

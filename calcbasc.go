//simples calculadora para raízes de equações do segundo grau.
package main
import ("fmt"; "math")
func calcularDelta (a float64, b float64, c float64) float64 {
  delta := b*b - 4*a*c
	return delta
}
func calcularRaizes ( a float64, b float64, delta float64) (float64, float64) {
	var x1, x2 float64
	x1 = (-b + math.Sqrt(delta))/(2 * a)
	x2 = (-b - math.Sqrt(delta))/(2 * a)

	return x1, x2
}

func main() {
//variáveis
	var a, b, c float64

//coletar input do usuário
	fmt.Println(`
	==================================================
	Digite o valor de A, B e C (Separados por Espaço):
	==================================================
	`)
	fmt.Scan(&a,&b,&c)
//calculo do delta
	delta := calcularDelta (a, b, c)
//verificação das propriedades e calculo das raizes
	switch {
	 case delta == 0:
		var x1, _ = calcularRaizes (a, b, delta)
		fmt.Println(`
		Sua equação possui uma única raiz real.
		Única Raiz:`, x1)
	 case delta < 0:
		fmt.Println("Sua equação não possui raízes reais")
	 default:
		x1, x2 := calcularRaizes(a, b, delta)
		fmt.Println(`
		Sua equação possui duas raízes reais
		Primeira raiz:`, x1,
		`| Segunda Raiz:`, x2)
	}
}

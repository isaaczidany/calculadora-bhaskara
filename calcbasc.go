//simples calculadora para raízes de equações do segundo grau.
package main
import ("fmt"; "math")
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
	delta := b*b - 4*a*c
//diferenciar propriedades do delta
	switch {
	 case delta == 0:
		x1 := (-b + math.Sqrt(delta)) / (2*a)
		fmt.Println(`
		Sua equação possui uma única raiz real.
		Única Raiz:`, x1)
	 case delta < 0:
		fmt.Println("Sua equação não possui raízes reais")
	 default:
		x1 := (-b + math.Sqrt(delta)) / (2*a)
		x2 := (-b - math.Sqrt(delta)) / (2*a)
		fmt.Println(`
		Sua equação possui duas raízes reais
		Primeira raiz:`, x1,
		`| Segunda Raiz:`, x2)
	}
}

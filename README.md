# 🧮 Calculadora de Bhaskara em Go

Projeto simples em Go para cálculo das raízes de uma equação do segundo grau utilizando a fórmula de Bhaskara.

O objetivo deste projeto é praticar conceitos fundamentais da linguagem Go, como:
- Funções
- Escopo de variáveis
- Retorno múltiplo
- Estrutura condicional (`switch`)
- Organização de código

---

## 📌 Fórmula utilizada

Para uma equação no formato:

ax² + bx + c = 0

O discriminante (Δ) é calculado como:

Δ = b² − 4ac

As raízes são calculadas por:

x = (-b ± √Δ) / (2a)

---

## ⚙️ Funcionalidades

- Leitura dos coeficientes `a`, `b` e `c`
- Cálculo do valor de Δ (delta)
- Verificação dos casos:
  - Δ < 0 → não possui raízes reais
  - Δ = 0 → uma raiz real
  - Δ > 0 → duas raízes reais
- Exibição das raízes no terminal

---

## 🛠️ Tecnologias utilizadas

- Go (Golang)
- Biblioteca padrão (`fmt`, `math`)

---

## ▶️ Como executar

1. Certifique-se de ter o Go instalado
2. Clone o repositório
3. No terminal, execute:

```bash
go run calcbasc.go

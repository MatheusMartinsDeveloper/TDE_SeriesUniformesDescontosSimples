package main

import (
	"fmt"
	"math"
)

type SerieUniforme struct {
	valorSemJuros    float64
	taxaDeJuros      float64
	numeroDeParcelas int
}

type DescontosSimples struct {
	descontoSimples float64
	valorNominal    float64
	taxaDeJuros     float64
	numeroPeriodo   int
}

func (sua *SerieUniforme) calcularSerieUniformeAntecipada() float64 {
	resultado := (sua.valorSemJuros / (1 + sua.taxaDeJuros)) / ((1 - math.Pow(1+sua.taxaDeJuros, -float64(sua.numeroDeParcelas))) / sua.taxaDeJuros)
	return resultado
}

func (sup *SerieUniforme) calcularSerieUniformePostecipada() float64 {
	resultado := sup.valorSemJuros / ((1 - math.Pow(1+sup.taxaDeJuros, -float64(sup.numeroDeParcelas))) / sup.taxaDeJuros)
	return resultado
}

func (cds *DescontosSimples) calcularDescontoSimples() float64 {
	resultado := (cds.valorNominal * cds.taxaDeJuros * float64(cds.numeroPeriodo)) / (1 + cds.taxaDeJuros*float64(cds.numeroPeriodo))
	return resultado
}

func (cvn *DescontosSimples) calcularValorNominal() float64 {
	resultado := (cvn.descontoSimples * (1 + cvn.taxaDeJuros*float64(cvn.numeroPeriodo))) / (cvn.taxaDeJuros * float64(cvn.numeroPeriodo))
	return resultado
}

func (ctj *DescontosSimples) calcularTaxaDeJuros() float64 {
	resultado := ctj.descontoSimples / (float64(ctj.numeroPeriodo) * (ctj.valorNominal - ctj.descontoSimples))
	return resultado
}

func (cnp *DescontosSimples) calcularNumeroPeriodo() int {
	resultado := cnp.descontoSimples / (cnp.taxaDeJuros * (cnp.valorNominal - cnp.descontoSimples))
	return int(resultado)
}

func main() {
	loop := true
	var escolha int

	for loop {
		fmt.Println("Bem-Vindo a Calculadora de Séries Uniformes e Descontos Simples!")
		fmt.Println("Digite 1 para entrar no menu Séries Uniformes!")
		fmt.Println("Digite 2 para entrar no menu Descontos Simples!")
		fmt.Println("Digite 3 para encerrar o programa!")
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			{
				var escolhaSeriesUniformes int
				var valorValorSemJuros float64
				var valorTaxaDeJuros float64
				var valorNumeroDeParcelas int

				fmt.Println("Você entrou no menu Séries Uniformes!")
				fmt.Println("Digite 1 para calcular o Valor das Parcelas Antecipadas!")
				fmt.Println("Digite 2 para calcular o Valor das Parcelas Postecipadas!")
				fmt.Println("Digite 3 para voltar ao menu principal!")
				fmt.Scanln(&escolhaSeriesUniformes)

				switch escolhaSeriesUniformes {
				case 1:
					{
						fmt.Println("Você entrou no menu Antecipadas!")
						fmt.Println("Digite o Valor sem Juros:")
						fmt.Scanln(&valorValorSemJuros)
						fmt.Println("Digite a Taxa de Juros (decimal):")
						fmt.Scanln(&valorTaxaDeJuros)
						fmt.Println("Digite o Número de Parcelas:")
						fmt.Scanln(&valorNumeroDeParcelas)

						instanciaSeriaUniformeAntecipada := SerieUniforme{
							valorSemJuros:    valorValorSemJuros,
							taxaDeJuros:      valorTaxaDeJuros,
							numeroDeParcelas: valorNumeroDeParcelas,
						}

						fmt.Printf("O valor das parcelas é: R$ %.2f\n", instanciaSeriaUniformeAntecipada.calcularSerieUniformeAntecipada())
					}
				case 2:
					{
						fmt.Println("Você entrou no menu Postecipadas!")
						fmt.Println("Digite o Valor sem Juros:")
						fmt.Scanln(&valorValorSemJuros)
						fmt.Println("Digite a Taxa de Juros (decimal):")
						fmt.Scanln(&valorTaxaDeJuros)
						fmt.Println("Digite o Número de Parcelas:")
						fmt.Scanln(&valorNumeroDeParcelas)

						instanciaSeriaUniformePostecipada := SerieUniforme{
							valorSemJuros:    valorValorSemJuros,
							taxaDeJuros:      valorTaxaDeJuros,
							numeroDeParcelas: valorNumeroDeParcelas,
						}

						fmt.Printf("O valor das parcelas é: R$ %.2f\n", instanciaSeriaUniformePostecipada.calcularSerieUniformePostecipada())
					}
				case 3:
					{
						continue
					}
				default:
					{
						fmt.Println("Essa opção não é válida, tente outra!")
					}
				}
			}
		case 2:
			{
				var escolhaDescontosSimples int
				var valorDescontoSimples float64
				var valorValorNominal float64
				var valorTaxaDeJuros float64
				var valorNumeroPeriodo int

				fmt.Println("Você entrou no menu Descontos Simples!")
				fmt.Println("Digite 1 para calcular o Desconto Racional Simples!")
				fmt.Println("Digite 2 para calcular o Valor Nominal!")
				fmt.Println("Digite 3 para calcular a Taxa de Juros!")
				fmt.Println("Digite 4 para calcular o Período de Antecipação!")
				fmt.Println("Digite 5 para voltar ao menu principal!")
				fmt.Scanln(&escolhaDescontosSimples)

				switch escolhaDescontosSimples {
				case 1:
					{
						fmt.Println("Você entrou no menu Desconto Racional Simples!")
						fmt.Println("Digite o Valor Nominal:")
						fmt.Scanln(&valorValorNominal)
						fmt.Println("Digite a Taxa de Juros (decimal):")
						fmt.Scanln(&valorTaxaDeJuros)
						fmt.Println("Digite o Período de Antecipação (meses):")
						fmt.Scanln(&valorNumeroPeriodo)

						instanciaDescontoSimples := DescontosSimples{
							valorNominal:  valorValorNominal,
							taxaDeJuros:   valorTaxaDeJuros,
							numeroPeriodo: valorNumeroPeriodo,
						}

						fmt.Printf("O valor do Desconto é: R$ %.2f\n", instanciaDescontoSimples.calcularDescontoSimples())
					}
				case 2:
					{
						fmt.Println("Você entrou no menu Valor Nominal!")
						fmt.Println("Digite o Desconto Racional:")
						fmt.Scanln(&valorDescontoSimples)
						fmt.Println("Digite a Taxa de Juros (decimal):")
						fmt.Scanln(&valorTaxaDeJuros)
						fmt.Println("Digite o Período de Antecipação (meses):")
						fmt.Scanln(&valorNumeroPeriodo)

						instanciaValorNominal := DescontosSimples{
							descontoSimples: valorDescontoSimples,
							taxaDeJuros:     valorTaxaDeJuros,
							numeroPeriodo:   valorNumeroPeriodo,
						}

						fmt.Printf("O valor do Valor Nominal é: R$ %.2f\n", instanciaValorNominal.calcularValorNominal())
					}
				case 3:
					{
						fmt.Println("Você entrou no menu Taxa de Juros!")
						fmt.Println("Digite o Desconto Racional Simples:")
						fmt.Scanln(&valorDescontoSimples)
						fmt.Println("Digite o Valor Nominal:")
						fmt.Scanln(&valorValorNominal)
						fmt.Println("Digite o Período de Antecipação (meses):")
						fmt.Scanln(&valorNumeroPeriodo)

						instanciaTaxaDeJuros := DescontosSimples{
							descontoSimples: valorDescontoSimples,
							valorNominal:    valorValorNominal,
							numeroPeriodo:   valorNumeroPeriodo,
						}

						fmt.Printf("O valor da Taxa de Juros é: %.4f\n", instanciaTaxaDeJuros.calcularTaxaDeJuros())
					}
				case 4:
					{
						fmt.Println("Você entrou no menu Período de Antecipação!")
						fmt.Println("Digite o Desconto Racional Simples:")
						fmt.Scanln(&valorDescontoSimples)
						fmt.Println("Digite o Valor Nominal:")
						fmt.Scanln(&valorValorNominal)
						fmt.Println("Digite a Taxa de Juros (decimal):")
						fmt.Scanln(&valorTaxaDeJuros)

						instanciaNumeroPeriodo := DescontosSimples{
							descontoSimples: valorDescontoSimples,
							valorNominal:    valorValorNominal,
							taxaDeJuros:     valorTaxaDeJuros,
						}

						fmt.Printf("O valor do Período de Antecipação é: %d\n", instanciaNumeroPeriodo.calcularNumeroPeriodo(), "meses")
					}
				case 5:
					{
						continue
					}
				default:
					{
						fmt.Println("Essa opção não é válida, tente outra!")
					}
				}
			}
		case 3:
			{
				fmt.Println("Você encerrou o programa!")
				loop = false
			}
		default:
			{
				fmt.Println("Essa opção não é válida, tente outra!")
			}
		}
	}
}

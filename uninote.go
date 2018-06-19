// Copyright 2018 Igor Brandao aka [isca]. All rights reserved.
// Use of this source code is governed by a LGPL-3.0
// license that can be found in the LICENSE file.

// What this code must do?
// > This code simple calculate your test Grades for Uninter University.
//
package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	//Apol eh o array publico que recebe as notas das apols em seus indices
	Apol [6]float64

	//Ap recebe as notas de atividade pratica e prova discursiva
	Ap [2]float64

	//Po recebe nota da prova objetiva
	Po float64

	//Colors

	//Red recebe a cor vermelha
	Red = color.New(color.FgRed).PrintfFunc()

	//Green recebe a cor verde
	Green = color.New(color.FgGreen).PrintfFunc()

	//Bwhite recebe a cor bold branco
	Bwhite = color.New(color.Bold, color.FgWhite).PrintfFunc()
)

//ColorNota apresenta o resultado da media colorido
func ColorNota(a float64) float64 {
	if a < 60 {
		Red("%v\r\n\n", a)
	} else if a >= 60 {
		Green("%v\r\n\n", a)
	}
	return a
}

//Nota das Apol 1,2,3,4 e 5
func apolMedia(a, b, c, d, e float64) float64 {
	media := (a + b + c + d + e) / 5
	Bwhite("Sua media das apols foi = ")
	return ColorNota(media)
}

//Nota da Atividade Pratica + Nota da Prova Discursiva
func apComPd(a, b float64) float64 {
	media := (a*4 + b*6) / 10
	Bwhite("Sua media da Prova Discursiva com Atividade Pratica foi = ")
	return ColorNota(media)

}

//Nota da Prova Objetiva + Media das Apols + Media da apComPd
func mediaGeral(a, b, c float64) float64 {
	media := (a*30 + b*20 + c*50) / 100
	Bwhite("Sua media Geral nesta materia eh de: ")
	return ColorNota(media)
}

//Sample reading stdin string with bufio
//		reader := bufio.NewReader(os.Stdin)
//		fmt.Printf("Sua nota na APOL2: ")
//		apol1, _ := reader.ReadString('\n')

//Receiveing apols notes on loop
func apolGrade() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Entre com a nota na APOL%v: ", i)
		fmt.Scan(&Apol[i])
	}
	return
}

func apComPdGrade() {
	for i := 1; i <= 2; i++ {
		if i == 1 {
			fmt.Printf("Entre sua nota de Atividade Pratica: ")
			fmt.Scan(&Ap[0])
		} else if i == 2 {
			fmt.Printf("Entre sua nota da Prova discursiva: ")
			fmt.Scan(&Ap[1])
		}
	}
}

func provaObjetiva() {
	fmt.Printf("Nota prova objetiva: ")
	fmt.Scan(&Po)
	return
}

func main() {

	//Processando notas das Apols
	apolGrade()
	n3 := apolMedia(Apol[1], Apol[2], Apol[3], Apol[4], Apol[5])

	//Processando notas da AP + PD
	apComPdGrade()
	n4 := apComPd(Ap[0], Ap[1])

	//Processando a  media da materia
	provaObjetiva()
	mediaGeral(Po, n3, n4)

}

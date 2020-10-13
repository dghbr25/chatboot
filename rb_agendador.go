package main

import (
	_ "fmt"
	"time"
)

func agendador() {
	var i int = 0
	for {
		time.Sleep(3 * time.Second)
		if i == 0 {
			sendMsg("5521990196631", "Você sabia que já existe o Título Eleitoral em App?")

			i = 1
		} else if i == 1 {
			sendMsg("5521990196631", "Esse ano Teremos Eleições On-line")

			i = 2
		} else if i == 2 {
			sendMsg("5521990196631", "Já escolheu seu candidato? Já sabe quem está fazendo o que ? Baixe nosso App")

			i = 3
		} else if i == 3 {
			sendMsg("5521990196631", "*(((((((9999999999)))))))* vote com o coração Rogerio Lisboa para Prefeito de Nova Iguacu")

		}

		/* _, err := getPacientesAssintomaticos()
		if err != nil {
			fmt.Printf("erro teste\n\n\n", err)
			return
		}
		*/
	}
}

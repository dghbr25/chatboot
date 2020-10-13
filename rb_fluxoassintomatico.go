package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func postRetiraPacientesFila(id string) (bool, error) {

	_, err := http.PostForm("http://covidbot.com.br/covidapi/paciente/notificacao/notpacientes/atualiza", url.Values{"id": {id}})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return false, err
	}

	return true, nil

}

func getPacientesAssintomaticos() (bool, error) {

	var url string = "http://covidbot.com.br/covidapi/paciente/notificacao/notpacientes"
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error body: %v\n", err)
		return false, err
	}
	var pacientesList []pAssintMsg
	if err := json.Unmarshal(body, &pacientesList); err != nil {
		fmt.Fprintf(os.Stderr, "error json: %v\n", err)
		return false, err
	}
	for i := 0; i < len(pacientesList); i++ {
		var numero = pacientesList[i].Num_telefone
		sendMsg(numero, "Olá "+pacientesList[i].Nome+".\n"+pacientesList[i].Texto)
		resp, err := postRetiraPacientesFila(pacientesList[i].Id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao retirar paciente da fila de msg: %v\n", err)
			return false, err
		}
		fmt.Printf("Menssagem enviada para paciente\n"+pacientesList[i].Nome, resp)
		globalContatos3[numero] = &pAssintMsg{"", "", "", ""}
		delete(globalContatos3, numero)
		time.Sleep(11 * time.Second)
	}
	return true, nil
}

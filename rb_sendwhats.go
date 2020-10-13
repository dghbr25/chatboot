package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
)

type JSONReturno struct {
	Probabilidade string `json:"probabilidade"`
}

func sendDataClouding(contato *Contato) (bool, error) {

	if contato.sintomasI == "" {
		contato.sintomasI = "6"
	}

	resp, err := http.PostForm("http://easystop.com.br/apicorona/cadastro/pacientes", url.Values{"nomePaciente": {contato.nome}, "cpf": {"000.000.000-00"}, "idade": {contato.idade}, "dt_nascimento": {""},
		"sintomas_primarios": {contato.sintomasI}, "sintomas_secundarios": {contato.sintomasII}, "contato_infectado": {contato.infect}, "gravidez": {contato.gravida},
		"isfebre": {contato.febre}, "ismelhoramedic": {contato.melhora}, "tmpfebre": {contato.tempoFebre}, "ultima_medicao_febre": {contato.temperatura}, "num_telefone": {contato.numero}, "sexo": {contato.sexo}, "cep": {contato.cep},
		"viagem_internacional": {contato.viagem}})

	fmt.Printf("nome: %s\n", contato.nome)
	fmt.Printf("sexo: %s\n", contato.sexo)
	fmt.Printf("idade: %s\n", contato.idade)
	fmt.Printf("cep: %s\n", contato.cep)
	fmt.Printf("stI: %s\n", contato.sintomasI)
	fmt.Printf("stII: %s\n", contato.sintomasII)
	fmt.Printf("febre: %s\n", contato.febre)
	fmt.Printf("tempo_febre: %s\n", contato.tempoFebre)
	fmt.Printf("temp: %s\n", contato.temperatura)
	fmt.Printf("infect: %s\n", contato.infect)
	fmt.Printf("viagem: %s\n", contato.viagem)
	fmt.Printf("remedios: %s\n", contato.remedio)
	fmt.Printf("numero: %s\n", contato.numero)
	fmt.Printf("ismelhoramedic: %s\n\n\n", contato.melhora)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error body: %v\n", err)
		return false, err
	}

	var jsonReturno JSONReturno
	if err := json.Unmarshal(body, &jsonReturno); err != nil {
		fmt.Fprintf(os.Stderr, "error json: %v\n", err)
		return false, err
	}

	if strings.ToLower(jsonReturno.Probabilidade) == "baixa" {
		sendMsg(contato.numero, QR00B)
		sendMsg(contato.numero, QRORB)
	} else if strings.ToLower(jsonReturno.Probabilidade) == "media" {
		sendMsg(contato.numero, QR00M)
		sendMsg(contato.numero, QRORM)
	} else if strings.ToLower(jsonReturno.Probabilidade) == "alta" {
		sendMsg(contato.numero, QR00A)
		sendMsg(contato.numero, QRORA)
	}

	delete(globalContatos, contato.numero)

	return true, nil
}

func sendMsg(numero string, texto string) (bool, error) {

	manda := fmt.Sprintf("%v@s.whatsapp.net", numero)
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: manda,
		},
		Text: texto,
	}

	_, err := globalWac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return false, err
	}

	return true, nil
}

type TemplateMessage struct {
	Info             MessageInfo
	HydratedTemplate *proto.HydratedFourRowTemplate
	Format           interface{}
	ContextInfo      ContextInfo
}

func getTemplateMessageProto(msg TemplateMessage) *proto.WebMessageInfo {

	p := getInfoProto(&msg.Info)

	contextInfo := getContextInfoProto(&msg.ContextInfo)

	if template, ok := msg.Format.(*proto.TemplateMessage_HydratedFourRowTemplate); ok {
		p.Message = &proto.Message{
			TemplateMessage: &proto.TemplateMessage{
				HydratedTemplate: msg.HydratedTemplate,
				Format:           template,
				ContextInfo:      contextInfo,
			},
		}
	}

	if template, ok := msg.Format.(*proto.TemplateMessage_FourRowTemplate); ok {
		p.Message = &proto.Message{
			TemplateMessage: &proto.TemplateMessage{
				HydratedTemplate: msg.HydratedTemplate,
				Format:           template,
				ContextInfo:      contextInfo,
			},
		}
	}
	return p
}
func sendPDF(numero string, arquivo string, titulo string) (bool, error) {
	pdf, err := os.Open(arquivo)
	if err != nil {
		return false, err
	}

	manda := fmt.Sprintf("%v@s.whatsapp.net", numero)
	msg := whatsapp.DocumentMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: manda,
		},
		Type:    "application/pdf",
		Title:   titulo,
		Content: pdf,
	}

	_, err = globalWac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return false, err
	}

	return true, nil
}
func sendImage(numero string, arquivo string, caption string) (bool, error) {
	img, err := os.Open(arquivo)
	if err != nil {
		return false, err
	}

	manda := fmt.Sprintf("%v@s.whatsapp.net", numero)
	msg := whatsapp.ImageMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: manda,
		},
		Type:    "image/jpeg",
		Caption: caption,
		Content: img,
	}

	_, err = globalWac.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return false, err
	}

	return true, nil
}

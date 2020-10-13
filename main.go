package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Rhymen/go-whatsapp"

	//"io"
	"encoding/json"
	_ "encoding/json"
	"os"
	_ "strconv"
	"time"

	_ "github.com/asaskevich/govalidator"
)

type waHandler struct{}

type politica struct {
	eleitor      string
	sexo         string
	faixa_etaria string
}

type Contato struct {
	passwd      string
	admin01     string
	caminho     int
	nome        string
	sexo        string
	idade       string
	remedio     string
	febre       string
	tempoFebre  string
	temperatura string
	infect      string
	melhora     string
	sintomasI   string
	sintomasII  string
	viagem      string
	gravida     string
	numero      string
	cep         string
}

type pAssint struct {
	Idpacientes          string `json:"Idpacientes"`
	Nome                 string `json:"nome"`
	Cpf                  string `json:"cpf"`
	Dt_nascimento        string `json:"Dt_nascimento"`
	Sintomas_primarios   string `json:"sintomas_primarios"`
	Sintomas_secundarios string `json:"sintomas_secundarios"`
	Contato_infectado    string `json:"contato_infectado"`
	Gravidez             string `json:"gravidez"`
	Isfebre              string `json:"isfebre"`
	Ultima_medicao_febre string `json:"ultima_medicao_febre"`
	Num_telefone         string `json:"num_telefone"`
	Sexo                 string `json:"Sexo"`
	Dt_Inclusao          string `json:"Dt_Inclusao"`
	CEP                  string `json:"CEP"`
	Viagem_internacional string `json:"viagem_internacional"`
	Tmpfebre             string `json:"tmpfebre"`
	Idade                string `json:"idade"`
	Ismelhoramedic       string `json:"ismelhoramedic"`
	Probabilidade        string `json:"probabilidade"`
}

type pAssintMsg struct {
	Id           string `json:"Id"`
	Nome         string `json:"Nome"`
	Texto        string `json:"Texto"`
	Num_telefone string `json:"num_telefone"`
}

var globalContatos = make(map[string]*Contato)
var globalContatos2 = make(map[string]*pAssint)
var globalContatos3 = make(map[string]*pAssintMsg)
var globalContatosPolitica = make(map[string]*pAssintMsg)

var globalAtual uint64
var globalWac *whatsapp.Conn

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func (*waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	if strings.Index(message.Info.RemoteJid, "-") != -1 {
		return
	}

	numero := message.Info.RemoteJid[:strings.Index(message.Info.RemoteJid, "@")]
	if message.Info.FromMe == false && message.Info.Timestamp > globalAtual {
		/*
			&& numero[:2] == "55"
				if numero != "5521969536121" && numero != "5521976676414" && numero != "5521980901502" && numero != "5521983062499" && numero != "5521996493601" && numero != "5521991821884" {
					return
				}
		*/
		fmt.Printf(" %v (%v/%v/%v): %v\n", message.Info.PushName, message.Info.RemoteJid, message.Info.Id, message.Info.Timestamp, message.Text)

		if _, ok := globalContatos[numero]; ok {
			if globalContatos[numero].caminho == 0 {
				// setando nome
				globalContatos[numero].nome = message.Text
				// informe a idade
				sendMsg(numero, politica02)
				globalContatos[numero].caminho = 1
			} else if globalContatos[numero].caminho == 1 {
				sendMsg(numero, politica03)
				globalContatos[numero].caminho = 2
			} else if globalContatos[numero].caminho == 2 {
				sendMsg(numero, politica31)
				globalContatos[numero].caminho = 3
			} else if globalContatos[numero].caminho == 3 {
				sendMsg(numero, politica04)
				globalContatos[numero].caminho = 4
			} else if globalContatos[numero].caminho == 4 {
				sendMsg(numero, politica05)
				globalContatos[numero].caminho = 5
			} else if globalContatos[numero].caminho == 5 {
				sendMsg(numero, politica06)
				globalContatos[numero].caminho = 6
			} else if globalContatos[numero].caminho == 6 {
				sendMsg(numero, politica07)
				sendMsg(numero, politica08)
				globalContatos[numero].caminho = 6
			}
		} else {
			sendMsg(numero, "*Bem-vindo ao Assistente Virtual Eleitoral Brasileiro*!")
			sendMsg(numero, politica01)
			sendMsg(numero, politica11)

			/*
				caminho int
				nome string
				sexo string
				idade string
				remedio string
				febre string
				tempoFebre string
				temperatura string
				infect string
				melhora string
				sintomasI string
				sintomasII string
				viagem string
				gravida string
				numero string
				cep string
			*/
			globalContatos[numero] = &Contato{"", "", 0, "", "", "", "", "", "", "", "", "", "", "", "", "", numero, ""}
		}
	}
}

func (*waHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	return
}

func (*waHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	return
}

func (*waHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	return
}

func (*waHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	return
}

func main() {

	globalAtual = uint64(time.Now().Unix())
	wac, err := whatsapp.NewConn(15 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	} else {
		globalWac = wac
	}
	wac.AddHandler(&waHandler{})
	err = login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	//go agendador()

	//go getbaseConhecimento("")

	select {}

}
func (*waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error occoured: %v", err)
}

func getbaseConhecimento(numeroRobo string) (bool, error) {

	// define data structure
	type Robo struct {
		ID        int    `json:"id_Robo"`
		Nome      string `json:"Nome"`
		Descricao string `json:"Descricao"`
	}

	type Numero struct {
		Num_telefone string `json:"num_telefone"`
		Id_Robo      int    `json:"id_Robo"`
	}

	type Fluxo struct {
		Id_Fluxo   int    `json:"id_Fluxo"`
		Id_Robo    int    `json:"id_Robo"`
		Nome       string `json:"Nome"`
		Descricao  string `json:"Descricao"`
		Dt_Criacao string `json:"Dt_Criacao"`
	}

	type LogicaFP struct {
		Id_Fluxo     int `json:"id_Fluxo"`
		Id_Pergunta  int `json:"id_Pergunta"`
		Seq_Pergunta int `json:"Seq_Pergunta"`
	}

	type Orientacao struct {
		Idorientacao     int    `json:"idorientacao"`
		Id_Fluxo         int    `json:"id_Fluxo"`
		Texto_orientacao int    `json:"texto_orientacao"`
		Tipo             string `json:"Tipo"`
		Sequencia        int    `json:"Sequencia"`
	}

	type OpcaoResposta struct {
		id_Resposta       int    `json:"id_Resposta"`
		id_Pergunta       int    `json:"id_Pergunta"`
		Texto             string `json:"Texto"`
		Sequencia         string `json:"Sequencia"`
		idProximaPergunta int    `json:"idProximaPergunta"`
		Dicionario        string `json:"Dicionario"`
	}

	type Pergunta struct {
		id_Pergunta        int    `json:"id_Pergunta"`
		Texto              string `json:"Texto"`
		PreExplicacao      string `json:"PreExplicacao"`
		PosExplicacao      string `json:"PosExplicacao"`
		Nome_Variavel_Post string `json:"Nome_Variavel_Post"`
	}

	robo, err := ioutil.ReadFile("arq/robo.json")
	numero, err := ioutil.ReadFile("arq/numero.json")
	fluxo, err := ioutil.ReadFile("arq/fluxo.json")
	logicafp, err := ioutil.ReadFile("arq/logicafluxo.json")
	orientacao, err := ioutil.ReadFile("arq/orientacao.json")
	opcaoresp, err := ioutil.ReadFile("arq/opcaoresposta.json")
	pergunta, err := ioutil.ReadFile("arq/pegunta.json")

	if err != nil {
		fmt.Print(err)
	}

	// json data

	var roboObj []Robo
	var numeroObj []Numero
	var fluxolist []Fluxo
	var logicafplist []LogicaFP
	var orientacaolist []Orientacao
	var opcapresplist []OpcaoResposta
	var perguntalist []Pergunta
	// Verifico erro no json

	// robo
	err = json.Unmarshal(robo, &roboObj)
	if err != nil {
		fmt.Println("Erro json robo:", err)
	}

	// numero
	err = json.Unmarshal(numero, &numeroObj)
	if err != nil {
		fmt.Println("Erro json numero:", err)
	}

	//fluxo
	err = json.Unmarshal(fluxo, &fluxolist)
	if err != nil {
		fmt.Println("Erro json fluxo:", err)
	}

	//logica fluxo principal
	err = json.Unmarshal(logicafp, &logicafplist)
	if err != nil {
		fmt.Println("Erro json fluxo:", err)
	}

	//Orientacao
	err = json.Unmarshal(orientacao, &orientacaolist)
	if err != nil {
		fmt.Println("Erro json fluxo:", err)
	}

	// opcao resposta
	err = json.Unmarshal(opcaoresp, &opcapresplist)
	if err != nil {
		fmt.Println("Erro json robo:", err)
	}

	// pergunta
	err = json.Unmarshal(pergunta, &perguntalist)
	if err != nil {
		fmt.Println("Erro json robo:", err)
	}

	/* can access using struct now
	fmt.Printf("Robo   :\n", roboObj)
	fmt.Printf("Numero :\n", numeroObj)
	fmt.Printf("Fluxo  :\n", fluxolist)
	fmt.Printf("Logica FP  :\n", logicafplist)
	fmt.Printf("Orientacao  :\n", orientacaolist)
	fmt.Printf("Opcao resposta  :\n", opcapresplist)
	fmt.Printf("pergunta  :\n", perguntalist)
	*/
	return true, nil
}

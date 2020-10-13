package main

var sinterr string = "*Por Favor selecione uma ou mais da opções acima separadas por ,*\n*Exemplo: 1, 2, 3, 4, 5*"
var retor string = "Desculpe, nao entendi sua resposta. Digite uma das opcoes abaixo para prosseguir."
var nome string = "Me diga seu nome:"
var idade string = "Qual a sua idade:"
var sexo string = "Me diga seu sexo:\n*F* - Feminino\n*M* - Masculino"
var QS000 string = "Teve algum sintoma de febre?\n*1* - Sim\n*2* - Nao"
var QS001 string = "Agora fale um pouco mais sobre a sua febre. Faz tempo que você está com febre?\n*1* - Dois dias ou menos.\n*2* - Mais de 3 dias."
var QN001 string = "Você fez o uso de algum medicamento?\n*1* - Sim\n*2* - Não"
var QS002 string = "Na sua medição, qual foi a temperatura mais alta?\n*1* -  Menos de 37,8.\n*2* - Entre 37,8 e 39.\n*3* - Acima de 39.\n*4* - Não consegui medir."
var QR001 string = "Os sintomas abaixo acontecem no início de resfriados leves. Você sente algum(ns)?  \n(coloque os números referentes aos seus sintomas, separando-os por vírgula)\n*1* - Espirro\n*2* - Nariz escorrendo (coriza)\n*3* - Congestão nasal(obstrução nasal)\n*4* - Nenhum dos sintoma acima."
var QR002 string = "Teve contato com alguém que foi infectado pelo Covid-19?\n*1* - Sim\n*2* - Não"
var QR003 string = "Notou alguma melhora, ao tomar o medicamento?\n*1* - Sim\n*2* - Não"
var QR004 string = "Você esteve recentemente em locais onde existam mais casos confirmados de coronavírus do que a sua própria região?\n*1* - Sim\n*2* - Não"
var QR005 string = "Preciso saber mais sobre outros sintomas, um pouco mais específicos. Lembra se teve:\n(coloque os números referentes aos seus sintomas, separando-os por vírgula) \n*1* - Falta de ar \n*2* - Tosse seca \n*3* - Tosse com secreção \n*4* - Dor no corpo(mialgia) \n*5* - Dor nas juntas(dor articular) \n*6* - Dor de cabeça \n*7* - Dor de garganta \n*8* - Nenhum dos sintomas acima."
var QR007 string = "Você está grávida?\n*1* - Sim\n*2* - Não"
var QRCOVID1901 string = "Bom, pelo que você me contou identificamos alguns sinais de alarme e é possivel que você esteja contaminado com o Covid-19 (Corona Virus). Minha orientação é que você busque atendimento em uma Unidade de Saúde perto de casa para ser avaliado por um profissional. Também recomento que você siga algumas orientações durante o seu trajeto até a Unidade de Saúde, para você se proteger e evitar transmitir o vírus para outras pessoas."
var QR008 string = "Acho que você deveria entrar em contato com um médico, ou se dirigir a uma unidade de saúde/hospital próximo à sua casa. Mas para não comprometer ainda mais sua saúde e de outras pessoas, evite ir de transporte público(se possível). Use uma máscara protetora(Caso não possua), cubra a boca com um pano para evitar a possibilidade de transmitir o vírus para outras pessoas."
var QR009 string = " * Utilize uma máscara em seu rosto durante todo o trajeto, desde a saída da sua casa até chegar na Unidade de Saúde*\n* Quando tossir e espirar, cubra a boca e o nariz com um lenço de papel descartável e jogue-o no lixo, ou proteja o com o antebraço, Não use suas mãos, a fim de evitar transmitir o vírus para objetos e pessoas.*\n* Lave as mãos com agua e sabão por pelo menos 20 segundos, com frequência. Quando não for possível, limpe e higienize suas mãos com álcoolgel 70%*\n\n*Evite contato próximo com pessoas que estejam com sintomas de resfriado ou gripe. *\n\n *Evite contato próximo com idósos, grávidas, mulheres que estejam amamentando ou pessoas que adoeçam com facilidade.*\n\n*Não compartilhe objetos de uso pessoal, como copos, talheres e pratos, e limpe objetos que sejam tocados com frequẽncia, como celulares.* "
var QR010 string = "Avise um familiar próximo da situação e sinalize que está se encaminhando para atendimento médico."
var QRLL1 string = "Você gostaria de ver unidades de saúde próximas? \n*1* - Sim\n*2* - Não"
var QS003 string = "Favor compartilhe sua localização para informar as Unidades de Saúde mais próxima. "
var QS004 string = "Favor informe o cep de sua residencia!"
var QRF01 string = "Obrigado. Estamos finalizando o ChatBoot e sua participacao e muito importante nesse momento de testes."

var QR00A string = "Falta de ar ou a persistência da febre por 3 dias ou mais, principalmente quando a febre não melhora com medicamento são sintomas de alerta para suspeita de coronavírus e por isso você deve fazer uma avaliação detalhada.  Você pode esclarecer suas dúvidas com profissionais de saúde pelo tel 0800 644 6543, de segunda à sexta das 8h às 17h30. Se estiver com falta de ar, não perca tempo, vá até um serviço de urgência mais próximo ou disque para o SAMU - 192"
var QRORA string = "\n\n- Utilize uma máscara em seu rosto durante todo o trajeto, desde a saída da sua casa até chegar na Unidade de Saúde;\n\n- Quando tossir e espirrar, cubra a boca e o nariz com um lenço de papel descartável e jogue-o no lixo, ou proteja com o antebraço. Não use suas mãos, a fim de evitar transmitir o vírus para objetos e pessoas;\n\n- Lave as mãos com água e sabão por pelo menos 20 segundos, com frequência. Quando não for possível, limpe e higienize suas mãos com álcool gel 70%;\n\n- Evite contato próximo com pessoas que estejam com sintomas de resfriado ou gripe;\n\n- Evite contato próximo com idosos, grávidas, mulheres que estejam amamentando ou pessoas que adoeçam com facilidade;\n\n- Não compartilhe objetos de uso pessoal, como copos, talheres e pratos, e limpe objetos que sejam tocados com frequência, como celulares."

var QR00M string = "Seu caso tem alguns indicativos que você apresente um quadro gripal leve. Você deve manter isolamento domiciliar e ficar tranquilo.  Caso tenha duvidas,  esclareça com profissionais de saúde, médicos e enfermeiros da atenção primária do SUS, pelo tel 0800 644 6543, de segunda à sexta das 8h às 17h30. Refaça seu teste amanhã para sabermos como você está evoluindo."
var QRORM string = "\n\n- Verifique se os sintomas que você sinalizou, ou que a febre que sentiu retornaram e ou estão piorando;\n\n- Evite dirigir-se aos hospitais/unidades de saúde, já que você não está com grande probabilidade de estar infectado;\n\n- Lave as mãos com água e sabão por pelo menos 20 segundos, com frequência. Quando não for possível, limpe e higienize suas mãos com álcool gel 70%;\n\n- Evite encostar as mãos não lavadas nos olhos, boca e nariz;\n\n- Quando tossir e espirrar, cubra a boca e o nariz com um lenço de papel descartável e jogue-o no lixo, ou proteja com o antebraço. Não use suas mãos, a fim de evitar transmitir o vírus para objetos e pessoas;\n\n- Evite locais muito cheios e com aglomeração de pessoas (shows, festas, shoppings etc);\n\n- Não compartilhe objetos de uso pessoal, como copos, talheres e pratos, e limpe objetos que sejam tocados com frequência, como celulares;\n\n- Evite contato próximo com pessoas que estejam com sintomas de resfriado ou gripe;"

var QR00B string = "Você tem baixa suspeita de infecção pelo coronavírus. Isso é ótimo! Mas pedimos que faça esse teste, todos os dias, durante o período de afastamento social. Você pode esclarecer suas dúvidas com profissionais de saúde pelo tel 0800 644 6543, de segunda à sexta das 8h às 17h30."
var QRORB string = "\n\n- Lave as mãos com água e sabão por pelo menos 20 segundos, com frequência. Quando não for possível, limpe e higienize suas mãos com álcool gel 70%;\n\n- Evite encostar as mãos não lavadas nos olhos, boca e nariz;\n- Quando tossir e espirrar, cubra a boca e o nariz com um lenço de papel descartável e jogue-o no lixo, ou proteja com o antebraço. Não use suas mãos, a fim de evitar transmitir o vírus para objetos e pessoas;\n\n- Evite locais muito cheios e com aglomeração de pessoas (shows, festas, shoppings etc);\n\n- Não compartilhe objetos de uso pessoal, como copos, talheres e pratos, e limpe objetos que sejam tocados com frequência, como celulares;\n\n- Evite contato próximo com pessoas que estejam com sintomas de resfriado ou gripe; \n\nHidrate-se, consuma no mínimo 2 litros de água por dia\nInclua frutas e verduras na sua alimentação\n\nTome sol, de 10 a 15 minutos por dia."

var QRMOR string = "Para você continuar saudável, eu recomendo que você siga as orientações a seguir."
var QRAOR string = "*Algumas orientações.*"
var QRCEP string = "Para que saiba quais são os bairros com mais pessoas com sintomas de refriados e gripes, onde consequentemente a ajuda é mais necessária, preciso que me informe o seu CEP:"

var QRAD1 string = "*SELECIONE A OPÇÃO DESEJADA*\n1 - Menssagem desativar o sistema.\n2 - Menssagem ativar  do sistema\n3 - Menssagem para pacientes"
var QRAD2 string = "Deseja enviar menssagens para pacientes com:\n1 - Baixa probabilidade\n2 - Média probabilidade\n3 - Alta probabilidade"

var politica01 string = "*Me chamo Lupa e estou aqui pra te ouvir.*"
var politica11 string = "Gostaria de 2 minutos de sua atenção pra podermos juntos construir uma Mesquita melhor. "
var politica02 string = "Faixa Etária:\n1) 16 a 24 anos \n2) 25 a 34 anos\n3) 35 a 44 anos\n4) 45 a 59 anos\n5) 60 anos ou mais"
var politica03 string = "Escolaridade:\n1) Ensino Fundamental (Pré à oitava série) (1º Grau)\n2) Ensino Médio completo ou incompleto (2º Grau)\n3) Ensino Superior completo ou incompleto"
var politica31 string = "Qual seu compromentimento com a política de sua cidade?\n1) Grande \n2) Medio\n3) Pequeno"
var politica04 string = "Se as eleições para Prefeito(a) de Mesquita fossem hoje, em quem o Sr(a) votaria?\n*DIGITE O NOME DO CANDIDATO OU SELECIONE AS OPÇÕES ABAIXO:*\n1) Não sabe\n2) Ninguém"
var politica05 string = "Se as eleições para Prefeito(a) do Mesquita fossem hoje e os(as) candidatos(as) fossem esses(as),\nem quem o Sr(a) votaria?\n1) Cris Gêmeas2) Jorge Miranda\n3) André Medico\n4) Daivid Luz\n5) Farid Abraão\n"
var politica06 string = "*AVALIAÇÃO DA ADMINISTRAÇÃO ESTADUAL*\n----------------------------------------------------------------------------------------------------------------------------------------------------\nEm sua opinião, a administração do Governador do Rio de Janeiro, Wilson Witzel, está sendo ótima, boa, regular, ruim ou péssima?\n1) Ótima\n2) Boa\n3) Regular\n4) Ruim\n5) Péssima\n6) Não sabe\n7) não opinar"
var politica07 string = "Olá Chegamos ao fim! Estou processando suas informações. Gostei muito de te conhecer. Temos muito mais informações importantes pra voce sobre sua cidade. Caso deseje mais infomações selecione uma das opções abaixo:\n1) - Sim\n2) - Não"
var politica08 string = "E não esqueça de lavar as mãos e usar mascaras segundo os orgãos de saúde. Vidas em primeiro lugar"

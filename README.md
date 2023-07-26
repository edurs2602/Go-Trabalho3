# Trabalho 3: Desenvolvimento de API RESTful

## Este Trabalho foi realizado em dupla

### Juntamente de [Pablo Santos](https://github.com/BiscuI)

<br>

# Tarefa:

<sub>Última atualização: 26/07/2023</sub>

## Sumário

- [Objetivos](#objetivo)
- [Tarefas](#tarefas)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Objetivo

O objetivo deste trabalho é colocar em prática o desenvolvimento de aplicações cliente-servidor utilizando APIs
RESTful, isto é, em conformidade com o estilo arquitetural REST (Representational State Transfer). São explorados
os seguintes elementos da programação em Go, cujos conhecimentos são, portanto, ora necessários:

- Entrada e saída formatada via console
- Tipos estruturados (_structs_)
- Funções
- Pacotes e módulos
- Manipulação de erros
- _Framework_ de testes
- Suporte ao desenvolvimento cliente-servidor

## Tarefas

As tarefas principais a serem realizadas neste trabalho é a implementação, na linguagem de programação Go, de
uma aplicação cliente-servidor fim-a-fim. Isto significa que deverão ser implementados (i) um programa servidor que
fornece funcionalidades/serviços na forma de uma API RESTful e (ii) a implementação de um programa cliente que
faz chamadas às funções disponibilizadas por essa API. Por questões de simplicidade, o programa servidor pode
executar em uma máquina local (localhost).

Para o programa servidor, as funções implementadas na API são de livre escolha, devendo, contudo, atender às
seguintes restrições: (i) estarem relacionadas a um objetivo de negócio que seja relevante e de complexidade consi-
derada razoável; (ii) oferecerem pelo menos cinco funções capazes de serem acessadas por meio de requisições
HTTP GET e POST, e; (iii) as requisições GET devem ser passíveis do envio de parâmetros na requisição. Op-
cionalmente, as funcionalidades oferecidas podem fazer uso de serviços externos providos por APIs públicas. Caso
as operações envolvam persistência, isto é, armazenamento de dados, não é obrigatório fazer uso de um banco de
dados típico, bastando, por exemplo, utilizar um arquivo existente no lado servidor emulando a ideia de uma base
de dados.

Com relação ao programa cliente, este deverá oferecer uma interface, através da entrada padrão (console) ou via
linha de comando, que possibilite a um usuário fornecer dados e fazer uso das operações disponibilizadas pela
API. Além disso, o programa cliente deve tratar as respostas enviadas pelo programa servidor como resultado do
processamento das requisições. O programa cliente deve ainda exibir na saída padrão mensagens resultantes do
processo de envio de requisição e tratamento de resposta.

A implementação do programa servidor correspondente à API deverá ter todas as suas operações devidamente
testadas utilizando o framework de testes provido pela linguagem Go. Portanto, deverá ser necessariamente im-
plementado um código fonte em separado contendo funções de testes unitários.

Extra: Será concedida pontuação adicional (q.v. Avaliação) à documentação da API feita com suporte do Swagger.
O Swagger é um conjunto de ferramentas de código aberto que possibilita o desenvolvimento e documentação de
APIs, além de possibilitar usuários interagirem diretamente com a API para testar suas funcionalidades de maneira
fácil. Esta tarefa não é obrigatória.

## Requisitos

A implementação deste trabalho requer os seguintes elementos instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre os integrantes da equipe. Após a implementação das soluções para os problemas propostos, o arquivo [`author.md`](author.md) presente no repositório deverá ser editado preenchendo as informações de identificação dos integrantes da equipe, na seção [Informações de Autoria](author.md#identificação-de-autoria).

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. **Trabalhos copiados no todo ou em parte de outros colegas ou da Internet serão anulados e receberão nota zero.**

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://git-scm.com) serão utilizados para possibilitar a entrega da implementação realizadas. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (_assignment_) no GitHub Classroom. Cada integrante de equipe deverá acessar este [_link_], aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelos integrantes da equipe e pelo docente, sob a organização [`ufrn-golang`](https://github.com/ufrn-golang). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo do arquivo executável gerado a partir da compilação do código fonte.

A implementação do programa objeto deste trabalho deverá ser realizada **até as 11:00 do dia 26 de junho de 2023** no respectivo repositório Git da equipe. Para fins de registro, o endereço do repositório também deverá ser **obrigatoriamente** enviado através da opção _Tarefas_ na Turma Virtual do SIGAA, devendo **um único membro da equipe** realizar esse envio. Além disso, **não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.**

## Avaliação

A avaliação do programa implementado contabilizará nota de até 10,0 pontos na 3ª unidade da disciplina. O
programa implementado será avaliado de acordo com os seguintes critérios:

- grau de criatividade e complexidade do objetivo de negócio da aplicação;
- utilização correta dos recursos providos pela linguagem de programação Go;
- corretude da execução do programa implementado;
- realização de testes unitários sobre as operações disponibilizadas pela API;
- aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte;
- correta utilização do repositório Git, incluindo documentação adequada por meio do arquivo README.md e o
  registro de todo o histórico da implementação por meio de commits.

O não cumprimento de algum dos critérios de avaliação especificados poderá resultar nos seguintes decréscimos,
aplicados sobre a nota obtida até então na avaliação:

| Falta                                                                                                       | Decréscimo |
| :---------------------------------------------------------------------------------------------------------- | ---------: |
| Falta de comentários no código fonte                                                                        |       -10% |
| Uso inadequado de controle de versão com Git                                                                |       -20% |
| Falta de especificação ou especificação incorreta da autoria do trabalho (arquivo [`author.md`](author.md)) |       -20% |
| Código fonte com legibilidade prejudicada (por exemplo, com identação ou nomenclatura inadequada)           |       -30% |
| Implementação realizada em desacordo com as especificações postas para o trabalho                           |       -50% |
| Programa apresenta erros de compilação, não executa ou apresenta saída incorreta                            |       -70% |
| Percentual de cobertura de testes inferior a 100%                                                           |       -40% |
| Erro na execução de testes unitários                                                                        |       -60% |
| Ausência ou incompletude de arquivos de teste                                                               |       -70% |
| Plagiarismo                                                                                                 |      -100% |

Caso a API seja documentada utilizando o Swagger, um percentual de 20% será somado à nota obtida até então
na avaliação. Caso a nota sem acréscimo tenha já alcançado a pontuação máxima prevista, ou seja, 10,0 pontos, a
pontuação excedente não será aproveitada.

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar _e-mail_ para o endereço;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.

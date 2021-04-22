# Mailing List API Web Go/MySQL/Templates/Bootstrap 
## Solução que utiliza usuários/email como exemplo para interação entre o cliente e o servidor através de templates nativos do Go, sendo altamente performáticos e intuitívos.

<p align="center">
  <img src="images/golang.png"/ alt="Golang">
</p>


A aplicação foi escrita totalmente em Go 🐹, visando utilizar o mínimo de dependências possíveis, tratar a maioria dos erros com os devidos cuidados e documentada com comentários de fácil entendimento

Pacotes utilizados

- Gorilla/Mux 🦍
- GORM
- Bootstrap


## Features

- Na aplicação é possível ver os usuários já adicionados previamente
- Adicionar novos usuários
- Ver usuários específicos
- Excluir algum usuário desejado
- Atualizar as informações


## Na prática


Iniciando a aplicação

![](images/1-starting.png)


Acessando os usuários já inseridos

![](images/2-adduser01.png)


Adicionando novos usuários

![](images/2-adduser02.png)

![](images/2-adduser03.png)

![](images/2-adduser04.png)

![](images/2-adduser05.png)

![](images/2-adduser06.png)

![](images/2-adduser07.png)

![](images/2-adduser08.png)


Atualizando o usuário desejado

![](images/3-updateuser01.png)

![](images/3-updateuser02.png)

![](images/3-updateuser03.png)

![](images/3-updateuser04.png)


Deletando o usuário desejado

![](images/4-deleteuser01.png)

![](images/4-deleteuser02.png)

![](images/4-deleteuser03.png)






## Requisitos

```sh
Golang: https://golang.org/dl/
MySQL instalado e configurado com as configurações desejadas (editar no arquivo credentials.go)
API Client de sua preferência (O que aparece nas imagens se chama Postman)

Após ter instalado o go
Gorilla mux: go get -u github.com/gorilla/mux
GORM: go get -u gorm.io/gorm
```


## Utilização

Basta utilizar o comando go build e aproveitar o aplicativo! 😊

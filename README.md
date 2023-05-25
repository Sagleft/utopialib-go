
![logo](logo.png)

Utopia Ecosystem API wrapper written in Golang

Docs
-----

[![GoDoc](https://godoc.org/github.com/sagleft/utopialib-go?status.svg)](https://godoc.org/gopkg.in/sagleft/utopialib-go.v1)
[![go-report](https://goreportcard.com/badge/github.com/Sagleft/utopialib-go)](https://goreportcard.com/report/github.com/Sagleft/utopialib-go)
[![Build Status](https://travis-ci.org/sagleft/utopialib-go.svg?branch=master)](https://travis-ci.org/sagleft/utopialib-go)

WARN! utopiago v1 deprecated, use v2 instead

Install
-----

```bash
go get github.com/Sagleft/utopialib-go/v2
```

then

```go
import (
	utopiago "github.com/Sagleft/utopialib-go/v2"
)
```

Usage
-----

```go
client := utopiago.NewClient("C17BF2E95821A6B545DC9A193CBB750B").
	SetProtocol("http").SetPort(22000).SetWsPort(25000)

myContactData, err := client.GetOwnContact()
if err != nil {
	log.Fatalln(err)
}

fmt.Println(myContactData.Pubkey)
```

or

```go
client := utopiago.UtopiaClient{
	Protocol: "http",
	Token:    "C17BF2E95821A6B545DC9A193CBB750B",
	Host:     "127.0.0.1",
	Port:     22791,
}

fmt.Println(client.CheckClientConnection())
```

How can this be used?
-----

* creating a web service that processes client requests;
* creation of a payment service;
* development of a bot for the channel;
* utility for working with uNS;
* experiments to explore web3.0;

package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Produtos struct {
	gorm.Model
	Id         uint `gorm:"primaryKey;autoIncrement"`
	Nome       string
	Descricao  string
	Preco      float32
	Quantidade int
}

func (p Produtos) String() string {
	return fmt.Sprintf("Id: %d, Nome: %s, Descricao: %s, Preco: %f, Quantidade: %d", p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
}

func generateRandomProduct() *Produtos {

	s1 := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(s1)

	randomColor := func() string {
		switch rnd.Intn(5) {
		case 1:
			return "Azul"
		case 2:
			return "Branca"
		case 3:
			return "Verde"
		case 4:
			return "Rosa"
		default:
			return "Amarela"
		}
	}()

	randomProduct := func() string {
		switch rnd.Intn(3) {
		case 1:
			return "Calça"
		case 2:
			return "Toalha"
		default:
			return "Camisa"
		}
	}()

	return &Produtos{
		Nome:       randomProduct,
		Descricao:  randomColor,
		Preco:      rnd.Float32() * 10,
		Quantidade: rnd.Intn(30),
	}
}

func randQuatidade() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(s1)
	return rnd.Intn(30)
}

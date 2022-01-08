package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGORM() {

	db := createCon()
	defer closeCon(db)

	db.AutoMigrate(&Produtos{})

	fmt.Println("")
	createOne(db)

	fmt.Println("")
	createBatch(db)

	fmt.Println("")
	findAll(db)

	fmt.Println("")
	findById(db)

	fmt.Println("")
	findByCondition(db)

	fmt.Println("")
	updateByCondition(db)

	fmt.Println("")
	updateSpecific(db)

	fmt.Println("")
	deleteSpecific(db)

	fmt.Println("")
	deleteAll(db)
}

func createCon() *gorm.DB {

	dbURI := "host=localhost user=db_user password=db_pass dbname=db_name port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db
}

func closeCon(db *gorm.DB) {
	if sql, err := db.DB(); err == nil {
		sql.Close()
	}
}

func createOne(db *gorm.DB) {
	produto := generateRandomProduct()
	produto.Quantidade = 29

	if result := db.Create(produto); result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Printf("createOne  -> %s \n", produto.String())
}

func createBatch(db *gorm.DB) {

	produtos := []*Produtos{}
	for i := 0; i < 5; i++ {
		produto := generateRandomProduct()
		produtos = append(produtos, produto)
	}

	if result := db.Create(produtos); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, elem := range produtos {
		fmt.Printf("createBatch  -> %s \n", elem.String())
	}
}

func findAll(db *gorm.DB) {
	var all []Produtos
	if result := db.Limit(10).Find(&all); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, elem := range all {
		fmt.Printf("FindAll  -> %s \n", elem.String())
	}
}

func findById(db *gorm.DB) *Produtos {
	var produto Produtos
	if result := db.Find(&produto, 1); result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	if produto.Id == 0 {
		fmt.Printf("FindById  -> not found \n")
		return nil
	}

	fmt.Printf("FindById  -> %s \n", produto.String())

	return &produto
}

func findByCondition(db *gorm.DB) {
	var all []Produtos

	// SELECT * FROM produtos WHERE nome LIKE 'Ca%';
	if result := db.Limit(10).Find(&all, "nome like ?", "Ca%"); result.Error != nil {
		//if result := db.Limit(10).Where("nome like ?", "Ca%").Find(&all); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, elem := range all {
		fmt.Printf("FindByCondition  -> %s \n", elem.String())
	}
}

func updateByCondition(db *gorm.DB) {

	// UpdateF ROM produtos SET quantidade=21 WHERE quantidade=21 ;
	if result := db.Model(&Produtos{}).Where("quantidade = ?", 29).Update("quantidade", randQuatidade()); result.Error != nil {
		fmt.Println(result.Error)
		return
	} else {
		fmt.Printf("UpdateByCondition  -> %d RowsAffected \n", result.RowsAffected)
	}
}

func updateSpecific(db *gorm.DB) {

	produto := findById(db)
	produtoRnd := generateRandomProduct()

	if result := db.Model(produto).Omit("id").Updates(produtoRnd); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	fmt.Printf("UpdateSpecific  -> %s \n", produto.String())
}

func deleteSpecific(db *gorm.DB) {

	produto := &Produtos{}
	db.Last(produto)
	fmt.Printf("DeleteSpecific  -> %s \n", produto.String())

	if result := db.Delete(produto); result.Error != nil {
		fmt.Println(result.Error)
		return
	} else {
		fmt.Printf("DeleteSpecific  -> %d RowsAffected \n", result.RowsAffected)
	}
}

func deleteAll(db *gorm.DB) {

	if result := db.Delete(&Produtos{}, "id > ?", 1); result.Error != nil {
		fmt.Println(result.Error)
		return
	} else {
		fmt.Printf("DeleteAll  -> %d RowsAffected \n", result.RowsAffected)
	}
}

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int     `gorm:"primaryKey"`
	Name  string  `gorm:"size:255"`
	Price float64 `gorm:"type:decimal(10,2)"`
}

func main() {
	// GORM: The fantastic ORM library for Golang, aims to be developer friendly.
	// https://gorm.io/

	// GORM 설치
	// go get -u gorm.io/gorm
	// go get -u gorm.io/driver/sqlite

	// GORM 문서
	// https://gorm.io/docs/

	// GORM GitHub
	// https://github.com/go-gorm/gorm

	dsn := "root:root@tcp(localhost:3306)/goexpert"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// SQLite
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// var product Product
	// db.First(&product, 1) // SELECT * FROM products WHERE id = 1;
	// println(product.ID, product.Name, product.Price)

	// db.First(&product, "name = ?", "Notebook") // SELECT * FROM products WHERE name = 'Notebook';
	// println(product.ID, product.Name, product.Price)

	// // select all
	// var products []Product
	// db.Find(&products) // SELECT * FROM products;
	// for _, p := range products {
	// 	println(p.ID, p.Name, p.Price)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products) // SELECT * FROM products; Limit 2 Offset 2
	// for _, p := range products {
	// 	println(p.ID, p.Name, p.Price)
	// }

	// Where
	// var p []Product
	// db.Where("price > ?", 3000).Find(&p)
	// for _, p := range p {
	// 	println(p.ID, p.Name, p.Price)
	// }

	// Like
	// var p2 []Product
	// db.Where("name LIKE ?", "%book%").Find(&p2)
	// for _, p := range p2 {
	// 	println(p.ID, p.Name, p.Price)
	// }

	var p Product
	db.First(&p, 1)
	p.Name = "Smartphone"
	p.Price = 2500.00
	db.Save(&p) // UPDATE products SET name='Smartphone', price=2500.00 WHERE id=1;

	println("After Update:")
	var updatedProduct Product
	db.First(&updatedProduct, 1)
	println(updatedProduct.ID, updatedProduct.Name, updatedProduct.Price)
	db.Delete(&updatedProduct) // DELETE FROM products WHERE id=1;

	println("After Delete:")
	var deletedProduct Product
	result := db.First(&deletedProduct, 1)
	if result.Error != nil {
		println("Product not found")
	} else {
		println(deletedProduct.ID, deletedProduct.Name, deletedProduct.Price)
	}

}

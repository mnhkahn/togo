// Package dmltogo
package dmltogo_test

import (
	"log"

	"github.com/mnhkahn/dmltogo"
)

func Example() {
	res, err := dmltogo.DmlToGo(`CREATE TABLE Persons (
                     PersonID int,
                     LastName varchar(255),
                     FirstName varchar(255),
                     Address varchar(255),
                     City varchar(255) 
                 );`)
	log.Println(res, err)
	/*  res is:
	package model

	type Person struct {
		Address   string `gorm:"column:Address" json:"Address"`
		City      string `gorm:"column:City" json:"City"`
		FirstName string `gorm:"column:FirstName" json:"FirstName"`
		LastName  string `gorm:"column:LastName" json:"LastName"`
		PersonID  int    `gorm:"column:PersonID" json:"PersonID"`
	}

	// TableName sets the insert table name for this struct type
	func (p *Person) TableName() string {
		return "Persons"
	}
	*/
}

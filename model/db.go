package model

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Book struct{
	Isbn    string
	Title   string
	Author  string
	Price   float32
}


var DB *mgo.Database
var books *mgo.Collection

func init(){

	s , err := mgo.Dial("mongodb://bond:moneypenny007@localhost:27017/bookstore")
	if err != nil{
		log.Fatalln(err)
	}

	err = s.Ping()
	if err != nil{
		log.Fatalln(err)
	}
	DB = s.DB("bookstore")
	books = DB.C("books")
	fmt.Println("database connected successfully")
}

func AllBooks()([]Book, error){
	bks := []Book{}
	err := books.Find(bson.M{}).All(&bks)
	return bks, err
}
func OneBooks(isbn string)Book{
	//rows  := DB.QueryRow("SELECT * FROM books where isbn = $1", isbn)
	var bk Book
	err := books.Find(bson.M{"isbn":isbn}).One(&bk)
	if err != nil{
		log.Fatalln(err)
	}
	return bk
}
func InsertBook(bk Book)error{
	err := books.Insert(bk)
	return  err
}
func UpdateBook(bk Book)error{
	err := books.Update(bson.M{"isbn":bk.Isbn},&bk)
	return  err
}
func DeleteBook(isbn string)error{
	err := books.Remove(bson.M{"isbn":isbn})
	return  err
}
package repository

import (
	"errors"
	"fmt"
	"log"
	"rest-go-demo/database/mysql"
	"rest-go-demo/entity"
)

//CreatePerson create function
func CreatePerson(person entity.Person) (interface{}, error) {
	err := mysql.Connector.Create(person)
	if err != nil {
		return nil, err.Error
	}
	res, _ := GetPerson(fmt.Sprint(person.ID))
	return res, nil
}

//GetPerson retrieves data from DB
func GetPerson(id string) (interface{}, error) {
	var person entity.Person
	rows, err := mysql.Connector.DB().Query("Select * from `people` where id = ? ", id)
	log.Println(rows.Next())
	if err != nil || !rows.Next() {
		log.Println("No Data found with id = ", id)
		err = errors.New("No data found with id = " + id)
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
		}
	}

	return person, err

}

//GetAllPersons retrieves all data
func GetAllPersons() ([]entity.Person, error) {
	var persons []entity.Person
	rows, err := mysql.Connector.DB().Query("SELECT * from `people`;")

	if err != nil || !rows.Next() {
		return nil, err
	}

	for rows.Next() {
		var person entity.Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
		persons = append(persons, person)
	}

	defer rows.Close()
	return persons, err
}

//DeletePerson deletes data for provided id
func DeletePerson(id string) error {

	res, err := mysql.Connector.DB().Exec("delete from `people` where id=?", id)
	if err != nil {
		return err
	}

	affectedRows, _ := res.RowsAffected()
	log.Printf("The statement affected %d rows\n", affectedRows)
	return nil
}

//UpdatePerson updates person
func UpdatePerson(person entity.Person) (interface{}, error) {
	rows, err := mysql.Connector.DB().Query("update `people` set first_name=?,last_name=?,age=? where id = ?;", person.FirstName, person.LastName, person.Age, person.ID)
	var personData interface{}
	if err != nil || !rows.Next() {
		if err == nil {
			err = errors.New("No such data")
		}
	} else {
		personData, _ = GetPerson(fmt.Sprint(person.ID))
	}
	rows.Close()
	return personData, err
}

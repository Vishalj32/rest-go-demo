package repository

import (
	"fmt"
	"log"
	"rest-go-demo/database/mysql"
	"rest-go-demo/entity"
)

//CreatePerson create function
func CreatePerson(person entity.Person) (entity.Person, error) {
	err := mysql.Connector.Create(person)
	if err != nil {
		return entity.Person{}, err.Error
	}
	res, _ := GetPerson(fmt.Sprint(person.ID))
	return res, nil
}

//GetPerson retrieves data from DB
func GetPerson(id string) (entity.Person, error) {
	var person entity.Person
	rows, err := mysql.Connector.DB().Query("Select * from `people` where id = ? ", id)
	if err != nil {
		return entity.Person{}, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
	}

	return person, nil
}

//GetAllPersons retrieves all data
func GetAllPersons() ([]entity.Person, error) {
	var persons []entity.Person
	rows, err := mysql.Connector.DB().Query("SELECT * from `people`;")

	if err != nil {
		return []entity.Person{}, err
	}

	for rows.Next() {
		var person entity.Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Age)
		persons = append(persons, person)
	}

	defer rows.Close()
	return persons, nil
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
func UpdatePerson(person entity.Person) (entity.Person, error) {
	_, err := mysql.Connector.DB().Query("update `people` set first_name=?,last_name=?,age=? where id = ?;", person.FirstName, person.LastName, person.Age, person.ID)

	if err != nil {
		return entity.Person{}, err
	}

	res, _ := GetPerson(fmt.Sprint(person.ID))
	return res, nil
}

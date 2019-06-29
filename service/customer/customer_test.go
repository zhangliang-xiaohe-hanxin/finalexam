package customer 

import (
	"testing"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	hostName = "postgres://dpsdgjur:qf4v1Qap7DKwpK3ZySXEWa7rB6B-VsJF@satao.db.elephantsql.com:5432/dpsdgjur"
)

var testingId int 


func TestTodo(t *testing.T) {
	t.Run("Insert", TestInsert)
	t.Run("Get", TestGetById)
	t.Run("Update", TestUpdate)
	t.Run("Delete", TestDelete)
}

func TestInsert(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	cus := Customer{ Name: "Lin Dan", Email: "Ball@gmail.com", Status: "Dead"} 
	err = insetIntoDB(&cus, session)
	if err != nil {
		t.Errorf("Cannot Insert Data    Errror: %v", err)
	}

	if cus.ID == 0 {
		t.Errorf("Got Zero Values in Id Field ")
	}

	testingId = cus.ID

}

func TestGetById(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	var cus Customer

	err = getByID(&cus, testingId, session)
	if err != nil {
		t.Errorf("Cannot get by Id  Errror: %v", err)
	}

	if cus.ID != testingId && cus.Name != "Lin Dan"  && cus.Email != "Ball@gmail.com" && cus.Status != "Dead" {
		t.Errorf("Data is not the same as Insert's Data")
	}
}

func TestUpdate(t *testing.T) {

	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	cus := Customer{ Name: "Lin Dan", Email: "a@gmail.com", Status: "Alive"} 
	err = updateByID(&cus, testingId, session)
	if err != nil {
		t.Errorf("Cannot Update  Errror: %v", err)
	}

	if cus.ID != testingId && cus.Name != "Lin Dan"  && cus.Email != "a@gmail.com" && cus.Status != "Alive" {
		t.Errorf("Data hasn't been updated yet")
	}
}

func TestDelete(t *testing.T) {
	session, err := sql.Open("postgres", hostName)
	if err != nil {
		t.Errorf("Cannot get Session   Errror: %v", err)
	}
	defer session.Close()

	err = delete(testingId, session)
	if err != nil {
		t.Errorf("Cannot Delete  Errror: %v", err)
	}

	var cus Customer

	err = getByID(&cus, testingId, session)
	if err == nil {
		t.Errorf("Data hasn't been deleted yet: %v", err)
	}

}
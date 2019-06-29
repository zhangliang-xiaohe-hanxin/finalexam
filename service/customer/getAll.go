package customer 

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/zhangliangxiaohehanxin/finalexam/database"
	"net/http"
	"log"
)

func (m Customer) GetStore(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	customers, err := getAll(m, session)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot query"})
		return
	}

	c.JSON(http.StatusOK, customers) 
}

func getAll(m Customer, session *sql.DB) ([]Customer, error) {

	var customers []Customer

	stmt, err := session.Prepare("SELECT id, name, email, status FROM customers")

	if err != nil {
		return []Customer{}, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return []Customer{}, err
	}

	for rows.Next() {
		err := rows.Scan(&m.ID, &m.Name, &m.Email, &m.Status)
		if err != nil {
			return []Customer{}, err
		}
		customers = append(customers, m)
	}

	return customers, nil
}
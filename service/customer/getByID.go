package customer 

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/zhangliangxiaohehanxin/finalexam/database"
	"strconv"
	"log"
)

func (m Customer) GetStoreByID(c *gin.Context) {

	session := db.Session

	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID Format"})
		return
	}
	
	err = getByID(&m, num, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Find by You ID"})
		return
	}

	c.JSON(http.StatusOK, m)
}

func getByID(m *Customer, id int, session *sql.DB) error {

	stmt, err := session.Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return err
	}
	row := stmt.QueryRow(id)
	err = row.Scan(&m.ID, &m.Name, &m.Email, &m.Status)
	if err != nil {
		return err
	}

	return nil
}

package customer 

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/finalexam/database"
	"strconv"
	"log"
	"net/http"
)

func (m *Customer) UpdateStoreByID(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")
	num, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID Format"})
		return
	}
	
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Cannot recieve Data"})
		return
	}

	err = updateByID(m, num, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Update"})
		return
	}

	c.JSON(http.StatusOK, m)
}

func updateByID(m *Customer, id int, session *sql.DB) error {

	stmt, err := session.Prepare("UPDATE customers SET name=$3, status=$2 WHERE id=$1 RETURNING ID;")
	if err != nil {
		return err
	}

	row := stmt.QueryRow(id, m.Status, m.Name)
	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}
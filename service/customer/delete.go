package customer

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/zhangliangxiaohehanxin/finalexam/database"
)

func (m Customer) DeleteStoreByID(c *gin.Context) {

	session, err := db.GetSession(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Cannot Get Session DB"})
		return
	}

	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed"})
		return
	}

	err = delete(num, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed"})
		return
	}

	c.JSON(200, gin.H{"message": "customer deleted"})
}

func delete(id int, session *sql.DB) error {

	stmt, err := session.Prepare("DELETE from customers WHERE id=$1;")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

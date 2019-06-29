package customer

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zhangliangxiaohehanxin/finalexam/database"
)

func (m *Customer) Insert(c *gin.Context) {

	session := db.Session

	if err := c.ShouldBindJSON(&m); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Something wrong with you request Body"})
		return
	}

	err := insetIntoDB(m, session)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot Insert"})
		return
	}

	c.JSON(http.StatusCreated, m)

}

func insetIntoDB(m *Customer, session *sql.DB) error {
	stmt, err := session.Prepare("INSERT INTO customers(name, email, status) VALUES($1, $2, $3) RETURNING ID;")
	if err != nil {
		return err
	}

	row := stmt.QueryRow(m.Name, m.Email, m.Status)
	err = row.Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

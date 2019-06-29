package routes 

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/finalexam/database"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type APIMethod interface {
	Insert(c *gin.Context) 
	GetStore(c *gin.Context)
	GetStoreByID(c *gin.Context)
	UpdateStoreByID(c *gin.Context)
	DeleteStoreByID(c *gin.Context)
}

type Route struct {
	API		APIMethod
	DBHost string
}

func (r Route) createSession() {

	session, err := sql.Open("postgres", r.DBHost)
	if err != nil {
		log.Fatal("can't open", err.Error())
	}
	db.Session = session
}

func (r Route) DestroySession() {
	db.Session.Close()
}

func authMiddleware(c *gin.Context){
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": "Invalid Authorization"})
		return
	}
	c.Next()
}

func (r Route)Init() *gin.Engine {
	db.CreateDB(r.DBHost)
	r.createSession()
	apiMethod := r.API
	routes := gin.Default()
	routes.Use(authMiddleware)
	routes.GET("/customers", apiMethod.GetStore)
	routes.GET("/customers/:id", apiMethod.GetStoreByID)
	routes.POST("/customers", apiMethod.Insert)
	routes.PUT("/customers/:id", apiMethod.UpdateStoreByID)
	routes.DELETE("/customers/:id", apiMethod.DeleteStoreByID)

	return routes
}
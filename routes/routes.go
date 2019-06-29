package routes 

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangliangxiaohehanxin/finalexam/database"
	"net/http/httputil"
	"bytes"
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


type buffer struct {
	gin.ResponseWriter
	response *bytes.Buffer
}

func (b buffer) Write(data []byte) (int, error) {
	b.response.Write(data)
	return b.ResponseWriter.Write(data)
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

func monitorAPI(c *gin.Context) {
	req, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{ "message": "Cannot Log Your Request"})
		return
	}

	log.Printf("\n\nrequest: %s\n", req)
	buff := bytes.Buffer{}
	bw := buffer{c.Writer, &buff}
	c.Writer = bw
	c.Next()
	log.Printf("\n\nresponse: %s\n\n", bw.response.String())

}

func (r Route)Init() *gin.Engine {
	db.CreateDB(r.DBHost)
	r.createSession()
	apiMethod := r.API
	routes := gin.Default()
	routes.Use(monitorAPI)
	routes.Use(authMiddleware)
	routes.GET("/customers", apiMethod.GetStore)
	routes.GET("/customers/:id", apiMethod.GetStoreByID)
	routes.POST("/customers", apiMethod.Insert)
	routes.PUT("/customers/:id", apiMethod.UpdateStoreByID)
	routes.DELETE("/customers/:id", apiMethod.DeleteStoreByID)

	return routes
}

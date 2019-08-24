// 2.2 Please setup a REST service for CRUD
package main
import (
	"github.com/gin-gonic/gin"
	"log"
)
type apartmentRequest struct{
	Name string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}
func main(){
	r := gin.Default()
	r.GET("/apartment", getApartments)
	r.POST("/apartment", addApartment)
	r.PUT("/apartment/:apartmentId", editApartment)
	r.DELETE("/apartment/:apartmentId", deleteApartment)
	r.Run(":8000")
}

func getApartments(c *gin.Context){
	apartments := getApartment()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	c.JSON(200, gin.H{"apartments":apartments})
	// c.String(200, "GET METHOD")

}

func addApartment(c *gin.Context){
	var req apartmentRequest
	if err := c.ShouldBind(&req); err != nil{
		log.Fatal(err)
	}
	id := saveApartment(req.Name, req.Address)
	if id != 0 {
		c.JSON(200, gin.H{"message":"apartment has been added"})
	}
	// c.String(200, req.Name)
	// c.String(200, "ADD METHOD")
}

func editApartment(c *gin.Context){
	c.JSON(200, gin.H{"message": c.Param("apartmentId")})
}

func deleteApartment(c *gin.Context){
	c.JSON(200, gin.H{"message": c.Param("apartmentId")})
}
package routes
import (
	agent "gin/controllers/agent"
	user "gin/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users/:userId", user.GetUser)
		api.GET("/users", user.GetAllUser)
		api.GET("/users/edit/:userId", user.UpdateUser)
		//api.DELETE("/users/:id", user.DeleteUser)
	}
	apiAgent := router.Group("/agent")
	{
		apiAgent.PUT("/refreshQrcode", agent.RefreshQrcode)
		//api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	_ = router.Run(":8000")
}
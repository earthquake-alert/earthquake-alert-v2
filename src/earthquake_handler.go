package src

import "github.com/gin-gonic/gin"

func EarthquakeHandler(c gin.Context) error {
	// ?force を指定すると強制的にJMAの電文からGETします
	// _, isForce := c.GetQuery("force")

	// if isForce {
	// 	...
	// }

	return nil
}

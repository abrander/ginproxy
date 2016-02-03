# ginproxy
A very simple proxy handler for gin-gonic

Can be used like this:

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/abrander/ginproxy"
)

func main() {
	router := gin.Default()
	g, _ := ginproxy.NewGinProxy("http://backend01.example.com/")
	router.Any("/api/*all", g.Handler)
	router.Run(":8080")
}
```

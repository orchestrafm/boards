package routing

import (
	"github.com/spidernest-go/mux"
)

var r *echo.Echo

const ErrGeneric = `{"errno": "404", "message": "Bad Request"}`

func ListenAndServe() {
	r = echo.New()

	v0 := r.Group("/api/v0")
	v0.POST("/board", createBoard)
	v0.GET("/board/:id", getBoard)
	v0.PUT("/board/:id", updateBoard)
	v0.PATCH("/board/:id", editBoard)
	v0.GET("/board/track/:id", getBoardsFromTrack)
	v0.DELETE("/board/:tid/:bid", deleteBoard)

	r.Start(":5000")
}

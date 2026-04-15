package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Lables struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type ReqestLables struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var lables = []Lables{}

func GetLables(c echo.Context) error {
	return c.JSON(http.StatusOK, lables)
}

func PostLables(c echo.Context) error {
	var req ReqestLables

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid req lables"})
	}

	result := req.Text

	if result == "" {
		return c.NoContent(http.StatusNoContent)
	}

	lable := Lables{
		Id:        uuid.NewString(),
		Text:      result,
		Completed: false,
	}

	lables = append(lables, lable)

	return c.JSON(http.StatusCreated, lable)

}

func completedTask(c echo.Context) error {
	id := c.Param("id")

	var req ReqestLables

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid reqest lables PATCH"})
	}

	status := !req.Completed

	for i, lable := range lables {
		if lable.Id == id {
			lables[i].Completed = status
			return c.JSON(http.StatusOK, lables[i])
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.RequestLogger())

	e.GET("/TODO", GetLables)
	e.POST("/TODO", PostLables)
	e.PATCH("/TODO/completedTask/:id", completedTask)
	if err := e.Start("localhost:8080"); err != nil {
		fmt.Println(err)
	}
}

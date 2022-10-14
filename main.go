package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	shapeManagers "shape/managers"
	"strconv"
	userManagers "user/managers"
)

func main() {
	r := gin.New()
	r.GET("/rectangle/area/:length/:width", func(context *gin.Context) {
		shapeManager := shapeManagers.GetShapeManager()

		length, err := strconv.ParseUint(context.Param("length"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		width, err := strconv.ParseUint(context.Param("width"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Rectangle Area",
			"body": gin.H{
				"area": shapeManager.GetAreaRectangle(length, width),
			},
		})
		return
	})
	r.GET("/rectangle/perimeter/:length/:width", func(context *gin.Context) {
		shapeManager := shapeManagers.GetShapeManager()

		length, err := strconv.ParseUint(context.Param("length"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		width, err := strconv.ParseUint(context.Param("width"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "Rectangle Area",
			"body": gin.H{
				"perimeter": shapeManager.GetPerimeterRectangle(length, width),
			},
		})
		return
	})
	r.GET("/users", func(c *gin.Context) {
		userManager := userManagers.GetUserManager()
		users, err := userManager.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		var userMap []map[string]any = nil
		for _, user := range users {
			userMap = append(userMap, user.GetMap())
		}
		if userMap == nil {
			userMap = []map[string]any{}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Users",
			"body": gin.H{
				"user": users,
			},
		})
		return
	})
	r.POST("/users", func(context *gin.Context) {
		userManager := userManagers.GetUserManager()
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		data := map[string]string{}

		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := userManager.Create(data["first_name"], data["last_name"])
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "User Created",
			"body": gin.H{
				"id":         user.Id,
				"first_name": user.FirstName,
				"last_name":  user.LastName,
			},
		})
		return
	})
	r.PATCH("/users/:id", func(context *gin.Context) {
		userManager := userManagers.GetUserManager()

		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		data := map[string]string{}

		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = userManager.Update(id, data["first_name"], data["last_name"])
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "User Updated",
			"body": gin.H{
				"id":         id,
				"first_name": data["first_name"],
				"last_name":  data["last_name"],
			},
		})
		return
	})
	r.DELETE("/users/:id", func(context *gin.Context) {
		userManager := userManagers.GetUserManager()
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = userManager.Delete(id)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "User deleted",
		})
		return
	})
	err := r.Run()
	if err != nil {
		return
	}
}

package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Sunchiii/champamker-service/configs"
	"github.com/Sunchiii/champamker-service/helpers"
	"github.com/Sunchiii/champamker-service/models"
	"github.com/Sunchiii/champamker-service/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newMember models.Member
		defer cancel()

		//check data user input
		if err := c.ShouldBindJSON(&newMember); err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "data invalid",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//check authentication
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := helpers.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, responses.Status{
				Status:  http.StatusUnauthorized,
				Massage: "UNAUTHORIZRED",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		member := models.Member{
			Id:           uuid.NewString(),
			FirstNameLao: newMember.FirstNameLao,
			LastNameLao:  newMember.LastNameLao,
			FirstNameENG: newMember.FirstNameENG,
			LastNameENG:  newMember.LastNameENG,
			Class:        newMember.Class,
			Tell:         newMember.Tell,
			Email:        newMember.Email,
			Role:         newMember.Role,
			Goal:         newMember.Goal,
			MagicWord:    newMember.MagicWord,
			Image:        newMember.Image,
		}

		if err := configs.MemberDB.WithContext(cxt).Create(&member).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "CAN'T_CREATE_A_NEWDATA",
				Data:    map[string]interface{}{"data": "can't create a new data"},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data:    map[string]interface{}{"data": member},
		})
	}
}

func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var member []models.Member
		defer cancel()

		//query date from database
		if err := configs.MemberDB.WithContext(ctx).Find(&member).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "SOMETHING_WRONG_ON_SERVER_SIDE",
				Data:    map[string]interface{}{"data": "something wrong on server side please try again"},
			})
			return
		}

		//response to user
		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data: map[string]interface{}{
				"data": member,
			},
		})

	}
}

func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var member models.Member
		id := c.Param("id")
		defer cancel()

		if err := configs.MemberDB.WithContext(ctx).Where("id = ?", id).First(&member).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "NOT_FOUND",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//response to user
		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data: map[string]interface{}{
				"data": member,
			},
		})
	}
}

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newData models.Member
		id := c.Param("id")
		defer cancel()

		//check user input
		if err := c.ShouldBindJSON(&newData); err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "data invalid",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//check authentication
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := helpers.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, responses.Status{
				Status:  http.StatusUnauthorized,
				Massage: "UNAUTHORIZRED",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//find data on data base
		//and define old data
		var currentData models.Member
		if err := configs.MemberDB.WithContext(ctx).Where("id = ?", id).First(&currentData).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "NOT_FOUND",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//update date to date base
		if err := configs.MemberDB.WithContext(ctx).Model(&currentData).Updates(newData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "CAN'T_NOT_UPDATE",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//response to user
		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data: map[string]interface{}{
				"data": currentData,
			},
		})
	}
}

func DeleteMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		defer cancel()

		//check authentication
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := helpers.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, responses.Status{
				Status:  http.StatusUnauthorized,
				Massage: "UNAUTHORIZRED",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//check member on data base befor the next process
		var currentMember models.Member
		if err := configs.MemberDB.WithContext(ctx).Where("id = ?", id).First(&currentMember).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "NOT_FOUND",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//delete member by id
		if err := configs.MemberDB.WithContext(ctx).Where("id = ?", id).Delete(&currentMember).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "CAN'T_NOT_DELETE",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		//response to user
		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data: map[string]interface{}{
				"data": currentMember,
			},
		})

	}
}

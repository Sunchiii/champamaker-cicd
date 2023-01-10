package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Sunchiii/champamker-service/configs"
	"github.com/Sunchiii/champamker-service/helpers"
	"github.com/Sunchiii/champamker-service/models"
	"github.com/Sunchiii/champamker-service/responses"
	"github.com/gin-gonic/gin"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newAdmin models.Admin
		defer cancel()

		//check data when user binding
		if err := c.ShouldBindJSON(&newAdmin); err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "data invalid",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//check user on database
		if err := configs.AdminDB.WithContext(ctx).Where(&newAdmin).First(&newAdmin).Error; err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "IS_NOT_ALREADY_EXIT",
				Data:    map[string]interface{}{"data": "please register to be admin"},
			})
			return
		}

		//generate token
		token, err := helpers.GenerateToken()
		if token == "" || err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "SOMETHING_WRONG_ON_SERVER_SIDE",
				Data:    map[string]interface{}{"data": "something wrong on server side"},
			})
			return
		}
		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data:    map[string]interface{}{"token": token},
		})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newAdmin models.Admin
		defer cancel()

		//check user request valid date
		if err := c.BindJSON(&newAdmin); err != nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "data invalid",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//
		//check user on database
		var check models.Admin
		if err := configs.AdminDB.WithContext(ctx).Where("username = ?", newAdmin.Username).First(&check).Error; err == nil {
			c.JSON(http.StatusBadRequest, responses.Status{
				Status:  http.StatusBadRequest,
				Massage: "ALREADY_EXIT",
				Data:    map[string]interface{}{"data": "already exit user"},
			})
			return
		}

		if err := configs.AdminDB.WithContext(ctx).Create(&newAdmin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "CAN'T_CREATE_A_NEWDATA",
				Data:    map[string]interface{}{"data": "can't create a new data"},
			})
			return
		}

		//generate token
		token, err := helpers.GenerateToken()
		if token == "" || err != nil {
			c.JSON(http.StatusInternalServerError, responses.Status{
				Status:  http.StatusInternalServerError,
				Massage: "SOMETHING_WRONG_ON_SERVER_SIDE",
				Data:    map[string]interface{}{"data": "something wrong on server side"},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Status{
			Status:  http.StatusOK,
			Massage: "SUCCESSFULLY",
			Data: map[string]interface{}{"data": gin.H{
				"username": newAdmin.Username,
				"token":    token,
			}},
		})

	}
}

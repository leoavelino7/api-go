package controllers

import (
	"api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (value *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, value.tweets)
}

func (value *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	value.tweets = append(value.tweets, *tweet)

	ctx.JSON(http.StatusCreated, tweet)
}	

func (value *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, tweet := range value.tweets {
		if tweet.ID == id {
			value.tweets = append(value.tweets[:i], value.tweets[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Tweet deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}

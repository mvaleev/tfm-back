package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var t = Thumb{
	ThumbURL: "some url",
	ThumbID:  "some id",
}

var v = Video{
	VideoURL:   "some url",
	Popularity: "some pop",
}

func todayThumbs(c *gin.Context) {
	ts := []Thumb{}
	ts = append(ts, t)

	c.JSON(http.StatusOK, ts)
}

func somedayThumbs(c *gin.Context) {
	id := c.Param("day")
	t.ThumbID = id
	ts := []Thumb{}
	ts = append(ts, t)

	c.JSON(http.StatusOK, ts)
}

func getThumb(c *gin.Context) {
	id := c.Param("id")
	t.ThumbID = id
	ts := []Thumb{}
	ts = append(ts, t)

	c.JSON(http.StatusOK, ts)
}

func getVideo(c *gin.Context) {
	id := c.Param("id")
	v.VideoURL = id
	vs := []Video{}
	vs = append(vs, v)

	c.JSON(http.StatusOK, vs)
}

package utils

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ErrorJson(c *gin.Context, v interface{}, code ...int) {
    status := http.StatusInternalServerError
    if len(code) > 0 {
        status = code[0]
    }
    c.JSON(status, v)
}

func JsonBytes(c *gin.Context, b []byte) {
    c.Data(http.StatusOK, "application/json", b)
}

package gemini

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) GetResponse(c *gin.Context) {
	var gemReq GeminiResponseReq
	if err := c.ShouldBindJSON(&gemReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.Service.GetResponse(c.Request.Context(), &gemReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetImageResponse(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	buf := make([]byte, header.Size)
	_, err = file.Read(buf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var gemReq GeminiImageQueryReq
	gemReq.Query = c.Request.PostFormValue("query")
	gemReq.File = buf
	gemReq.FileType = header.Header.Get("Content-Type")

	res, err := h.Service.GetImageResponse(c.Request.Context(), &gemReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

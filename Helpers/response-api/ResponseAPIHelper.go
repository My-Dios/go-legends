package ResponseAPIHelper

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStandarization(statusCode int, message string, c *gin.Context, data interface{}, count int, limit int, currentPage int, totalPages int) {
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	fullURL := protocol + "://" + c.Request.Host + c.Request.URL.String()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	nextPageUrl := ""
	prevPageUrl := ""
	if currentPage < totalPages {
		nextPageUrl = buildURLWithPage(fullURL, currentPage+1, totalPages)
	}
	if currentPage > 1 {
		prevPageUrl = buildURLWithPage(fullURL, currentPage-1, totalPages)
	}
	responseData := struct {
		Path      string `json:"path"`
		Timestamp string `json:"timestamp"`
		Result    struct {
			Message     string      `json:"message"`
			Count       int         `json:"count"`
			TotalPages  int         `json:"totalPages"`
			CurrentPage int         `json:"currentPage"`
			Data        interface{} `json:"data,omitempty"`
			NextPageUrl string      `json:"next"`
			PrevPageUrl string      `json:"prev"`
		} `json:"result"`
	}{
		Path:      fullURL,
		Timestamp: timestamp,
		Result: struct {
			Message     string      `json:"message"`
			Count       int         `json:"count"`
			TotalPages  int         `json:"totalPages"`
			CurrentPage int         `json:"currentPage"`
			Data        interface{} `json:"data,omitempty"`
			NextPageUrl string      `json:"next"`
			PrevPageUrl string      `json:"prev"`
		}{
			Message:     message,
			Count:       count,
			TotalPages:  totalPages,
			CurrentPage: currentPage,
			Data:        data,
			NextPageUrl: nextPageUrl,
			PrevPageUrl: prevPageUrl,
		},
	}
	c.JSON(statusCode, responseData)
}

func buildURLWithPage(url string, page, totalPages int) string {
	if page <= 0 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}
	pageParam := "page=" + strconv.Itoa(page)
	if strings.Contains(url, "?") {
		if strings.Contains(url, "page=") {
			re := regexp.MustCompile(`page=\d+`)
			url = re.ReplaceAllString(url, pageParam)
		} else {
			url += "&" + pageParam
		}
	} else {
		url += "?" + pageParam
	}
	return url
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

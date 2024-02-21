package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"

    "iTools/page"
    "iTools/script"
)

func ShareGet(c *gin.Context) {
    pageCount := 10
    currentPage := c.Query("page")

    files := script.ListFiles(UploadPath)
    p := page.GetPageTable(currentPage, len(files), pageCount)
    files = files[p.StartIdx:p.EndIdx]

    c.HTML(http.StatusOK, "share.html", gin.H{
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "files":         files,
        "fileCount":     len(files),
        "peviousPage":   p.Previous,
        "currentPage":   p.Current,
        "nextPage":      p.Next,
        "currentTables": p.CurrentTables,
        "lastPage":      p.Last,
    })
}

func ShareDownload(c *gin.Context) {
    base64FileName := c.Param("base64FileName")

    fileName := script.DeBase64(base64FileName)

    filePath := filepath.Join(UploadPath, fileName)

    // if err != nil || !script.FileExist(filePath) {
    //   CantFound(c)
    // }

    c.FileAttachment(filePath, fileName)
}

var _ = fmt.Println

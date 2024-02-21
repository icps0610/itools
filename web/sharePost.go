package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"

    "iTools/script"
)

func SharePost(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["uploads"]

    for _, file := range files {
        fileName := script.FileNameWithRandNum(file.Filename)

        path := filepath.Join(UploadPath, fileName)
        c.SaveUploadedFile(file, path)
    }

    c.Redirect(http.StatusMovedPermanently, "/share")
}

func ShareClean(c *gin.Context) {
    var cmd string
    if IsWin {
        cmd = fmt.Sprintf(`del /q %s\*`, UploadPath)
    } else {
        cmd = fmt.Sprintf(`rm %s/*`, UploadPath)
    }
    script.RunCmd(cmd)

    c.Redirect(http.StatusMovedPermanently, "/share")
}

package web

import (
    "fmt"
    "html/template"
    "os"
    "path/filepath"
    "runtime"
    "strings"

    "github.com/gin-contrib/gzip"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/gin-gonic/gin"

    "iTools/conf"
    "iTools/script"
)

var (
    RootDirPath = getRootDirPath()
    TempDirPath = getTempDirPath()

    UploadPath = TempDirPath + `share`

    staticDirPath = filepath.Join(RootDirPath, `static`)
    htmlPath      = filepath.Join(RootDirPath, `templates`, `*.html`)

    IsWin = (runtime.GOOS == "windows")
)

func init() {
    os.Mkdir(UploadPath, 777)
}

func Service() *gin.Engine {
    if conf.Mode != `debug` {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()
    router.Use(gzip.Gzip(gzip.DefaultCompression))
    // router.Use(gin.Logger())

    router.Static("/static", staticDirPath)
    router.Static("/tmp", TempDirPath)
    router.Static("/img", UploadPath)

    router.SetFuncMap(template.FuncMap{
        "mode": modeFunc,
    })

    router.LoadHTMLGlob(htmlPath)

    store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("itools", store))

    return router
}

func getRootDirPath() string {
    path, err := os.Executable()
    printError(err)
    if string(path[:11]) == `z:\go-build` {
        path, err = os.Getwd()
        printError(err)
        return path + `\`
    }
    path = filepath.Dir(path)
    if runtime.GOOS == "windows" {
        return path + `\`
    }
    return path + `/`
}

func getTempDirPath() string {
    if runtime.GOOS == "windows" {
        return `z:\`
    }
    return `/tmp/`
}

func setMeg(c *gin.Context, typeName, msg string) {
    var bgColor string
    if typeName == "alert" {
        bgColor = `#ff9999`
    } else if typeName == "success" {
        bgColor = `#88dd88`
    }
    setSession(c, "bgColor", bgColor)
    setSession(c, "msg", msg)
}

func getMsg(c *gin.Context) string {
    var flash string
    msg := getSession(c, "msg")
    if msg != "" {
        flash = msg
        setSession(c, "msg", "")
    }
    return flash
}

func getBgColor(c *gin.Context) string {
    var flash string
    msg := getSession(c, "bgColor")
    if msg != "" {
        flash = msg
        setSession(c, "bgColor", "")
    }
    return flash
}

func setSession(c *gin.Context, sessionName, sessionValue string) {
    session := sessions.Default(c)
    session.Set(sessionName, sessionValue)
    session.Save()
}

func getSession(c *gin.Context, sessionName string) string {
    session := sessions.Default(c)
    msg := session.Get(sessionName)

    return To_interface_s(msg)
}

func To_interface_s(i interface{}) string {
    var text string
    is, ok := i.(string)
    if ok {
        text = is
    }
    return text
}

func query(c *gin.Context, query string) string {
    text := c.Query(query)
    return strings.Replace(text, " ", "+", -1)
}

func queryAndDeBase64(c *gin.Context, queryStr string) string {
    text := query(c, queryStr)
    return script.DeBase64(text)
}

func modeFunc() string {
    return conf.Mode
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

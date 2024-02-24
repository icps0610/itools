package main

import (
    "iTools/conf"
    . "iTools/web"
)

func main() {
    router := Service()

    router.GET("/", IndexGet)

    router.GET("/url", UrlGet)
    router.POST("/url", UrlPost)

    router.GET("/rem", RemGet)
    router.POST("/rem", RemPost)
    router.POST("/remID", RemIDPost)

    router.GET("/indent", IndentGet)
    router.POST("/indent", IndentPost)

    router.GET("/carry", CarryGet)
    router.POST("/carry", CarryPost)

    router.GET("/passwd", PasswdGet)
    router.POST("/passwd", PasswdPost)

    router.GET("/base64", Base64Get)
    router.POST("/enBase64", EnBase64Post)
    router.POST("/deBase64", DeBase64Post)

    router.GET("/urlCode", UrlCodeGet)
    router.POST("/enUrl", EnUrlPost)
    router.POST("/deUrl", DeUrlPost)

    if conf.Mode == `pi` || conf.Mode == `debug` {
        router.GET("/share", ShareGet)
        router.GET("/share/:base64FileName", ShareDownload)

        router.POST("/share", SharePost)
        router.POST("/shareClean", ShareClean)
    }

    if conf.Mode == `pi` {
        router.Run(":8010")
    } else {
        router.Run(":3000")
    }

}

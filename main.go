package main

import (
    "github.com/gin-gonic/gin"
    "github.com/pquerna/otp/totp"
    "html/template"
    "net/http"
    "time"
    "strings"
)

func main() {
    r := gin.Default()

    // 加载模板
    r.SetFuncMap(template.FuncMap{})
    r.LoadHTMLGlob("templates/*")

    // 首页路由
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "OTP":       "",
            "Countdown": 0,
            "SecretKey": "",
        })
    })

    // 处理表单提交
    r.POST("/", func(c *gin.Context) {
        secretKey := c.PostForm("secret_key")
        secretKey = strings.ReplaceAll(secretKey, " ", "") // 去除空格
        otp, err := totp.GenerateCode(secretKey, time.Now())
        if err != nil {
            c.HTML(http.StatusOK, "index.html", gin.H{
                "OTP":       "无效的密钥: " + err.Error(),
                "Countdown": 0,
                "SecretKey": secretKey,
            })
            return
        }

        // 计算剩余时间
        now := time.Now().Unix()
        remainingTime := 30 - (now % 30) // 计算距离下一个 30 秒时间点的剩余时间

        c.HTML(http.StatusOK, "index.html", gin.H{
            "OTP":       otp,
            "Countdown": remainingTime,
            "SecretKey": secretKey,
        })
    })

    // 启动服务器
    r.Run(":5656")
}
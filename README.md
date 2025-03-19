# 2FA Generator

🚀 一个轻量级的 2FA（双因素认证）验证码生成器，支持二维码生成，方便集成到你的应用中。

## 🔥 功能特点
- 生成基于 TOTP（时间一次性密码）的 2FA 验证码
- 生成二维码，方便扫描添加到身份验证器（如 Google Authenticator）
- docker 轻量化安装部署

---

## 🛠️ 安装 & 运行

### **1. 使用 Docker 运行**
你可以直接使用 Docker 运行该应用：
```sh
docker run -d -p 5656:5656 --name 2fa-generator terrysiu/2fa-generator:latest
```

然后在浏览器中访问：
```sh
http://localhost:5656
```

---

## 📜 许可证

本项目遵循 MIT 许可证，欢迎贡献！🎉

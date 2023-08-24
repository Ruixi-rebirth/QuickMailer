## 功能

发送者(仅支持qq邮箱)可发送附件(可多个)和正文给接受者(可多个)

## 配置解释

```toml
filePaths = ["path/to/file1.txt", "path/to/file2.txt"] #待发送的附件路径
senderEmail = "your_email@qq.com" #发送者qq邮箱
senderPassword = "your_password" #在 [账号与安全](https://wx.mail.qq.com/account) --安全设置--SMTP/IMAP服务 中开启服务并获取授权码
subject = "Your Subject" #邮件主题
body = """
Hello,

This is the body of the email.

Regards,
Your Name
""" #邮件正文,不需要可设置为 body = ""

[[receivers]]
name = "Receiver1Name" #接受者姓名(作用仅仅只是标识)
email = "receiver1_email@example.com" #接受者邮箱

[[receivers]]
name = "Receiver2Name"
email = "receiver2_email@example.com"

```

## 如何使用?

1. 下载自己电脑cpu架构对应的二进制文件 [release下载](https://github.com/Ruixi-rebirth/QuickMailer/releases)

2. 将这个`可执行文件、config.toml、待发送附件`放在同一级目录下面,执行这个可执行文件,即可瞬间发送。

注意 可将项目中 `example-config.toml` 的内容复制到你的配置文件`config.toml`，改一改就可以了

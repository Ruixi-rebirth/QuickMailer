## 功能

发送者(仅支持qq邮箱)发送附件(可多个)给接受者(可多个)

## 配置解释

```
{
"filePath": "./xxx.txt", // 将被发送的附件路径
"senderEmail": "xxx@qq.com", // 发送者qq邮箱
"senderPassword": "xxx", // 在 [账号与安全](https://wx.mail.qq.com/account) --安全设置--SMTP/IMAP服务 中开启服务并获取授权码替换 xxx
"receiverEmail": "xxx@qq.com", // 接受者qq邮箱
"subject": "xxx" //邮件主题
}
```

## 如何使用?

下载自己电脑cpu架构对应的二进制文件,然后将这个`可执行文件、config.json、待发送附件`放在同一级目录下面,执行这个可执行文件,即可瞬间发送。
注意 可将项目中 `example-config.json` 的内容复制到你的配置文件，改一改就可以了

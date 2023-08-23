## 功能

发送者(仅支持qq邮箱)发送附件(可多个)给接受者(可多个)

## 配置解释

```
{
  "filePaths": [
    "/path/to/file1", // 待发送的附件路径
    "/path/to/file2"
  ],
  "senderEmail": "xxx@qq.com", // 发送者qq邮箱
  "senderPassword": "xxx", // 在 [账号与安全](https://wx.mail.qq.com/account) --安全设置--SMTP/IMAP服务 中开启服务并获取授权码替换 xxx
  "receivers": [
    {
      "name": "xxx", // 接受者姓名(作用仅仅只是标识)
      "email": "xxx@qq.com"  // 接受者qq邮箱
    },
    {
      "name": "xxx",
      "email": "xxx@qq.com"
    }
  ],
  "subject": "xxx"  //邮件主题
}
```

## 如何使用?

1. 下载自己电脑cpu架构对应的二进制文件 [release下载](https://github.com/Ruixi-rebirth/QuickMailer/releases)

2. 将这个`可执行文件、config.json、待发送附件`放在同一级目录下面,执行这个可执行文件,即可瞬间发送。

注意 可将项目中 `example-config.json` 的内容复制到你的配置文件，改一改就可以了

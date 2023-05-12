# <center><strong>BTG</strong></center>
 
<br>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JNyaa/BT2Telegram?label=Go%20Version&style=flat-square)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/JNyaa/BT2Telegram?label=Release%20Version&style=flat-square)](https://github.com/JNyaa/BT2Telegram/release)
[![Apache-2.0](https://img.shields.io/github/license/JNyaa/BT2Telegram?style=flat-square)](https://github.com/JNyaa/BT2Telegram/blob/master/LICENSE)

[![GitHub issues](https://img.shields.io/github/issues/JNyaa/BT2Telegram?label=Sticker%20Issues&style=flat-square)](https://github.com/JNyaa/BT2Telegram/issues)
[![GitHub Repo stars](https://img.shields.io/github/stars/JNyaa/BT2Telegram?label=Stars&style=flat-square)](https://github.com/JNyaa/BT2Telegram/stargazers)
![GitHub release (latest by date)](https://img.shields.io/github/downloads/JNyaa/BT2Telegram/latest/total?label=Downloads%40Latest&style=flat-square)
![GitHub repo size](https://img.shields.io/github/repo-size/JNyaa/BT2Telegram?style=flat-square)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/JNyaa/BT2Telegram?style=flat-square)

<br>

---

<p align="center"><strong>宝塔面板的Telegram消息推送通道<strong></p>
<p align="center"><strong>可能以后会推出更多功能<strong></p>


<br>
<p align="center"><a href="https://github.com/JNyaa/BT2Telegram/releases">Download</a> | <a href="https://github.com/JNyaa/BT2Telegram/issues">Issues</a> | <a href="#notice">Notice</a> | <a href="#donate">Donate</a> </p>


<br>

# 需求

1. 宝塔需求: 消息通道支持邮件,官方版/开心版均可,不限大陆/国际版
2. 系统需求: Linux建议Kernel 4+, Windows建议内部版本最低18362

<br>

# 配置文件
配置文件为 `btg.yml`
```
bot:
  key: BotAPI的Key，从@BotFather处获取
  chat: ChatID, 如果不知道ChatID，如果想在群组里发送消息，先把机器人拉进群，配置好key后启动，在群内输入指令 `/my`，将ChatID的值复制到配置文件的`bot.chat`中，重新启动机器人即可。chat支持的类型: 群组、频道、私聊(需要先私聊机器人以创建会话)
  api_server: 自定义Telegram Bot API地址 (https://example.com/)
  send: 消息发送模式
smtp:
  port: 邮件服务器端口，请务必在防火墙封禁端口
  auth:
    status: 邮件服务器基础验证，默认为false(关)，设置为true以启用验证
    user: 邮件服务器验证用户名
    pass: 邮件服务器验证密码
```

# 开始搭建

## 安装

1. [下载](https://github.com/JNyaa/BT2Telegram/releases) 二进制文件，将其移动到任意文件夹并给予root 777权限
2. 从Github仓库中找到 btg.yml，将其下载并和二进制放在一起 (给予读写权限)
3. 配置btg.yml。
4. [配置systemd](#systemd配置) (必选)
5. [配置宝塔](#宝塔配置) (必选)

<br>

## systemd配置

1. 从Github仓库中找到 btg.service，将其下载并和二进制放在一起 (给予读写权限)
2. 修改service文件中的目录
3. `mv btg.service /usr/lib/systemd/system/btg.service`
4. `systemectl start btg`
5. 如果需要开机启动，再额外输入命令 `systemctl enable btg`


## 宝塔配置
1. 点击 `面板设置-告警通知-告警设置` (如果有问题，点击 `通知设置-设置推送`，点击窗口左下角的`更新列表`即可，配置完后还需要在此再点击一遍`更新列表`)
2. 找到`邮箱-点击设置`
3. 发送人邮箱和SMTP密码随便填 (全填1也行，如果在btg.yml中开启验证，则须与配置文件中一致)
4. SMTP服务器设置为127.0.0.1，端口设置为配置文件中的smtp.port端口
5. 点击保存即可

后面的收件人邮箱也是随便填，第一次会自动初始化。


# Notice

1. 不接受无脑Issues，态度恶劣的直接删除
2. 机器人功能简单，无需增强，不接受合并请求


<br>

# Donate

USDT/ETH: `0xF1C61348613489C5007a1A7aFbfd72bfdC4a3C3A`

TON: `UQCKEwhmCtTBmFNwESxeEmb8Ij4hSAU1wFs_inPy3tBbLExf`


<br>

# License
This software is distributed under Apache-2.0 license.
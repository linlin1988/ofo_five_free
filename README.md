# ofo_five_free


# 编译源码, 使用go

```
    go build ofo.go

```

# 开始薅羊毛

```

./ofo -k xxxxxxx(invitekey) -t 13510000002

```


# 运行效果


```
[root@test code]# ./ofo -k xxxxxxx(invitekey) -t 13530564560
[13530564561]邀请的号码成功
[13530564562]邀请的号码成功
[13530564563]邀请的号码成功
[13530564564]邀请的号码成功
[13530564565]邀请的号码成功


#这个程序是for循环++1 发遍全国
```

# 打开APP看收获

<img src='https://raw.githubusercontent.com/linlin1988/ofo_five_free/master/ofo.jpg'> </br>



# 如何获得自己的invitekey

1. 电脑网页版打开  https://ofo-misc.ofo.com/pay/index.html#/ 进行登录 </br>
2. 登录后访问 https://ofo-misc.ofo.com/invite/index.html#/</br>
3. 打开浏览器开发者模式</br>
4. f5刷新页面</br>
5. 找到请求列表里的 inviteConfig 这一条，查看响应结果中的inviteKey</br>

如
```
{"errorCode":200,"msg":"操作成功","values":{"inviterName":"","inviteKey":"xxxxxxxx","isInviteGroup":0,"invitedNum":113252,"invitedPacketNum":56}}

```

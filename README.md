# 海康威视iSecure后渗透

## 简介

*主要用于获取权限后通过数据库信息进行资产信息确认*

## 使用

------
**参数**

``` shell
  -P int
        database port.            # 数据库端口
  -c    change database password. # 密码修改
  -f string
        config file path.         # 配置文件地址
  -h string
        database host name.       # 数据库主机地址
  -p string
        database password.        # 数据库密码(可直接粘贴密文)
  -u string
        database user name.       # 数据库用户
```

**使用方法一**

上传对应可执行文件到目标服务器后使用-f参数指定配置文件使用。(默认配置文件 **/opt/hikvision/web/components/postgresql11linux64.1/conf/config.properties**)

![](img\1.png)

**使用方法二**

手动设定账户密码主机等信息，因为部分海康系统的postgres端口(7092)会对外网开放，所以可以在获取密码密文后进行参数指定运行。

![](img\2.png)
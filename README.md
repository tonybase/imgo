# IM-GO

## 概要
这是一个使用go语言开发的高性能IM服务端。单台服务器可以承受近百万并发(经测试单台服务器可开启90W+协程,30W+TCP长连接)，实现具体IM基本功能，用户能够达到一般IM所拥有的使用功能。
## 服务端设计
![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/login.png)
## 服务端特点
	具有稳定性，可拓展性，高效性，高并发等特点
## 客户端运行截图
登录界面

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/client_login.png)

主界面

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/client_main.png)

聊天界面

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/client_chat.png)

## 客户端设计
	客户端是新版本的IQQ项目。使用Swing开发。可以在win,osx,linux上良好运行。具有良好的界面交互效果
## 测试数据
	服务器配置：
	操作系统：CentOS 6.5
	CPU： 2核心
	内存： 4G
	系统盘： 12G
	SSD：40G
	带宽：5M

测试结果：

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/start_server.png)

空闲时服务器信息：

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/server_idle.png)

建立5000TCP连接之后的服务器信息：

![image](https://raw.githubusercontent.com/im-qq/imgo/master/docs/images/server_5k.png)

## 未来
#### 开源版本：
	1. 发送表情
	2. 发送文件
	3. 好友分组
	4. 离线重连机制处理
	5. 系统重构提高性能
#### 商业版本：
	1. 集成微信
	2. 开发App
	3. 开发webchat
	4. 添加管理后台
	5. 系统重构为分布式设计
	6. 添加缓存机制
	7. 简单的数据采集和BI

## 版权声明
	开发者:Tony、Itnik

	该系统遵守GPL3.0开源协议。
	商业使用请联系作者。

## Link

+ [IM协议文档](https://github.com/im-qq/imgo/blob/master/docs/IM%E6%8E%A5%E5%8F%A3%E5%8D%8F%E8%AE%AEV1.0.1.docx)
+ [客户端](https://github.com/im-qq/italk)

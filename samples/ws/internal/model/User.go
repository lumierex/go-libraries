package model

// 启动扫描
// {type: xxx, xxxx}   知道username

// web和sever之间建立一个 channel
// 每个人都有自己的专属channel

// map[username:scanTypeSource:timestamp]
// Manager管理所有的 ws-管理

// front 扫描 => 创建ws => 根据对应的 handleRequestType
// Type Mysql
// HandleMysql

// redis lpush brpop
// 获取当前的扫描机器
//
// 当前机器向客户端发起ws连接
// 每次扫描调用sendResult

type User struct {
}

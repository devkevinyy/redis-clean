### Redis 数据清理工具

#### 概述
    该工具的目标是无阻塞的对 Redis 的数据进行清理，采用 Scan 的方式可针对 Big Key 进行无阻塞的清理。

#### 使用方式
    * 在 release 目录获取对应平台的二进制可执行文件，目前已编译有 MacOS、Linux、Windows 平台。
    * 运行文件，按照说明传入对应参数：
    >   --host            string         redis实例连接地址
    >   --auth            string         用户名和密码
    >   --db              int            指定DB
    >   --pattern         string         Key匹配模式
    >   --count           int            每次scan的数量

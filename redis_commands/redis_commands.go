package redis_commands

const (
	DEL               = "DEL"               // 用于删除 key
	DUMP              = "DUMP"              // 序列化给定 key ，并返回被序列化的值
	EXISTS            = "EXISTS"            // 检查给定 key 是否存在
	EXPIRE            = "EXPIRE"            // 为给定 key 设置过期时间
	EXPIREAT          = "EXPIREAT"          // 用于为 key 设置过期时间，接受的时间参数是 UNIX 时间戳
	PEXPIRE           = "PEXPIRE"           // 设置 key 的过期时间，以毫秒计
	PEXPIREAT         = "PEXPIREAT"         // 设置 key 过期时间的时间戳(unix timestamp)，以毫秒计
	KEYS              = "KEYS"              // 查找所有符合给定模式的 key
	MOVE              = "MOVE"              // 将当前数据库的 key 移动到给定的数据库中
	PERSIST           = "PERSIST"           // 移除 key 的过期时间，key 将持久保持
	PTTL              = "PTTL"              // 以毫秒为单位返回 key 的剩余的过期时间
	TTL               = "TTL"               // 以秒为单位，返回给定 key 的剩余生存时间(
	RANDOMKEY         = "RANDOMKEY"         // 从当前数据库中随机返回一个 key
	RENAME            = "RENAME"            // 修改 key 的名称
	RENAMENX          = "RENAMENX"          // 仅当 newkey 不存在时，将 key 改名为 newkey
	TYPE              = "TYPE"              // 返回 key 所储存的值的类型
	SET               = "SET"               // 设置指定 key 的值
	GET               = "GET"               // 获取指定 key 的值
	GETRANGE          = "GETRANGE"          // 返回 key 中字符串值的子字符
	GETSET            = "GETSET"            // 将给定 key 的值设为 value ，并返回 key 的旧值 ( old value )
	GETBIT            = "GETBIT"            // 对 key 所储存的字符串值，获取指定偏移量上的位 ( bit )
	MGET              = "MGET"              // 获取所有(一个或多个)给定 key 的值
	SETBIT            = "SETBIT"            // 对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)
	SETEX             = "SETEX"             // 设置 key 的值为 value 同时将过期时间设为 seconds
	SETNX             = "SETNX"             // 只有在 key 不存在时设置 key 的值
	SETRANGE          = "SETRANGE"          // 从偏移量 offset 开始用 value 覆写给定 key 所储存的字符串值
	STRLEN            = "STRLEN"            // 返回 key 所储存的字符串值的长度
	MSET              = "MSET"              // 同时设置一个或多个 key-value 对
	MSETNX            = "MSETNX"            // 同时设置一个或多个 key-value 对
	PSETEX            = "PSETEX"            // 以毫秒为单位设置 key 的生存时间
	INCR              = "INCR"              // 将 key 中储存的数字值增一
	INCRBY            = "INCRBY"            // 将 key 所储存的值加上给定的增量值 ( increment )
	INCRBYFLOAT       = "INCRBYFLOAT"       // 将 key 所储存的值加上给定的浮点增量值 ( increment )
	DECR              = "DECR"              // 将 key 中储存的数字值减一
	DECRBY            = "DECRBY"            // 将 key 所储存的值减去给定的减量值 ( decrement )
	APPEND            = "APPEND"            // 将 value 追加到 key 原来的值的末尾
	HDEL              = "HDEL"              // 用于删除哈希表中一个或多个字段
	HEXISTS           = "HEXISTS"           // 用于判断哈希表中字段是否存在
	HGET              = "HGET"              // 获取存储在哈希表中指定字段的值
	HGETALL           = "HGETALL"           // 获取在哈希表中指定 key 的所有字段和值
	HINCRBY           = "HINCRBY"           // 为存储在 key 中的哈希表指定字段做整数增量运算
	HKEYS             = "HKEYS"             // 获取存储在 key 中的哈希表的所有字段
	HLEN              = "HLEN"              // 获取存储在 key 中的哈希表的字段数量
	HSET              = "HSET"              // 用于设置存储在 key 中的哈希表字段的值
	HVALS             = "HVALS"             // 用于获取哈希表中的所有值
	BLPOP             = "BLPOP"             // 移出并获取列表的第一个元素
	BRPOP             = "BRPOP"             // 移出并获取列表的最后一个元素
	BRPOPLPUSH        = "BRPOPLPUSH"        // 从列表中弹出一个值，并将该值插入到另外一个列表中并返回它
	LINDEX            = "LINDEX"            // 通过索引获取列表中的元素
	LINSERT           = "LINSERT"           // 在列表的元素前或者后插入元素
	LLEN              = "LLEN"              // 获取列表长度
	LPOP              = "LPOP"              // 移出并获取列表的第一个元素
	LPUSH             = "LPUSH"             // 将一个或多个值插入到列表头部
	LPUSHX            = "LPUSHX"            // 将一个值插入到已存在的列表头部
	LRANGE            = "LRANGE"            // 获取列表指定范围内的元素
	LREM              = "LREM"              // 移除列表元素
	LSET              = "LSET"              // 通过索引设置列表元素的值
	LTRIM             = "LTRIM"             // 对一个列表进行修剪(trim)
	RPOP              = "RPOP"              // 移除并获取列表最后一个元素
	RPOPLPUSH         = "RPOPLPUSH"         // 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	RPUSH             = "RPUSH"             // 在列表中添加一个或多个值
	RPUSHX            = "RPUSHX"            // 为已存在的列表添加值
	SADD              = "SADD"              // 向集合添加一个或多个成员
	SCARD             = "SCARD"             // 获取集合的成员数
	SDIFF             = "SDIFF"             // 返回给定所有集合的差集
	SDIFFSTORE        = "SDIFFSTORE"        // 返回给定所有集合的差集并存储在 destination 中
	SINTER            = "SINTER"            // 返回给定所有集合的交集
	SINTERSTORE       = "SINTERSTORE"       // 返回给定所有集合的交集并存储在 destination 中
	SISMEMBER         = "SISMEMBER"         // 判断 member 元素是否是集合 key 的成员
	SMEMBERS          = "SMEMBERS"          // 返回集合中的所有成员
	SMOVE             = "SMOVE"             // 将 member 元素从 source 集合移动到 destination 集合
	SPOP              = "SPOP"              // 移除并返回集合中的一个随机元素
	SRANDMEMBER       = "SRANDMEMBER"       // 返回集合中一个或多个随机数
	SREM              = "SREM"              // 移除集合中一个或多个成员
	SUNION            = "SUNION"            // 返回所有给定集合的并集
	SUNIONSTORE       = "SUNIONSTORE"       // 所有给定集合的并集存储在 destination 集合中
	SSCAN             = "SSCAN"             // 迭代集合中的元素
	ZADD              = "ZADD"              // 向有序集合添加一个或多个成员，或者更新已存在成员的分数
	ZCARD             = "ZCARD"             // 获取有序集合的成员数
	ZCOUNT            = "ZCOUNT"            // 计算在有序集合中指定区间分数的成员数
	ZINCRBY           = "ZINCRBY"           // 有序集合中对指定成员的分数加上增量 increment
	ZINTERSTORE       = "ZINTERSTORE"       // 计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 key 中
	ZLEXCOUNT         = "ZLEXCOUNT"         // 在有序集合中计算指定字典区间内成员数量
	ZRANGE            = "ZRANGE"            // 通过索引区间返回有序集合成指定区间内的成员
	ZRANGEBYLEX       = "ZRANGEBYLEX"       // 通过字典区间返回有序集合的成员
	ZRANGEBYSCORE     = "ZRANGEBYSCORE"     // 通过分数返回有序集合指定区间内的成员
	ZRANK             = "ZRANK"             // 返回有序集合中指定成员的索引
	ZREM              = "ZREM"              // 移除有序集合中的一个或多个成员
	ZREMRANGEBYLEX    = "ZREMRANGEBYLEX"    // 移除有序集合中给定的字典区间的所有成员
	ZREMRANGEBYRANK   = "ZREMRANGEBYRANK"   // 移除有序集合中给定的排名区间的所有成员
	ZREMRANGEBYSCORE  = "ZREMRANGEBYSCORE"  // 移除有序集合中给定的分数区间的所有成员
	ZREVRANGE         = "ZREVRANGE"         // 返回有序集中指定区间内的成员，通过索引，分数从高到底
	ZREVRANGEBYSCORE  = "ZREVRANGEBYSCORE"  // 返回有序集中指定分数区间内的成员，分数从高到低排序
	ZREVRANK          = "ZREVRANK"          // 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	ZSCORE            = "ZSCORE"            // 返回有序集中，成员的分数值
	ZUNIONSTORE       = "ZUNIONSTORE"       // 计算一个或多个有序集的并集，并存储在新的 key 中
	ZSCAN             = "ZSCAN"             // 迭代有序集合中的元素（包括元素成员和元素分值）
	PSUBSCRIBE        = "PSUBSCRIBE"        // 订阅一个或多个符合给定模式的频道。
	PUBSUB            = "PUBSUB"            // 查看订阅与发布系统状态。
	PUBLISH           = "PUBLISH"           // 将信息发送到指定的频道。
	PUNSUBSCRIBE      = "PUNSUBSCRIBE"      // 退订所有给定模式的频道。
	SUBSCRIBE         = "SUBSCRIBE"         // 订阅给定的一个或多个频道的信息。
	UNSUBSCRIBE       = "UNSUBSCRIBE"       // 指退订给定的频道。
	DISCARD           = "DISCARD"           // 取消事务，放弃执行事务块内的所有命令
	EXEC              = "EXEC"              // 执行所有事务块内的命令
	MULTI             = "MULTI"             // 标记一个事务块的开始
	UNWATCH           = "UNWATCH"           // 取消 WATCH 命令对所有 key 的监视
	WATCH             = "WATCH"             // 监视一个(或多个) key
	GEOHASH           = "GEOHASH"           // 返回一个或多个位置元素的 Geohash 表示
	GEOPOS            = "GEOPOS"            // 从key里返回所有给定位置元素的位置（经度和纬度）
	GEODIST           = "GEODIST"           // 返回两个给定位置之间的距离
	GEORADIUS         = "GEORADIUS"         // 以给定的经纬度为中心， 找出某一半径内的元素
	GEOADD            = "GEOADD"            // 将指定的地理空间位置（纬度、经度、名称）添加到指定的key中
	GEORADIUSBYMEMBER = "GEORADIUSBYMEMBER" // 找出位于指定范围内的元素，中心点是由给定的位置元素决定
	HINCRBYFLOAT      = "HINCRBYFLOAT"      // 为哈希表 key 中的指定字段的浮点数值加上增量 increment
	HMGET             = "HMGET"             // 获取所有给定字段的值
	HMSET             = "HMSET"             // 同时将多个 field-value (域-值)对设置到哈希表 key 中
	HSETNX            = "HSETNX"            // 只有在字段 field 不存在时，设置哈希表字段的值
	HSCAN             = "HSCAN"             // 迭代哈希表中的键值对
	HSTRLEN           = "HSTRLEN"           // 返回哈希表 key 中， 与给定域 field 相关联的值的字符串长度
	PING              = "PING"              // ping
	WITHSCORES        = "WITHSCORES"        // 参数，分数
)

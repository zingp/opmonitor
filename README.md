### opmonitor
    check proccess status, if proc is dead, make it runing. Count items from your custom logs.

### 存在问题

    不能统计多个日志中的字段，修改时应该注意map不是线程安全的

    之前认为自己实现的corn比较占用资源，改成第三方cron库实现；实际上占用资源的原因是日志过大，过虐太消耗资源
    统计的日志量越小，性能消耗越低
### 数据库的连接

mysql 命令是忽略大小写的

MySQL 的数据类型

1. 整数型
    - tinyint 最大长度4
    - smallint 最大长度6
    - mediumint 最大长度8
    - int 最大长度 11, 第一位表示符号 + 或 -
    - bigint 最大长度 20

注意: int(M) M表示最大显示宽度, 在 int(M) 中, M 的值跟 int(M) 所占多少存储空间并无任何关系
和数字位数也无关系 int(3) int(4) int(8) 在磁盘上都是占用 4 bytes的存储空间

int(11), tinyint(1) bigint(20), 后面的数字, 不代表占用空间容量, 而代表最小显示位数.

设计 mySql 数据库时, 建表时, mysql 会自动分配长度

2. 浮点型
    - float
    - double
    - decimal如果不指定精度, 默认为 (10, 0)

浮点数相对是整数优点在长度一定的情下, 浮点数表示更大的范围, 缺点时引起精度问题

3. 字符串型

| 值         | char(4) | 存储需求    | varchar(4) | 存储需求    |
|-----------|---------|---------|------------|---------|
| ""        | ""      | 4 bytes | ""         | 1 bytes |
| "ab"      | "ab "   | 4 bytes | "ab"       | 3 bytes |
| "abcd"    | "abcd"  | 4 bytes | "abcd"     | 5 bytes |
| "abcdefg" | "abcd"  | 4 bytes | "abcd"     | 5 bytes |

varchar 使用额外的 1-2 字节内来存储值长度, 列长度 <= 255 使用 1 字节保存

4. 备注型 (详情等使用)

| 类          | 描述               |
|------------|------------------|
| tinytext   | 字符串, 最大 255个字符   |
| text       | 字符串, 最大 65535个字符 |
| mediumtext | 字符串, 16777215    |
| longtext   | 4294967295       |

5. 日期型

datetime "0000-00-00 00:00:00"

```mysql
mysql -u root -p

Enter password: *****
```

显示当前存在的数据库

```mysql
show databases;
```

创建一张数据库

```mysql
create database books;
```

选择你所需要操作的数据库

```mysql
use dataBase(数据库名);
# Database changed
```

数据库中创建一张表

```mysql
create table users
(
    id       int(11),
    username varchar(255),
    age      int(3),
    sex      int(1)
);

# 班级表
create table class
(
    id    int(11) NOT NULL AUTO_INCREMENT,
    name  varchar(255),
    email varchar(255),
    score tinyint(4),
    PRIMARY KEY (`id`)
);
```

查看当前数据库的表

```mysql
show tables;

# users
```

查看表中的所有数据

```mysql
select * from users;
```

指定字段查找数据

```mysql
select id, username from users;
```

查找指定数据

```mysql
selcect id from users where id=1;
selcect * from users where id>1;
selcect id from users where id<1 and age=15;
select * from users where username = "yym";

# is null 为空
# is not null 不为空
# between 在范围内
select * from class where email is null;

# between 查找成绩大于等于 70 小于等于 90 的数据
select * from class where score >= 70 and score <= 90;
select * from class where score between 70 and 90;
select * from class where score not between 70 and 90;

# 查找 scroe 不是 89 和 98 的数据
select * from class where score not in (89, 98);

# or and
select * from class where score=85 or score=60;
select * from class where score>=85 and score<=99;

# like 模糊查询
select *  from class where email like "%qq%"
select *  from class where email like "ja%" # ja 开头的
select *  from class where email like "%163.com" # 163 结尾的
select *  from class where email not like "%163.com" # 不是 163 结尾的
```

显示表的结构

```mysql
describe users;
```

给表中新增一条数据

```mysql
insert into users(id, username, age, sex)
values (1, "wxf", 23, 1);
insert into users(username) values ("zzh");

insert into `class` values (1, "张三", "1@qq.com", 55);
```

修改表中指定的数据

```mysql
update users set id = 10 where username = "yym";
# 修改多个参数
update users set age=17, sex=1 where id = 2;
```

删除表中指定的数据

```mysql
delete from users where id = 1;
```

按指定的数据排序

```mysql
# 按照 status 倒序排序
select * from users order by status desc;
# desc 降序 asc 升序
select * from users order by id desc;
select * from users order by id desc, sex asc;
```

统计数量

```mysql
# 1 代表第一列
select count(1) from users;
select count(*) from users;
```

Limit

```mysql
select id, name from users limit 2;
# 2, 3 跳过两条, 查询三条
select id, name from users limit 2, 3; 
```

删除指定的表

```mysql
drop table users;
```

删除数据库

```mysql
drop database book;
```

MySql 分组函数

```mysql
# AVG(column) 求平均值
select avg(score) from class;

# count(column) 统计行数

# max(column) 最大值
select * from class where score in(select max(score) from class);

# min(column) 最小值

# sum(column 求和
```


Mysql 别名 `as`

```mysql
select name as a from class;
select min(score) as minscore from class;
```

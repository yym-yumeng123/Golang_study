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


### MySQL 表与表之间的关系

- 一对一 例如: 一个文章只能属于一个分类
- 一对多 例如: 文章分类 对文章, 一个分类可以有多个文章
- 多对多 例如: 一个学生可以选择多个课程, 一个课程可以被多个学生选择

MySQL `笛卡尔积连接, 内连接, 左外连接, 右外连接`

查询数据的时候能不用连接语句尽量不用, 笛卡尔积连接查询较慢, 项目中用的多的是内连接

```mysql
# 1. 查找文章显示文章分类

# 笛卡尔积连接
select article.id, article.title, article.state, article_cate as cate 
from article, article_cate where acticle.cate_id=acticle_cate.id

# 内连接 inner join ... on
# select 后面要查询的内容
select article.id, article.title, article.state, article_cate as cate
from article inner join article_cate on acticle.cate_id=acticle_cate.id
```

```mysql
# 多对多, 可以建一张 中间表 
# 学生和课程 lesson_id student_id
# A 同学选修了 那些课程

# 查询A同学选修的课程 id
select lesson_id from lesson_student where student_id = 1;
# 查询出课程id 对应的课程 简单查询 in
select * from lesson where id in (select lesson_id from lesson_student where student_id = 1);

# 内连接查询
select lesson.id, lesson.name from lesson inner join lesson_student
on lesson.id=lesson_student.lesson_id and lesson_student.student_id = 1;


# 课程被那些同学选修了, 课程为1的学生
select student_id from lesson_student where lesson_id = 1;
select * from student where id in (select student_id from lesson_student where lesson_id = 1);
```

```mysql
# 左外连接 left join on => lesson 表所有信息输出, lesson_student 表不满足条件都是空
select * from lesson left join lesson_student
on lesson.id=lesson_student.lesson_id and lesson_student.student_id = 1;

# 右外连接 right join ... on
select * from lesson right join lesson_student
on lesson.id=lesson_student.lesson_id and lesson_student.student_id = 1;
```


### 索引

MySQL 索引的建立对于 MySQL的高效运行是很重要的, 索引可以大大提高检索速度

如果没有索引, 执行查询的时候必须从第一条记录开始, 扫描整个表的记录, 直到符合要求的记录.如果
有了索引, mysql 无需扫描任何记录即可顺序找到目标记录的位置. 简单来说, 索引就是提高查找速度, 数据量越多, 效果越明显

MySQL中常见的索引

- 普通索引
- 唯一索引 => 字段名称不能重复
- 全文索引
- 空间索引 Spatial

```mysql
# 设置普通索引

# class(name) 表的字段 index_name 索引名称
create index index_name on my_table(name);

alter table table_name add index index_name(name)
alter table table_name drop index index_name

# 查看索引 \G 格式化输出
show index from class\G; 

# 删除索引
drop index index_name on class;

# 创建唯一索引  (主键是一种唯一索引)
create unique index index_name on table(name)
```

### 事务

事务处理用来维护数据库的完整性, 保证成批的 SQL 语句要么全部执行, 要么全不执行.

```mysql
# begin 开始一个事务
# rollback 事务回滚
# commit 事务确认

# 转账操作
begin;
update users set balance = balance - 100 where username="小王";
update users set balance = balance + 100 where username="小李";
commit;
```

### 锁

Mysql 中的锁有: 表级锁和行级锁

**表级锁**

1. 添加读锁

可以并发读,但是不能并发写, 读锁期间, 没释放锁之前不能进行写操作

```mysql
# 确保无人对这个记录进行 update 和 delete 操作
lock table user read; # 给表设置一个读锁
     
unlock tables;
```

2. 添加写锁

只有锁表的用户可以进行读写操作, 其它用户不行 (并发下对商品库存的操作)

```mysql
lock table user write;
unlock tables;
```
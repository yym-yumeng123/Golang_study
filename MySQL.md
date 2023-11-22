### 数据库的连接

mysql 命令是忽略大小写的

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
create table users(
    id int(11),
    username varchar(255),
    age int(3),
    sex int(1)
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
select * from users where username="yym";
```

显示表的结构

```mysql
describe users;
```

给表中新增一条数据

```mysql
insert into users(id, username, age, sex) values (1, "wxf", 23, 1);
insert into users(username) values ("zzh");
```

修改表中指定的数据

```mysql
update users set id = 10 where username="yym";
# 修改多个参数
update users set age=17, sex=1 where id=2;
```

删除表中指定的数据

```mysql
delete from users where id=1;
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
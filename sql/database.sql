CREATE DATABASE IF NOT EXISTS user_grpc;
USE user_grpc;


DROP IF EXISTS users;
CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  email varchar(50) not null unique,
  document varchar(20) not null,
  age int(3) not null,
  createdAt timestamp default current_timestamp()
)ENGINE=INNODB;
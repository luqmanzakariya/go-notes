CREATE DATABASE belajar_golang_restful_api

use belajar_golang_restful_api

CREATE TABLE category(
	id integer primary key auto_increment,
	name varchar(200) not null
) engine = InnoDb;

SELECT * FROM category c 
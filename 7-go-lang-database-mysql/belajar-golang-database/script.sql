-- 145-membuka-koneksi-ke-database_test
-- 146-database-pooling_test
-- 147-eksekusi-perintah-sql_test
CREATE TABLE customer (
	id VARCHAR(100) NOT NULL,
	name VARCHAR(100) NOT NULL,
	PRIMARY KEY (id)
) ENGINE = InnoDb;

SHOW tables;
describe customer;

-- 148-query-sql_test
SELECT * FROM customer c ;

-- 149-tipe-data-column_test
DELETE FROM customer;

ALTER TABLE customer 
	ADD COLUMN email VARCHAR(100),
	ADD COLUMN balance INT DEFAULT 0,
	ADD COLUMN rating DOUBLE DEFAULT 0.0,
	ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	ADD COLUMN birth_date DATE,
	ADD COLUMN married BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer (id, name, email, balance, rating, birth_date, married)
VALUES 	('eko', 'Eko', 'eko@gmail.com', 1000000, 90.0, '1999-10-10', true),
 		('budi', 'Budi', 'budi@gmail.com', 500000, 85.5, '1986-06-10', true),
		('joko', 'Joko', 'joko@gmail.com', 750000, 88.5, '1988-06-10', false);
		
UPDATE customer 
SET email = NULL, birth_date = NULL 
WHERE id = 'joko';

-- 150-sql-injection_test.go
CREATE TABLE user(
	username VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	PRIMARY KEY (username)
) ENGINE = InnoDB;

SELECT * FROM user;

INSERT INTO user(username, password) VALUES ('admin', 'admin');

-- 152-auto-increment_test.go
CREATE TABLE comments(
	id INT NOT NULL AUTO_INCREMENT,
	email VARCHAR(100) NOT NULL,
	comment TEXT,
	PRIMARY KEY (id)
) ENGINE = InnoDB;

DESC comments;

SELECT * FROM comments;
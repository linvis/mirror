DROP TABLE iF EXISTS user;

CREATE TABLE user (
	user_id INT(11) PRIMARY KEY AUTO_INCREMENT,
	email CHAR(32) NOT NULL,
	user_name CHAR(32) NOT NULL UNIQUE,
	passwd CHAR(32) NOT NULL,
	male TINYINT,
	age TINYINT,
	location CHAR(32),
	role CHAR(32),
    avatar CHAR(255)
);

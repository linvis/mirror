DROP TABLE IF EXISTS activity;

CREATE TABLE activity
(
  event_id int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  event_type int(11) DEFAULT NULL,
  start_time timestamp NULL DEFAULT NULL,
  end_time timestamp NULL DEFAULT NULL,
	duration CHAR(32) DEFAULT NULL,
	num int(11),
	comment VARCHAR(255)
);
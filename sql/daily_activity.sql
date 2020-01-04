DROP TABLE IF EXISTS daily_evt;
CREATE TABLE daily_evt
(
  evt_id int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(11) NOT NULL,
	evt_type int(11) NOT NULL,
  evt_item int(11) NOT NULL,
	evt_date BIGINT not NULL,
  start_time int(11) NULL DEFAULT NULL,
  end_time int(11) NULL DEFAULT NULL,
	duration int(11) DEFAULT NULL,
	num int(11),
	comment VARCHAR(255)
);

DROP TABLE IF EXISTS daily_evt_type;
CREATE TABLE daily_evt_type
(
  evt_type int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  evt_name char(64) NOT NULL,
  user_id int(11) NOT NULL
);

INSERT INTO daily_evt_type(evt_name, user_id)
value ('睡眠', 0), ('工作', 0), ('阅读',0), ('编程',0), ('运动',0), ('娱乐',3);

DROP TABLE IF EXISTS daily_evt_item;
CREATE TABLE daily_evt_item
(
  evt_item int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
	evt_item_name char(60) NOT NULL,
  evt_type int(11) NOT NULL,
  user_id int(11) NOT NULL
);
INSERT INTO daily_evt_item(evt_item_name, evt_type, user_id)
VALUE ('刷题', 4, 3), ('跑步', 5, 3);
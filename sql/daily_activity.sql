DROP TABLE IF EXISTS daily_act;

CREATE TABLE daily_act
(
  act_id int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id int(11) NOT NULL,
  act_type int(11) DEFAULT NULL,
  start_time bigint NULL DEFAULT NULL,
  end_time bigint NULL DEFAULT NULL,
	duration CHAR(32) DEFAULT NULL,
	num int(11),
	comment VARCHAR(255)
);

DROP TABLE IF EXISTS daily_act_type;
CREATE TABLE daily_act_type
(
  act_id int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  act_name char(64) NOT NULL,
  user_id int(11) DEFAULT NULL
);


INSERT INTO daily_act_type(act_name, user_id)
value ('睡眠', 0), ('工作', 0), ('阅读',0), ('刷题'，0), ('健身',0 ), ('娱乐',0);


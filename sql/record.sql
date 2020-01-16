DROP TABLE IF EXISTS sleep_record;
CREATE TABLE sleep_record
(
  record_id INT(11)
  UNSIGNED ZEROFILL NOT NULL PRIMARY KEY AUTO_INCREMENT,
  user_id INT
  (11) NOT NULL,
	evt_date INT NOT NULL,
  start_time MEDIUMINT DEFAULT NULL,
  end_time MEDIUMINT DEFAULT NULL,
	duration MEDIUMINT NOT NULL
);

# demo data
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1577721600, -13, 527, 540);
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1577808000, -11, 467, 478);
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1577894400, 40, 504, 464);
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1577980800, 93, 542, 449);
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1578067200, -2, 529, 531);
  INSERT INTO sleep_record
    (user_id, evt_date, start_time, end_time, duration)
  VALUES
    (3, 1578153600, 27, 529, 502);

-- DROP TABLE IF EXISTS record_type;
-- CREATE TABLE record_type
-- (
--   type_id SMALLINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
--   type_name VARCHAR(30) NOT NULL,
--   group_id SMALLINT NOT NULL,
--   user_id INT(11) NOT NULL
-- );

-- DROP TABLE IF EXISTS record_group;
-- CREATE TABLE record_group
-- (
--   group_id SMALLINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
--   group_name VARCHAR(30) NOT NULL,
--   user_id INT(11) NOT NULL
-- );

-- INSERT INTO record_type(type_name, group_id, user_id)
-- VALUE ('刷题', 4, 3),('测试', 4, 3), ('跑步', 5, 3);

-- INSERT INTO record_group(group_name, user_id)
-- value ('睡眠', 0), ('工作', 0), ('阅读',0), ('编程',0), ('运动',0), ('娱乐',3);
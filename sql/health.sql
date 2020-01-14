DROP TABLE IF EXISTS sleep_records;

CREATE TABLE sleep_records
(
    sleep_id INT(11) PRIMARY KEY AUTO_INCREMENT,
    record_date BIGINT NOT NULL,
    start_time int(11) DEFAULT NULL,
    end_time int(11) DEFAULT NULL,
	duration int(11) DEFAULT NULL
);

DROP TABLE IF EXISTS weight_records;
CREATE TABLE weight_records
(
    weight_id INT(11) PRIMARY KEY AUTO_INCREMENT,
    record_date BIGINT NOT NULL,
    body_weight INT(11) NOT NULL,
    body_fat_rate TINYINT
);
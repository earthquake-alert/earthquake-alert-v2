# 重複チェック用
CREATE TABLE IF NOT EXISTS `JmaXmlEntries` (
    `id` VARCHAR(127) NOT NULL,
    `updated` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

# Twitterスレッド保存用
# 震度速報、震源に関する情報、震源・震度に関する情報などは同じ地震に対してなので
# これらはスレッドとして送信させる
CREATE TABLE IF NOT EXISTS `TwitterThreads` (
    `event_id` BIGINT NOT NULL,
    `tweet_id` VARCHAR(31) NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`event_id`)
);

# 地震情報
CREATE TABLE IF NOT EXISTS `Earthquakes` (
    `event_id` BIGINT NOT NULL,
    `lat` INT,
    `lon` INT,
    `depth` INT,
    `epicenter_name` TEXT,
    `max_int` VARCHAR(3) NOT NULL,
    `magnitude` INT,
    `magnitude_type` VARCHAR(3),
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`event_id`)
);

# 津波情報と地震情報をつなげるやつ
# 津波1に対して地震多or無があるので
CREATE TABLE IF NOT EXISTS `TsunamiConnects` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `tsunami_id` INT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

# 津波情報
CREATE TABLE IF NOT EXISTS `TsunamiInfos` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 震源・震度に関する情報
CREATE TABLE IF NOT EXISTS `EarthquakeInfos` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `lat` INT NOT NULL,
    `lon` INT NOT NULL,
    `depth` INT NOT NULL,
    `epicenter_name` TEXT NOT NULL,
    `max_int` VARCHAR(3),
    `magnitude` VARCHAR(3),
    `magnitude_type` VARCHAR(3),
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 震源に関する情報
CREATE TABLE IF NOT EXISTS `EarthquakeEpicenters` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `lat` INT NOT NULL,
    `lon` INT NOT NULL,
    `depth` INT NOT NULL,
    `epicenter_name` TEXT NOT NULL,
    `magnitude` INT NOT NULL,
    `magnitude_type` VARCHAR(3) NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 震度速報
CREATE TABLE IF NOT EXISTS `EarthquakeReports` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `max_int` VARCHAR(3) NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `EarthquakeActivity` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `EarthquakeUpdate` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `EarthquakeCount` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 重複チェック用
CREATE TABLE `JmaXmlEntries` (
    `id` VARCHAR(127) NOT NULL,
    `updated` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

# Twitterスレッド保存用
# 震度速報、震源に関する情報、震源・震度に関する情報などは同じ地震に対してなので
# これらはスレッドとして送信させる
CREATE TABLE `TwitterThreads` (
    `event_id` INT NOT NULL,
    `tweet_id` VARCHAR(31) NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

# 地震情報
CREATE TABLE `Earthquakes` (
    `event_id` INT NOT NULL,
    `lat` INT,
    `lon` INT,
    `depth` INT,
    `epicenter_name` TEXT,
    `max_int` VARCHAR(3) NOT NULL,
    `magnitude` INT,
    `magnitude_type` VARCHAR(3),
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

# 津波情報と地震情報をつなげるやつ
# 津波1に対して地震多or無があるので
CREATE TABLE `Tsunamis` (
    `tsunami_id` INT NOT NULL,
    `event_id` INT NOT NULL,
    `created` DATETIME NOT NULL
);

# 津波情報
CREATE TABLE `TsunamiInfos` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL
);

# 震源・震度に関する情報
CREATE TABLE `EarthquakeInfos` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` INT NOT NULL,
    `lat` INT NOT NULL,
    `lon` INT NOT NULL,
    `depth` INT NOT NULL,
    `epicenter_name` TEXT NOT NULL,
    `max_int     ` VARCHAR(3),
    `magnitude` VARCHAR(3),
    `magnitude_type` VARCHAR(3),
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL
);

# 震源に関する情報
CREATE TABLE `EarthquakeEpicenters` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` INT NOT NULL,
    `lat` INT NOT NULL,
    `lon` INT NOT NULL,
    `depth` INT NOT NULL,
    `epicenter_name` TEXT NOT NULL,
    `magnitude` INT NOT NULL,
    `magnitude_type` VARCHAR(3) NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL
);

# 震度速報
CREATE TABLE `EarthquakeReports` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` INT NOT NULL,
    `max_int` VARCHAR(3) NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL
);

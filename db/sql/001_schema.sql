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
    `event_id` BIGINT UNSIGNED NOT NULL,
    `tweet_id` VARCHAR(31) NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`event_id`)
);

# 地震情報
CREATE TABLE IF NOT EXISTS `Earthquakes` (
    `event_id` BIGINT UNSIGNED NOT NULL,
    `lat` DOUBLE,
    `lon` DOUBLE,
    `depth` INT,
    `epicenter_name` TEXT,
    `max_int` VARCHAR(3) NOT NULL,
    `magnitude` TEXT,
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
    `event_id` BIGINT UNSIGNED NOT NULL,
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
    `event_id` BIGINT UNSIGNED NOT NULL,
    `lat` DOUBLE,
    `lon` DOUBLE,
    `depth` INT,
    `epicenter_name` TEXT NOT NULL,
    `max_int` VARCHAR(3),
    `magnitude` TEXT,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 震源に関する情報
CREATE TABLE IF NOT EXISTS `EarthquakeEpicenters` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,
    `lat` DOUBLE,
    `lon` DOUBLE,
    `depth` INT,
    `epicenter_name` TEXT NOT NULL,
    `magnitude` TEXT,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 震度速報
CREATE TABLE IF NOT EXISTS `EarthquakeReports` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,
    `max_int` VARCHAR(3) NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 地震の活動状況等に関する情報
CREATE TABLE IF NOT EXISTS `EarthquakeActivity` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 顕著な地震の震源要素更新のお知らせ
CREATE TABLE IF NOT EXISTS `EarthquakeUpdate` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,
    `lat` DOUBLE,
    `lon` DOUBLE,
    `depth` INT,
    `magnitude` TEXT,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 地震回数に関する情報
CREATE TABLE IF NOT EXISTS `EarthquakeCount` (
    `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `event_id` BIGINT UNSIGNED NOT NULL,
    `date` DATETIME NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `row` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

# 都道府県とcodeの対応
CREATE TABLE IF NOT EXISTS `PrefectureCodes` (
    `id` INT UNSIGNED NOT NULL,
    `name` TEXT NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

# 地域とcodeの対応
CREATE TABLE IF NOT EXISTS `AreaCodes` (
    `id` INT UNSIGNED NOT NULL,
    `name` TEXT NOT NULL,
    `prefecture_id` INT UNSIGNED NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

# 震度観測点とcodeの対応
CREATE TABLE IF NOT EXISTS `IntensityStationCodes` (
    `id` INT UNSIGNED NOT NULL,
    `name` TEXT NOT NULL,
    `lat` DOUBLE NOT NULL,
    `lon` DOUBLE NOT NULL,
    `prefecture_id` INT UNSIGNED NOT NULL,
    `area_id` INT UNSIGNED NOT NULL,
    `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

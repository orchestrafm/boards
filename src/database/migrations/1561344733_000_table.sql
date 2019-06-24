CREATE TABLE `boards` (
    `track_id` INT(8) UNSIGNED NOT NULL,
    `id` INT(8) UNSIGNED NOT NULL,
    `date_created` DATETIME NOT NULL DEFAULT NOW(),
    `date_modified` DATETIME,
    `sha3` CHAR(512) NOT NULL,
    `jacket` BLOB NOT NULL,
    `charters` VARCHAR(128) UNSIGNED NOT NULL,
    `difficulty_name` TINYINT(3) UNSIGNED NOT NULL,
    `difficulty_rating` TINYINT(2) UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`track_id`) REFERENCES `tracks` (`id`) ON DELETE CASCADE
)

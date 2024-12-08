CREATE TABLE `homepages` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`main_image` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`main_title` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`main_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`kalkulator_title` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`kalkulator_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`pers_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`publikasi_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2
;

CREATE TABLE `donation_list` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`title` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`description` TEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`code` INT(11) NULL DEFAULT NULL,
	`donation_category_id` BIGINT(20) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `donation_list_donation_category_id_foreign` (`donation_category_id`) USING BTREE,
	CONSTRAINT `donation_list_donation_category_id_foreign` FOREIGN KEY (`donation_category_id`) REFERENCES `donation_category` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
AUTO_INCREMENT=14
;

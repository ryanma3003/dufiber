CREATE TABLE `charity_list` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`title` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`description` TEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`charity_category_id` BIGINT(20) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `charity_list_charity_category_id_foreign` (`charity_category_id`) USING BTREE,
	CONSTRAINT `charity_list_charity_category_id_foreign` FOREIGN KEY (`charity_category_id`) REFERENCES `charity_category` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
;

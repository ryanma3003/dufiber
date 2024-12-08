CREATE TABLE `donations` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`name` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`email` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`phone` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`amount` DOUBLE(8,2) NOT NULL,
	`status` VARCHAR(191) NOT NULL DEFAULT 'pending' COLLATE 'utf8mb4_unicode_ci',
	`reference` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`snap_token` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`donation_list_id` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
	`charity_list_id` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
	`users_id` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
	`orderId` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `donations_users_id_foreign` (`users_id`) USING BTREE,
	INDEX `donations_donation_list_id_foreign` (`donation_list_id`) USING BTREE,
	INDEX `donations_charity_list_id_foreign` (`charity_list_id`) USING BTREE,
	CONSTRAINT `donations_charity_list_id_foreign` FOREIGN KEY (`charity_list_id`) REFERENCES `charity_list` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE,
	CONSTRAINT `donations_donation_list_id_foreign` FOREIGN KEY (`donation_list_id`) REFERENCES `donation_list` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE,
	CONSTRAINT `donations_users_id_foreign` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
ROW_FORMAT=COMPACT
AUTO_INCREMENT=2
;

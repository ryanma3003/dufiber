CREATE TABLE `users` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`name` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`email` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`email_verified_at` TIMESTAMP NULL DEFAULT NULL,
	`password` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`role` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`remember_token` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `users_email_unique` (`email`) USING BTREE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
AUTO_INCREMENT=3
;

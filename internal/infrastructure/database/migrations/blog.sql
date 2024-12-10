CREATE TABLE `blog` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`title` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`slug` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`image` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`content` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`author` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`blog_category_id` BIGINT(20) UNSIGNED NULL DEFAULT NULL,
	`users_id` BIGINT(20) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `blog_users_id_foreign` (`users_id`) USING BTREE,
	INDEX `blog_blog_category_id_foreign` (`blog_category_id`) USING BTREE,
	CONSTRAINT `blog_blog_category_id_foreign` FOREIGN KEY (`blog_category_id`) REFERENCES `blog_category` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE,
	CONSTRAINT `blog_users_id_foreign` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
ROW_FORMAT=COMPACT
AUTO_INCREMENT=2
;

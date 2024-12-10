CREATE TABLE `galeris` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`title` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`slug` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`image` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`galery_tag_id` BIGINT(20) UNSIGNED NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `galeris_galery_tag_id_foreign` (`galery_tag_id`) USING BTREE,
	CONSTRAINT `galeris_galery_tag_id_foreign` FOREIGN KEY (`galery_tag_id`) REFERENCES `galeri_tags` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
;

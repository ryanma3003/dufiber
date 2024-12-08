CREATE TABLE `harga_zakats` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`donation_list_id` BIGINT(20) UNSIGNED NOT NULL,
	`title` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`price` DECIMAL(20,2) NOT NULL,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `harga_zakats_donation_list_id_foreign` (`donation_list_id`) USING BTREE,
	CONSTRAINT `harga_zakats_donation_list_id_foreign` FOREIGN KEY (`donation_list_id`) REFERENCES `donation_list` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
AUTO_INCREMENT=6
;

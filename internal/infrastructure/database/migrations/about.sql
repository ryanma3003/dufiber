CREATE TABLE `abouts` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`updated_at` TIMESTAMP NULL DEFAULT NULL,
	`latar_title` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`latar_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`visi_misi_title` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`visi_title` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`visi_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`misi_title` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`misi_text` MEDIUMTEXT NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`misi_text2` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_title` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_title1` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_text1` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_image1` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_title2` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_text2` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_image2` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_title3` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_text3` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_image3` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_title4` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_text4` MEDIUMTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`nilai_image4` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_unicode_ci',
	`struktur_title` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	`struktur_image` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2
;

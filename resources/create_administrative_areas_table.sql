CREATE TABLE `administrative_areas` (
  `id` VARCHAR(100) UNIQUE NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `type` VARCHAR(25) NOT NULL,
  `parent_id` VARCHAR(100) NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);
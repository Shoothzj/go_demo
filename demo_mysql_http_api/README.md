## refer
https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d
## sql
```sql
  CREATE TABLE `messages` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NULL,
  `body` VARCHAR(200) NULL,
  `created_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `title_UNIQUE` (`title` ASC));
```
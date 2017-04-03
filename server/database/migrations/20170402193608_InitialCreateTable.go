package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up_20170402193608, Down_20170402193608)
}

func Up_20170402193608(tx *sql.Tx) error {
	sql :=
		"DROP TABLE IF EXISTS `readers`; " +
			"DROP TABLE IF EXISTS `users`; " +
			"DROP TABLE IF EXISTS `links`; " +
			"DROP TABLE IF EXISTS `comments`; " +
			"DROP TABLE IF EXISTS `entries`; " +
			"CREATE TABLE `entries` ( " +
			"`id` int(11) NOT NULL AUTO_INCREMENT, " +
			"`user_id` int(11) NOT NULL DEFAULT '0', " +
			"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
			"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " +
			"`visible` tinyint(4) NOT NULL DEFAULT '0', " +
			"`entry` text, " +
			"`title` varchar(255) DEFAULT NULL, " +
			"PRIMARY KEY (`id`), " +
			"UNIQUE KEY `idx_unique_title` (`title`) USING BTREE, " +
			"KEY `fk_users` (`user_id`), " +
			"CONSTRAINT `fk_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION " +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8; " +
			"CREATE TABLE `comments` ( " +
			"`id` int(11) NOT NULL AUTO_INCREMENT, " +
			"`entry_id` int(11) NOT NULL DEFAULT '0', " +
			"`user_id` int(11) NOT NULL DEFAULT '0', " +
			"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
			"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, " +
			"`comment` text, " +
			"PRIMARY KEY (`id`), " +
			"KEY `fk_user_id` (`user_id`), " +
			"KEY `fk_entry_id` (`entry_id`), " +
			"CONSTRAINT `fk_entry_id` FOREIGN KEY (`entry_id`) REFERENCES `entries` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION, " +
			"CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION " +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8; " +
			"CREATE TABLE `links` ( " +
			"`id` int(11) NOT NULL AUTO_INCREMENT, " +
			"`entry_id` int(11) NOT NULL DEFAULT '0', " +
			"`link` varchar(255) NOT NULL DEFAULT '', " +
			"PRIMARY KEY (`id`), " +
			"KEY `fk_entry_link_id` (`entry_id`), " +
			"CONSTRAINT `fk_entry_link_id` FOREIGN KEY (`entry_id`) REFERENCES `entries` (`id`) ON DELETE CASCADE " +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8; " +
			"CREATE TABLE `users` ( " +
			"`id` int(11) NOT NULL AUTO_INCREMENT, " +
			"`first_name` varchar(128) NOT NULL DEFAULT '', " +
			"`last_name` varchar(128) DEFAULT NULL, " +
			"`middle_name` varchar(128) DEFAULT NULL, " +
			"`role` smallint(6) NOT NULL DEFAULT '0', " +
			"`email` varchar(128) NOT NULL DEFAULT '', " +
			"`password` varbinary(255) NOT NULL DEFAULT '', " +
			"`salt` varchar(255) NOT NULL DEFAULT '', " +
			"PRIMARY KEY (`id`), " +
			"UNIQUE KEY `idx_email` (`email`) USING BTREE " +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8; " +
			"CREATE TABLE `readers` ( " +
			"`id` int(11) NOT NULL AUTO_INCREMENT, " +
			"`writer_id` int(11) NOT NULL DEFAULT '0', " +
			"`reader_id` int(11) NOT NULL DEFAULT '0', " +
			"PRIMARY KEY (`id`), " +
			"KEY `fk_reader_id` (`reader_id`), " +
			"KEY `fk_writer_id` (`writer_id`), " +
			"CONSTRAINT `fk_reader_id` FOREIGN KEY (`reader_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION, " +
			"CONSTRAINT `fk_writer_id` FOREIGN KEY (`writer_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION " +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8; "
	_, err := tx.Exec(sql)
	return err
}

func Down_20170402193608(tx *sql.Tx) error {
	sql :=
		"DROP TABLE IF EXISTS `readers`; " +
			"DROP TABLE IF EXISTS `users`; " +
			"DROP TABLE IF EXISTS `links`; " +
			"DROP TABLE IF EXISTS `comments`; " +
			"DROP TABLE IF EXISTS `entries`; "
	_, err := tx.Exec(sql)
	return err
}

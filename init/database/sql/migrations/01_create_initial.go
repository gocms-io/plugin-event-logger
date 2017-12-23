package migrations

import (
	"github.com/rubenv/sql-migrate"
)

func CreateInitial() *migrate.Migration {

	create_initial := migrate.Migration{
		Id: "1",
		Up: []string{`
			CREATE TABLE gocms_plugin_event_logger_settings (
			id int(11) NOT NULL AUTO_INCREMENT,
			name varchar(30) NOT NULL UNIQUE,
			value blob NOT NULL,
			description varchar(255) NOT NULL,
			created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			lastModified DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
			`, `
	        CREATE TABLE gocms_plugin_event_logger_request (
			id int(11) NOT NULL AUTO_INCREMENT,
			path text NOT NULL,
			type varchar(10) NOT NULL,
			created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
			`, `
			CREATE TABLE gocms_plugin_event_logger_request_headers (
			id int(11) NOT NULL AUTO_INCREMENT,
			requestId int(11) NOT NULL,
			created datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (requestId)
				REFERENCES gocms_plugin_event_logger_request (id)
				ON DELETE CASCADE
			) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
			`, `
			INSERT INTO gocms_plugin_event_logger_settings (name, value, description) VALUES('SETTINGS_REFRESH_RATE', '60', 'Minutes between each settings refresh from the database to memory.');
			`,
		},
		Down: []string{
			"DROP TABLE gocms_plugin_event_logger_request_headers;",
			"DROP TABLE gocms_plugin_event_logger_request;",
			"DROP TABLE gocms_plugin_event_logger_settings;",
		},
	}

	return &create_initial
}

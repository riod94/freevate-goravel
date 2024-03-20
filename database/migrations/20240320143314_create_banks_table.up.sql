CREATE TABLE banks (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  code varchar(255) NOT NULL UNIQUE,
  status enum('active', 'inactive') NOT NULL DEFAULT 'active',
  category enum('domestic', 'regional', 'foreign', 'private', 'sharia', 'other') NOT NULL DEFAULT 'private',
  link varchar(255),
  contact varchar(255),
  logo_filepath varchar(255),
  created_at datetime(3) NOT NULL,
  updated_at datetime(3) NOT NULL,
  deleted_at datetime(3),
  PRIMARY KEY (id),
  KEY idx_banks_created_at (created_at),
  KEY idx_banks_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS users (
  id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  fullname text NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now()),
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS vendors (
  id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  fullname text NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  password text NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now()),
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS products (
  id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  vendor_id bigint unsigned NOT NULL,
  name varchar(255) NOT NULL,
  photo_url text NOT NULL,
  description text,
  price decimal NOT NULL DEFAULT 0,
  stock bigint unsigned NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT (now()),
  created_by bigint unsigned NOT NULL,
  updated_at timestamp,
  updated_by bigint,
  FOREIGN KEY (vendor_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS checkout (
  guid varchar(36) NOT NULL PRIMARY KEY DEFAULT(uuid()),
  user_id bigint unsigned NOT NULL,
  product_id bigint unsigned NOT NULL,
  qty int NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT (now()),
  updated_at timestamp,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
ALTER TABLE users ADD COLUMN id UUID;
ALTER TABLE users ADD COLUMN first_name varchar(128);
ALTER TABLE users ADD COLUMN last_name varchar(128);
ALTER TABLE users ADD COLUMN acces_token varchar(256);
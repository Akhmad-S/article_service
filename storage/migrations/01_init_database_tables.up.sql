CREATE TABLE author (
    id	CHAR(36) PRIMARY KEY,
	first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE article (
    id CHAR(36) PRIMARY KEY,
	title VARCHAR(255) UNIQUE NOT NULL,
	body text NOT NULL,
	author_id CHAR(36),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES author (id);

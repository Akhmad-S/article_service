ALTER TABLE article DROP CONSTRAINT IF EXISTS fk_article_author;
DROP TABLE author;
DROP TABLE article;
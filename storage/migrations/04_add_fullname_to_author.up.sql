ALTER TABLE author ADD COLUMN fullname VARCHAR(255);

UPDATE author SET fullname = first_name || ' ' || last_name ||
(SELECT CASE WHEN middle_name IS NULL THEN '' ELSE  ' ' || middle_name END AS middlename);

ALTER TABLE author DROP COLUMN first_name;
ALTER TABLE author DROP COLUMN last_name;
ALTER TABLE author DROP COLUMN middle_name;

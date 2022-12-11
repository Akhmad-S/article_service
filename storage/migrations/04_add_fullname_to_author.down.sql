ALTER TABLE author ADD COLUMN first_name VARCHAR(55);
ALTER TABLE author ADD COLUMN last_name VARCHAR(55);
ALTER TABLE author ADD COLUMN middle_name VARCHAR(55);



UPDATE author SET first_name = (SELECT split_part(fullname, ' ', 1));
UPDATE author SET last_name = (SELECT split_part(fullname, ' ', 2));
UPDATE author SET middle_name = (SELECT split_part(fullname, ' ', 3));

ALTER TABLE author DROP COLUMN fullname;
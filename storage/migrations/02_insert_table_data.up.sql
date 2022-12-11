INSERT INTO author (id, first_name, last_name) VALUES ('f4edbadf-6153-4d31-8955-918af3f967a4', 'John', 'Doe')ON CONFLICT DO NOTHING;
INSERT INTO author (id, first_name, last_name) VALUES ('349b2748-3480-4c33-ac67-6e44d23555fe', 'Peter', 'Parker')ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('1b569b84-48d9-414f-8882-b265c1dec5cc', 'Lorem', 'Ipsume lorem something...', 'f4edbadf-6153-4d31-8955-918af3f967a4')ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('27322749-999b-46da-82a6-17ef71dc2bfe', 'Spiderman', 'About hero Spiderman...', '349b2748-3480-4c33-ac67-6e44d23555fe')ON CONFLICT DO NOTHING;

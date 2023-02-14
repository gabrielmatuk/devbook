insert into usuarios(nome, nick, email, senha)
values
("user 1", "user_1", "usuario1@gmail.com", "$2a$10$B1VVtynrJo22FUwxHUmJTOdyzGjFsLZMJJg6CMJwGph/o5xQAQDnG"),
("user 2", "user_2", "usuario2@gmail.com", "$2a$10$B1VVtynrJo22FUwxHUmJTOdyzGjFsLZMJJg6CMJwGph/o5xQAQDnG"),
("user 3", "user_3", "usuario3@gmail.com", "$2a$10$B1VVtynrJo22FUwxHUmJTOdyzGjFsLZMJJg6CMJwGph/o5xQAQDnG");
-- pswd 156416854
insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3),
(2,3);
INSERT INTO users (name, nick, email, password)
VALUES
("Usuario 1", "user1", "user1@email.com.br", "$2a$10$c4ptxqojf.doZ756ryi9J.tph6wSlP9jsnpq1/HXxHkUAn5EosQjm"),
("Usuario 2", "user2", "user2@email.com.br", "$2a$10$c4ptxqojf.doZ756ryi9J.tph6wSlP9jsnpq1/HXxHkUAn5EosQjm"),
("Usuario 3", "user3", "user3@email.com.br", "$2a$10$c4ptxqojf.doZ756ryi9J.tph6wSlP9jsnpq1/HXxHkUAn5EosQjm"),
("Usuario 4", "user4", "user4@email.com.br", "$2a$10$c4ptxqojf.doZ756ryi9J.tph6wSlP9jsnpq1/HXxHkUAn5EosQjm");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 2),
(1, 3);

INSERT INTO publications (title, content, created_id)
VALUES
("Pulicacao 1", "Conteudo 1", 1),
("Pulicacao 2", "Conteudo 2", 2),
("Pulicacao 3", "Conteudo 3", 3);
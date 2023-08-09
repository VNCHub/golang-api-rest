create table characters(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO characters(name, history) VALUES
('Harry Potter', 'O Menino Que Sobreviveu.'),
('Hermione', 'Inteligente e curiosa.'),
('Ron Weasley', 'Amigo leal.'),
('Dumbledore',	'Diretor de Hogwarts.'),
('Voldemort', 'O Lorde das Trevas.'),
('Snape', 'Complexo e misterioso.'),
('Hagrid', 'Gentil, amigo dos animais.'),
('Luna Lovegood', 'Estranha e sonhadora.'),
('Draco Malfoy', 'Nasceu em uma família de comensais.'),
('McGonagall', 'Professora de Transfiguração.');
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE News (
    Id SERIAL PRIMARY KEY,
    Title TEXT NOT NULL,
    Content TEXT NOT NULL
);

CREATE TABLE NewsCategories (
    NewsId BIGINT NOT NULL,
    CategoryId BIGINT NOT NULL,
    PRIMARY KEY (NewsId, CategoryId)
);


--Добавляем новости
INSERT INTO News (Id, Title, Content)
VALUES (64, 'Lorem ipsum', 'Dolor sit amet <b>foo</b>');

--Добавленим катеогрии для новости с Id=64
INSERT INTO NewsCategories (NewsId, CategoryId) VALUES (64,1);
INSERT INTO NewsCategories (NewsId, CategoryId) VALUES (64,2);
INSERT INTO NewsCategories (NewsId, CategoryId) VALUES (64,3);

--Добавляем новости
INSERT INTO News(Id, Title, Content)
VALUES (1, 'first', 'tratata');

--Добавленим катеогрии для новости с Id=1
INSERT INTO NewsCategories (NewsId, CategoryId) VALUES (1,1);
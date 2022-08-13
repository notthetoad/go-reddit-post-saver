DROP TABLE IF EXISTS post;
CREATE TABLE post (
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    body VARCHAR(255),
    url VARCHAR(255) NOT NULL
    PRIMARY KEY (`id`)
)

INSERT INTO post (title, body, url) VALUES 
    ('Foo', 'foobarbaz', 'example.org')
    ('Bar', 'barfoobaz', 'example.org/1')
    ('Baz', 'bazbarfoo', 'example.org/2')

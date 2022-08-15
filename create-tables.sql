DROP TABLE IF EXISTS goreddit;
CREATE TABLE post (
    id INT VARCHAR(128) NOT NULL,
    url VARCHAR(255) NOT NULL,
    title VARCHAR(128) NOT NULL,
    permalink VARCHAR(255) NOT NULL,
    selftext VARCHAR(255),
    subreddit VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
)

INSERT INTO post (title, body, url) VALUES 
    ('Foo', 'foobarbaz', 'example.org')
    ('Bar', 'barfoobaz', 'example.org/1')
    ('Baz', 'bazbarfoo', 'example.org/2')

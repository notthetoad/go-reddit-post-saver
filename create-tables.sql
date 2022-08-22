CREATE TABLE reddit_post (
    id INT AUTO_INCREMENT NOT NULL,
    post_id VARCHAR(64) NOT NULL UNIQUE,
    permalink VARCHAR(128) NOT NULL,
    url VARCHAR(255) NOT NULL,
    title VARCHAR(512) NOT NULL,
    selftext TEXT,
    subreddit VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE reddit_comment (
    id INT AUTO_INCREMENT NOT NULL,
    comment_id VARCHAR(64) NOT NULL UNIQUE,
    permalink VARCHAR(128) NOT NULL,
    selftext TEXT,
    subreddit VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);

--CREATE TABLE misc_data (
--    id INT AUTO_INCREMENT NOT NULL,
--    filename VARCHAR(255),
--    type_key
--)
--
--CREATE TABLE misc_data_type (
--    id INT AUTO_INCREMENT,
--
--)

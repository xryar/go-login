CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    fullname VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDB;
CREATE TABLE songs(
    id INT NOT NULL AUTO_INCREMENT,
    title TEXT NOT NULL,
    year TEXT NOT NULL,
    genre TEXT NOT NULL,
    performer TEXT NOT NULL,
    duration INT NOT NULL,
    album_id INT NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_album_id FOREIGN KEY (album_id)
        REFERENCES albums(id)
        ON DELETE CASCADE
) ENGINE = InnoDB;
CREATE TABLE TodoList(
    id SERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    descriptions VARCHAR(65535) NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP,
    PRIMARY KEY(id)
);
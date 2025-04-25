-- +goose Up
CREATE TABLE PRODUCTS
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR   NOT NULL,
    description VARCHAR   NOT NULL,
    uom         VARCHAR   NOT NULL,
    image       VARCHAR   NOT NULL,
    price       NUMERIC   NOT NULL DEFAULT 0,
    stock       INT       NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    discarded_at TIMESTAMP          DEFAULT NULL
);


INSERT INTO PRODUCTS (name, description, uom, image, price, stock)
VALUES ('SANDAL',
        'Match on your feet',
        'GRAM',
        'https://www.google.com/imgres?q=onrus&imgurl=https%3A%2F%2Fdown-id.img.susercontent.com%2Ffile%2Fid-11134207-7rasi-m2uoca47px1tac&imgrefurl=https%3A%2F%2Fshopee.co.id%2FONRUS-OnHike-Navis-Sandal-Gunung-Sandal-Hiking-Sandal-Tracking-Sandal-Jepit-Gunung-i.1302211767.26611372951&docid=a5CUtk0edGzHVM&tbnid=AJ80rTfsxbrHZM&vet=12ahUKEwiUzN-fpfOMAxWRRmcHHSuTAnYQM3oECB4QAA..i&w=1024&h=1024&hcb=2&ved=2ahUKEwiUzN-fpfOMAxWRRmcHHSuTAnYQM3oECB4QAA',
        95000,
        1000);
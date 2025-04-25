-- +goose Up
INSERT INTO PRODUCTS (name, description, uom, image, price, stock)
VALUES ('SEPATU',
        'Match on your feet',
        'GRAM',
        'https://www.google.com/imgres?q=onrus&imgurl=https%3A%2F%2Fdown-id.img.susercontent.com%2Ffile%2Fid-11134207-7rasi-m2uoca47px1tac&imgrefurl=https%3A%2F%2Fshopee.co.id%2FONRUS-OnHike-Navis-Sandal-Gunung-Sandal-Hiking-Sandal-Tracking-Sandal-Jepit-Gunung-i.1302211767.26611372951&docid=a5CUtk0edGzHVM&tbnid=AJ80rTfsxbrHZM&vet=12ahUKEwiUzN-fpfOMAxWRRmcHHSuTAnYQM3oECB4QAA..i&w=1024&h=1024&hcb=2&ved=2ahUKEwiUzN-fpfOMAxWRRmcHHSuTAnYQM3oECB4QAA',
        250000,
        1000);
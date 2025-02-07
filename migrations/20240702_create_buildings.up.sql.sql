CREATE TABLE IF NOT EXISTS cities
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS years
(
    id   BIGSERIAL PRIMARY KEY,
    year INT NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS floors
(
    id    BIGSERIAL PRIMARY KEY,
    count INT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS buildings
(
    id       BIGSERIAL PRIMARY KEY,
    name     TEXT   NOT NULL,
    city_id  BIGINT NOT NULL REFERENCES cities (id) ON DELETE CASCADE,
    year_id  BIGINT NOT NULL REFERENCES years (id) ON DELETE CASCADE,
    floor_id BIGINT NOT NULL REFERENCES floors (id) ON DELETE CASCADE
);

ALTER TABLE buildings
    ADD CONSTRAINT unique_city_year_floor_name
        UNIQUE (city_id, year_id, floor_id, name);
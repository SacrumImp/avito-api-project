CREATE TABLE IF NOT EXISTS user_type (
    id SERIAL PRIMARY KEY,
    title Text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(254) NOT NULL,
    password_hash VARCHAR(50) NOT NULL,
    user_type_id SMALLINT NOT NULL REFERENCES user_type(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS developer (
    id SERIAL PRIMARY KEY,
    title Text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS house (
    id SERIAL PRIMARY KEY,
    address Text UNIQUE NOT NULL,
    year_of_construction SMALLINT NOT NULL,
    developer_id SMALLINT REFERENCES developer(id) ON DELETE CASCADE,
    inserted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS flat (
    house_id INTEGER UNIQUE REFERENCES house(id) ON DELETE CASCADE,
    number SMALLINT,
    price BIGINT NOT NULL,
    NumberOfRooms SMALLINT NOT NULL,
    PRIMARY KEY (house_id, number)
);

CREATE INDEX idx_flats_by_house ON flat(house_id);
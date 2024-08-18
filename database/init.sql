CREATE TABLE IF NOT EXISTS user_type (
    id SERIAL PRIMARY KEY,
    title Text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(254) NOT NULL,
    password_hash VARCHAR(50) NOT NULL,
    user_type_id INTEGER NOT NULL REFERENCES user_type(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS developer (
    id SERIAL PRIMARY KEY,
    title Text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS house (
    id SERIAL PRIMARY KEY,
    address Text UNIQUE NOT NULL,
    year_of_construction SMALLINT NOT NULL,
    developer_id INTEGER REFERENCES developer(id) ON DELETE CASCADE,
    inserted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_flat_added_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS status (
    id SERIAL PRIMARY KEY,
    title Text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS flat (
    house_id INTEGER REFERENCES house(id) ON DELETE CASCADE,
    number SMALLINT,
    price BIGINT NOT NULL,
    number_of_rooms SMALLINT NOT NULL,
    status_id INTEGER NOT NULL REFERENCES status(id) ON DELETE CASCADE,
    PRIMARY KEY (house_id, number)
);

CREATE INDEX idx_flats_by_house ON flat(house_id);

CREATE OR REPLACE FUNCTION update_flat_added_at()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE house
    SET last_flat_added_at = NOW()
    WHERE id = NEW.house_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_flat_insert
AFTER INSERT ON flat
FOR EACH ROW
EXECUTE FUNCTION update_flat_added_at();

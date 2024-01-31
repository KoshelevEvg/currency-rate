CREATE SCHEMA IF NOT EXISTS test;

CREATE TABLE IF NOT EXISTS test.name_currency (
      currency_id serial UNIQUE ,
      char_name varchar,
      name varchar
);

CREATE TABLE IF NOT EXISTS test.date_currency (
    id serial primary key,
    date_start timestamp,
    date_end timestamp,
    value decimal,
    currency_id integer references test.name_currency(currency_id)
);

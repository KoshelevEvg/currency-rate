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



INSERT INTO test.name_currency (currency_id, char_name, name) VALUES (2, 'test1', 'ttt1') returning name_currency.currency_id;


SELECT to_char(date_start, 'yyyy/mm/dd'), date_end, value, char_name, name from test.date_currency as t right join test.name_currency on t.currency_id = name_currency.currency_id
where char_name = 'AUD' and to_char(date_start, 'yyyy/mm/dd') = '2024/01/19';

SELECT date_start, date_end, value, char_name, name from test.date_currency as t right join test.name_currency on t.currency_id = name_currency.currency_id
where char_name = 'AUD' and date_start = '';
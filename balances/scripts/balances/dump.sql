DROP TABLE IF EXISTS balances;

CREATE TABLE balances (
    id varchar(255) primary key, 
    account_id varchar(255), 
    amount float not null,
    updated_at datetime not null
);

INSERT INTO balances (id, account_id, amount, updated_at) values ('9ca862ce-193e-408d-8061-d4679b51b56a', '8f11f884-1ec3-49e3-82c7-d51057983d83', 1000, now());
INSERT INTO balances (id, account_id, amount, updated_at) values ('baa05ded-2c04-42d6-ba04-81b53c68a313', 'd81ecac7-9af9-4005-b58c-1fd51f6a4720', 1000, now());
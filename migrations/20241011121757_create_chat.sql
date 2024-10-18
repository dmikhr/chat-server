-- +goose Up
-- +goose StatementBegin
begin;

create table chat (
    id serial primary key,
    name text not null,
    created_at timestamp with time zone NOT NULL
);

-- для хранения пользователей принадлежащих определенному чату
create table chat_users
(
    id serial primary key,
    chat_id int not null,
    user_id int not null,
    FOREIGN KEY (chat_id) REFERENCES chat(id) ON DELETE CASCADE
);

create table message (
    id serial primary key,
    name text not null,
    msg text not null,
    created_at timestamp with time zone NOT NULL
);

commit;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
begin;
DROP TABLE chat;
DROP TABLE chat_users;
DROP TABLE message;
commit;

-- +goose StatementEnd

create table songs(
    id serial primary key,
    created_at timestamp not null default now(),
    title text,
    song_path text,
    image_path text,
    author text,
    user_id integer references users (id) on delete CASCADE
);
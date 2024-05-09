create table liked_songs (
    user_id int not null references users (id) on delete cascade,
    song_id int not null references songs (id) on delete cascade,
    created_at timestamp default now(),
    PRIMARY KEY(user_id, song_id)
);
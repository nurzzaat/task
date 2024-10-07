create table if not exists
    songs (
        id serial primary key,
        group_name text default '' not null,
        name text default '' not null,
        release_date text default '' not null,
        lyric text default '' not null,
        link text default ''
    );
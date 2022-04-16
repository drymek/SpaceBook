create table bookings
(
    id             serial
            constraint bookings_pk
            primary key,
    custom_id      varchar(255) not null,
    firstname      varchar(255) not null,
    lastname       varchar(255) not null,
    gender         varchar(32)  not null,
    birthday       date         not null,
    launchpad_id   varchar(255),
    destination_id varchar(255) not null,
    "launch_date"   date         not null
);

create unique index bookings_custom_id_uindex
    on bookings (custom_id);

create index bookings_id_index
    on bookings (id);

create unique index bookings_id_uindex
    on bookings (id);


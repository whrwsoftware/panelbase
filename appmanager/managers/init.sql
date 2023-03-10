create table if not exists app_info
(
    name     varchar(200)  not null default '' primary key,
    icon     varchar(2000) not null default '',
    provider varchar(200)  not null default '',
    workdir  varchar(500)  not null default '',
    app_type integer       not null default 0
);

create table if not exists app_version
(
    app_name   varchar(200) not null default '',
    version    varchar(50)  not null default '',
    version_id integer      not null default 0,
    installed  integer      not null default 0,
    log_path   varchar(500) not null default '',
    log_file   varchar(200) not null default '',
    primary key (app_name, version)
);

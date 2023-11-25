CREATE TABLE IF NOT EXISTS test_table
(
    data text
);

create unique index idx_data on test_table (data);
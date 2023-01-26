create table default.errata
(
    errata_id            String,
    errata_prefix        String,
    errata_num           Int64,
    errata_update_count  Int32,
    errata_creation_date DateTime,
    errata_change_date   DateTime
)
    engine = MergeTree ORDER BY errata_prefix
        SETTINGS index_granularity = 2;
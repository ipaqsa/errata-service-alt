create table if not exists ErrataID
(
    errata_id            String,
    errata_prefix        String,
    errata_year          UInt32,
    errata_num           UInt32,
    errata_update_count  UInt32,
    errate_creation_date DateTime,
    errata_change_date   DateTime
) engine = MergeTree ORDER BY errata_prefix
        SETTINGS index_granularity = 2;
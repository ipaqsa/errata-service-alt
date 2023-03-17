CREATE TABLE IF NOT EXISTS ErrataID
(
    errata_id            String,
    errata_prefix        String,
    errata_year          UInt32,
    errata_num           UInt32,
    errata_update_count  UInt32,
    errata_creation_date DateTime,
    errata_change_date   DateTime
)
ENGINE = MergeTree
ORDER BY (errata_prefix, errata_year);

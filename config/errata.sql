create table if not exists errata
(
    errataID            String,
    errataPrefix        String,
    errataNum           Int64,
    errataUpdateCount  Int32,
    errataCreationDate DateTime,
    errataChangeDate   DateTime
)
    engine = MergeTree ORDER BY errataPrefix
        SETTINGS index_granularity = 2;
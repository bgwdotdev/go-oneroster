CREATE TABLE IF NOT EXISTS orgs (
    Id Integer PRIMARY KEY AUTOINCREMENT,
    sourcedId text,
    status text,
    dateLastModified text,
    name text,
    type text,
    identifier text,
    parentSourcedId text
);

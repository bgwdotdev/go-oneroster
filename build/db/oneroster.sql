CREATE TABLE IF NOT EXISTS "orgs" (
    "Id" Integer PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "name" text,
    "type" text,
    "identifier" text,
    "parentSourcedId" text,
    CONTRAINT "FK_orgs_orgs_parentSourcedId" FOREIGN KEY ("parentSourcedId") REFERENCES "orgs" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "academicSessions" (
    "Id" Integer PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "title" text,
    "type" text,
    "startDate" text,
    "endDate" text,
    "parentSourcedId" text,
    "schoolYear" text,
    CONSTRAINT "FK_academicSessions_academicSessions_parentSourcedId" FOREIGN KEY ("parentSourcedId") REFERENCES "academicSessions" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "courses" (
    "Id" Integer PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "schoolYearSourcedId" text,
    "title" text,
    "courseCode" text,
    "grades" text,
    "orgSourcedId" text,
    "subjects" text,
    "subjectCodes" text,
    CONSTRAINT "FK_courses_academicSessions_schoolYearSourcedId" FOREIGN KEY ("schoolYearSourcedId") REFERENCES "academicSessions" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "classes" (

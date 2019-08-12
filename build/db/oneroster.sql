CREATE TABLE IF NOT EXISTS "orgs" (
    "Id" Integer PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "name" text,
    "type" text,
    "identifier" text,
    "parentSourcedId" text,
    CONSTRAINT "FK_orgs_orgs_parentSourcedId" 
        FOREIGN KEY ("parentSourcedId")
        REFERENCES "orgs" ("sourcedId")
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
    CONSTRAINT "FK_academicSessions_academicSessions_parentSourcedId" 
        FOREIGN KEY ("parentSourcedId")
        REFERENCES "academicSessions" ("sourcedId")
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
    CONSTRAINT "FK_courses_academicSessions_schoolYearSourcedId" 
        FOREIGN KEY ("schoolYearSourcedId") 
        REFERENCES "academicSessions" ("sourcedId"),
    CONSTRAINT "FK_courses_orgs_orgSourcedId" 
        FOREIGN KEY ("orgSourcedId") 
        REFERENCES "orgs" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "classes" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "title" text,
    "grades" text,
    "courseSourcedId" text,
    "classCode" text,
    "classType" text,
    "location" text,
    "schoolSourcedId" text,
    "termSourcedIds" text,
    "subjects" text,
    "subjectCodes" text,
    "periods" text,
    CONSTRAINT "FK_classes_courses_courseSourcedId"
        FOREIGN KEY ("courseSourcedId")
        REFERENCES "courses" ("surcedId"),
    CONSTRAINT "FK_classes_orgs_schoolSourcedId"
        FOREIGN KEY ("schoolSourcedId")
        REFERENCES "orgs" ("sourcedId"),
    CONSTRAINT "FK_classes_academicSessions_termSourcedId"
        FOREIGN KEY ("termSourcedIds")
        REFERENCES "academicSessions" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "enabledUser" boolean,
    "orgSourcedIds" text,
    "role" text,
    "username" text,
    "userIds" text,
    "givenName" text,
    "familyName" text,
    "middleName" text,
    "identifier" text,
    "email" text,
    "sms" text,
    "phone" text,
    "agentSourcedIds" text,
    "grades" text,
    "password" text,
    CONSTRAINT "FK_users_orgs_orgSourcedId"
        FOREIGN KEY ("orgSourcedIds")
        REFERENCES "orgs" ("sourcedId"),
    CONSTRAINT "FK_users_users_agentSourcedIds"
        FOREIGN KEY ("agentSourcedIds")
        REFERENCES "users" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "enrollments" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "sourcedId" text,
    "status" text,
    "dateLastModified" text,
    "classSourcedId" text,
    "schoolSourcedId" text,
    "userSourcedId" text,
    "role" text,
    "primary" text,
    "beginDate" text,
    "endDate" text,
    CONSTRAINT "FK_enrollments_classes_classSourcedId"
        FOREIGN KEY ("classSourcedId")
        REFERENCES "classes" ("sourcedId"),
    CONSTRAINT "FK_enrollments_orgs_schoolSourcedId"
        FOREIGN KEY ("schoolSourcedId") 
        REFERENCES "orgs" ("sourcedId"),
    CONSTRAINT "FK_enrollments_users_userSourcedId"
        FOREIGN KEY ("userSourcedId")
        REFERENCES "users" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "userAgents" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "agentSourcedId" text,
    "userSourcedId" text,
    CONSTRAINT "FK_userAgents_users_agentSourcedId" 
        FOREIGN KEY ("agentSourcedId")
        REFERENCES "users" ("sourcedId"),
    CONSTRAINT "FK_userAgents_users_subjectSourcedId"
        FOREIGN KEY ("userSourcedId")
        REFERENCES "users" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "userOrgs" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "orgSourcedId" text,
    "userSourcedId" text,
    CONSTRAINT "FK_userOrgs_orgs_orgSourcedId"
        FOREIGN KEY ("orgSourcedId")
        REFERENCES "orgs" ("sourcedId"),
    CONSTRAINT "FK_userOrgs_users_userSourcedId"
        FOREIGN KEY ("userSourcedId")
        REFERENCES "users" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "userIds" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "userSourcedId" text,
    "type" text,
    "identifier" text,
    CONSTRAINT "FK_userIds_users_userSourcedId"
        FOREIGN KEY ("userSourcedId")
        REFERENCES "users" ("sourcedId")
);

CREATE TABLE IF NOT EXISTS "classAcademicSessions" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "academicSessionSourcedId" text,
    "classSourcedId" text,
    CONSTRAINT "FK_classAcademicSessions_academicSessions_academicSessionSourcedId"
        FOREIGN KEY ("academicSessionSourcedId")
        REFERENCES "academicSessions" ("sourcedId"),
    CONSTRAINT "FK_classAcademicSessions_classes_classSourcedId"
        FOREIGN KEY ("classSourcedId")
        REFERENCES "classes" ("sourcedId")
);


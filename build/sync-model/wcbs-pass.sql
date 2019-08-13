/** Orgs **/
SELECT
    SCHOOL_ID AS sourcedId,
    CODE as name,
    'school' AS type,
    DESCRIPTION AS identifier,
    LAST_AMEND_DATE AS dateLastModified,
    IN_USE AS status,
FROM 
    dbo.SCHOOL
ORDER BY
    SCHOOL_ID

/** academicSession - academic Years **/
SELECT
    Y.YEAR_ID AS sourcedId,
    Y.IN_USE AS status,
    Y.LAST_AMEND_DATE AS dateLastModified,
    Y.DESCRIPTION AS title,
    SC.YEAR_START AS startDate,
    SC.YEAR_END AS endDate,
    Y.CODE AS schoolYear
FROM 
    dbo.YEAR AS Y INNER JOIN
    dbo.SCHOOL_CALENDAR AS SC ON SC.ACADEMIC_YEAR = Y.CODE
ORDER BY 
    Y.YEAR_ID
/** academicSession - terms **/

/** courses **/
SELECT
    S.SUBJECT_ID AS sourcedId,
    S.IN_USE AS status,
    S.LAST_AMEND_DATE AS dateLastModified,
    /* schoolYearSourcedId */
    S.DESCRIPTION AS title,
    S.CODE AS courseCode,
    /* grades - N/A */
    org.SCHOOL_ID AS orgSourcedId,
    S.DESCRIPTION AS subjects
    /* subjectCodes - SQA codes? */ 
FROM
    dbo.SUBJECT AS S 
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = S.SCHOOL
ORDER BY
    S.SUBJECT_ID

/** classes - scheduled **/ 
SELECT
    S.SUBJECT_SET_ID AS sourcedId,
    S.IN_USE AS status,
    S.LAST_AMEND_DATE AS dateLastModified,
    S.DESCRIPTION AS title,
    /* null AS grades, */
    SUB.SUBJECT_ID AS courseSoucedId,
    S.SET_CODE AS classCode,
    'scheduled' AS classType,
    S.ROOM AS location,
    org.SCHOOL_ID AS schoolSourcedId,
    S.SUBJECT_SET_ID AS termSourcedIds, 
    SUB.DESCRIPTION AS subjects
    /* SQA CODES? AS subjectCodes, */
    /* null AS periods */ 
FROM
    dbo.SUBJECT_SET AS S
        INNER JOIN
    dbo.SCHOOL as org
        on org.CODE = S.SCHOOL
        INNER JOIN
    dbo.YEAR as Y
        on Y.CODE = S.ACADEMIC_YEAR
        INNER JOIN
    dbo.SUBJECT AS SUB
        ON SUB.CODE = S.SUBJECT
ORDER BY
    S.SUBJECT_SET_ID    
/** classes - homeroom(form) **/
SELECT
    F.FORM_ID AS sourcedId,
    F.IN_USE AS status,
    F.LAST_AMEND_DATE AS dateLastModified,
    F.DESCRIPTION AS title,
    /* null AS grades, */
    /* null AS courseSourcedId */
    F.CODE AS classCode,
    'homeroom' AS classType,
    F.ROOM AS location,
    org.SCHOOL_ID AS schoolSourcedId,
    F.FORM_ID AS termSourcedIds
    /* null AS subjects */
    /* null AS subjectCodes */
    /* null as periods */
FROM
    dbo.FORM AS F
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = F.SCHOOL
WHERE
    F.ACADEMIC_YEAR = '2019'
ORDER BY
    F.FORM_ID

/** users - pupils **/
SELECT
    P.NAME_ID AS sourcedId,
    P.IN_USE AS status,
    P.LAST_AMEND_DATE as dateLastModified,
    P.IN_USE AS enalbedUser,
    P.NAME_ID AS orgSourcedIds,
    'student' AS role,
    N.EMAIL_ADDRESS AS username,
    P.NAME_ID AS userIds,
    N.PREFERRED_NAME AS givenName,
    N.SURNAME AS familyName,
    /* null AS middlename, */
    P.CODE AS identifier,
    N.EMAIL_ADDRESS AS email,
    /* null AS sms, */
    /* null AS phone, */
    P.NAME_ID AS agentSourcedIds,
    formYear.DFEE_YEAR_NUMBER AS grades
    /* null AS password */
FROM
    dbo.PUPIL AS P
        INNER JOIN
    dbo.NAME AS N
        ON P.NAME_ID = N.NAME_ID
        INNER JOIN
    dbo.FORM AS form
        ON P.FORM = form.CODE
        INNER JOIN
    dbo.FORM_YEAR AS formYear
        ON form.YEAR_CODE = formYear.CODE
WHERE 
    P.ACADEMIC_YEAR = '2019'
    AND
    form.ACADEMC_YEAR = '2019'
ORDER BY
    P.NAME_ID
/** users - staff **/
SELECT 
    U.NAME_ID AS sourcedId,
    U.IN_USE AS status,
    U.LAST_AMEND_DATE AS dateLastModified,
    U.NAME_ID AS orgSourcedIds,
    'teacher' AS role,
    U.INTERNAL_EMAIL_ADDRESS AS username,
    U.NAME_ID as userIds,
    N.PREFERRED_NAME AS givenName,
    N.SURNAME AS familyname,
    /* null AS middlename, */
    U.CODE AS identifier,
    U.INTERNAL_EMAIL_ADDRESS AS email,
    /* null AS sms, */
    /* null AS phone, */
    U.NAME_ID AS agentSourcedIds
    /* null AS grades, */
    /* null AS password */
FROM
    dbo.STAFF as U
        INNER JOIN
    dbo.NAME AS N
        ON N.NAME_ID = U.NAME_ID
WHERE
    U.CATEGORY = 'TEA001'
    OR
    U.CATEGORY = 'SUPPLY'
    OR
    U.CATEGORY = 'EARLY'
ORDER BY
    U.NAME_ID
/** users - support staff **/
SELECT 
    U.NAME_ID AS sourcedId,
    U.IN_USE AS status,
    U.LAST_AMEND_DATE AS dateLastModified,
    U.NAME_ID AS orgSourcedIds,
    'aide' AS role,
    U.INTERNAL_EMAIL_ADDRESS AS username,
    U.NAME_ID as userIds,
    N.PREFERRED_NAME AS givenName,
    N.SURNAME AS familyname,
    /* null AS middlename, */
    U.CODE AS identifier,
    U.INTERNAL_EMAIL_ADDRESS AS email,
    /* null AS sms, */
    /* null AS phone, */
    U.NAME_ID AS agentSourcedIds
    /* null AS grades, */
    /* null AS password */
FROM
    dbo.STAFF as U
        INNER JOIN
    dbo.NAME AS N
        ON N.NAME_ID = U.NAME_ID
WHERE
    U.CATEGORY = 'NON001'
    OR
    U.CATEGORY = 'COACH'
ORDER BY
    U.NAME_ID
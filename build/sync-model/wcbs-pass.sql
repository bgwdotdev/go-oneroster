/** Orgs **/
SELECT
    SCHOOL_ID AS sourcedId,
    CODE as name,
    'school' AS type,
    DESCRIPTION AS identifier,
    LAST_AMEND_DATE as dateLastModified,
    IN_USE AS status,
FROM 
    dbo.SCHOOL

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
    dbo.SCHOOL_CALENDAR AS SC on SC.ACADEMIC_YEAR = Y.CODE

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

/** classes - scheduled **/ 
SELECT
    S.SUBJECT_SET_ID AS sourcedId,
    S.IN_USE AS status,
    S.LAST_AMEND_DATE AS dateLastModified,
    S.DESCRIPTION AS title,
    /* grades - N/A */
    /* course soucedid */
    S.SET_CODE AS classCode,
    'scheduled' AS classType,
    S.ROOM AS location,
    org.SCHOOL_ID AS schoolSourcedId,
    S.SUBJECT_SET_ID AS termSourcedIds, 
    SUB.DESCRIPTION AS subjects
    /* subjectCodes - SQA codes? */
    /* periods - N/A? */ 
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

/** classes - homeroom(form) **/

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
    dbo.SCHOOL as org
        ON org.CODE = P.SCHOOL
        INNER JOIN
    dbo.NAME AS N
        ON P.NAME_ID = N.NAME_ID
        INNER JOIN
    dbo.FORM AS form
        ON P.FORM = form.CODE
        INNER JOIN
    dbo.FORM_YEAR AS formYear
        ON form.YEAR_CODE = formYear.CODE


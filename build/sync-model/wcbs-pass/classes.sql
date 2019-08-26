/** classes - scheduled **/ 
SELECT
    S.SUBJECT_SET_ID AS sourcedId,
    S.IN_USE AS status,
    S.LAST_AMEND_DATE AS dateLastModified,
    S.DESCRIPTION AS title,
    /* null AS grades, */
    SUB.SUBJECT_ID AS courseSourcedId,
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
    sourcedId
WHERE
    S.ACADEMIC_YEAR = '2019'
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
    sourcedId



/** enrollments - pupils **/
SELECT
    P.PUPIL_SET_ID AS sourcedId,
    SS.IN_USE AS status,
    /* null AS dateLastModified, */
    P.SUBJECT_SET_ID AS classSourcedId,
    org.SCHOOL_ID AS schoolSourcedId,
    PUPIL.NAME_ID AS userSourcedId,
    'student' AS role
    /* null AS primary, */
    /* null AS beginDate, */
    /* null AS endDate */
FROM
    dbo.PUPIL_SET AS P
        INNER JOIN
    dbo.SUBJECT_SET AS SS
        ON SS.SUBJECT_SET_ID = P.SUBJECT_SET_ID
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = SS.SCHOOL
        INNER JOIN
    dbo.PUPIL
        ON PUPIL.PUPIL_ID = P.PUPIL_ID
WHERE
    SS.ACADEMIC_YEAR = '2019'
ORDER BY
    sourcedId
/** enrollments - Teacher 1 **/
DECLARE @T bit
SET @T=1
SELECT
    CONCAT(SS.SUBJECT_SET_ID, S.NAME_ID) AS sourcedId,
    SS.IN_USE AS status,
    /* null AS dateLastModified */
    SS.SUBJECT_SET_ID AS classSourcedId,
    org.SCHOOL_ID as schoolSourcedId,
    S.NAME_ID AS userSourcedId,
    'teacher' AS role,
    @T AS 'primary'
    /* null AS begindate, */
    /* null AS endDate */
FROM
    dbo.SUBJECT_SET AS SS
        INNER JOIN
    dbo.STAFF AS S
        ON SS.TUTOR = S.CODE
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = SS.SCHOOL
WHERE
    SS.ACADEMIC_YEAR = '2019'
ORDER BY
    sourcedId
/** enrollments - Teacher 2 **/
DECLARE @F bit
SET @F=0
SELECT
    CONCAT(SS.SUBJECT_SET_ID, S.NAME_ID) AS sourcedId,
    SS.IN_USE AS status,
    /* null AS dateLastModified */
    SS.SUBJECT_SET_ID AS classSourcedId,
    org.SCHOOL_ID as schoolSourcedId,
    S.NAME_ID AS userSourcedId,
    'teacher' AS role,
    @F AS 'primary'
    /* null AS begindate, */
    /* null AS endDate */
FROM
    dbo.SUBJECT_SET AS SS
        INNER JOIN
    dbo.STAFF AS S
        ON SS.TUTOR_2 = S.CODE
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = SS.SCHOOL
WHERE
    SS.ACADEMIC_YEAR = '2019'
ORDER BY
    sourcedId

/** enrollments - Teacher 3 **/
DECLARE @F bit
SET @F=0
SELECT
    CONCAT(SS.SUBJECT_SET_ID, S.NAME_ID) AS sourcedId,
    SS.IN_USE AS status,
    /* null AS dateLastModified */
    SS.SUBJECT_SET_ID AS classSourcedId,
    org.SCHOOL_ID as schoolSourcedId,
    S.NAME_ID AS userSourcedId,
    'teacher' AS role,
    @F AS 'primary'
    /* null AS begindate, */
    /* null AS endDate */
FROM
    dbo.SUBJECT_SET AS SS
        INNER JOIN
    dbo.STAFF AS S
        ON SS.TUTOR_3 = S.CODE
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = SS.SCHOOL
WHERE
    SS.ACADEMIC_YEAR = '2019'
ORDER BY
    sourcedId


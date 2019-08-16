/** userOrgs - Pupil **/
SELECT
    P.NAME_ID AS userSourcedId,
    org.SCHOOL_ID AS orgSourcedId
FROM
    dbo.PUPIL AS P
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = P.SCHOOL
WHERE
    P.ACADEMIC_YEAR = '2019'
ORDER BY
    orgSourcedId, userSourcedId
/** userOrgs - staff **/
SELECT 
    S.NAME_ID AS userSourcedId,
    org.SCHOOL_ID AS orgSourcedId
FROM
    dbo.STAFF AS S
        INNER JOIN
    dbo.SCHOOL AS org
        ON org.CODE = S.SCHOOL
ORDER BY
    orgSourcedId, userSourcedId



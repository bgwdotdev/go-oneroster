/** classAcademicSessions **/
/* scheduled */
SELECT
    SS.SUBJECT_SET_ID AS classSourcedId,
    Y.YEAR_ID AS academicSessionSourcedId
FROM
    dbo.SUBJECT_SET AS SS
        INNER JOIN
    dbo.YEAR AS Y
        ON Y.CODE = SS.ACADEMIC_YEAR
WHERE
    SS.ACADEMIC_YEAR = '2019'
ORDER BY
    academicSessionSourcedId, classSourcedId

/* homeroom */
SELECT
    F.FORM_ID as classSourcedId,
    Y.YEAR_ID as academicSessionSourcedId
FROM
    dbo.FORM as F
        INNER JOIN
    dbo.YEAR as Y
        ON Y.CODE = F.ACADEMIC_YEAR
WHERE
    F.ACADEMIC_YEAR = '2019'
ORDER BY
    academicSessionSourcedId, classSourcedId

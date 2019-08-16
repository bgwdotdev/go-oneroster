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
    sourcedId



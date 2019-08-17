/** Orgs **/
SELECT
    SCHOOL_ID AS sourcedId,
    CODE as name,
    'school' AS type,
    DESCRIPTION AS identifier,
    LAST_AMEND_DATE AS dateLastModified,
    IN_USE AS status
FROM 
    dbo.SCHOOL
ORDER BY
    sourcedId 

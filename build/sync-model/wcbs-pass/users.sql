/** users - pupils **/
SELECT
    P.NAME_ID AS sourcedId,
    P.IN_USE AS status,
    P.LAST_AMEND_DATE as dateLastModified,
    P.IN_USE AS enabledUser,
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
    formYear.AGE_RANGE AS grades
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
    form.ACADEMIC_YEAR = '2019'
ORDER BY
    sourcedId
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
    sourcedId
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
    sourcedId



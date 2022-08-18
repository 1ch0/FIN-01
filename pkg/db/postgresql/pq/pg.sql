
CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    department character varying(500) NOT NULL,
    Created date,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
    WITH (OIDS=FALSE);

CREATE TABLE userdetail
(
    uid integer,
    intro character varying(100),
    profile character varying(100)
)
    WITH(OIDS=FALSE);
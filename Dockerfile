FROM scratch

COPY ./conf/conf.yaml /conf/conf.yaml
COPY OneRoster.sqlite3 .
COPY credstore .
COPY gorapi-static /

EXPOSE 3000/tcp

ENV GOR_AUTH_KEY='secret'
ENV GOR_AUTH_KEYALG='HS256'
ENV GOR_AUTH_DBDRIVER='sqlite3'
ENV GOR_AUTH_DBNAME='credstore'

CMD ["/gorapi-static"]

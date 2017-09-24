FROM shouldbee/flyway
LABEL maintainer Sergey Nesterov

COPY ./db/migrations /flyway/db/migrations

WORKDIR /flyway
RUN cd db/migrations && ls
CMD ["-url=jdbc:postgresql://postgres/telehorn", "-user=docker", "-password=docker", "migrate", "-locations=filesystem:db/migrations"]

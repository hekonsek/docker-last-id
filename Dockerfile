FROM hekonsek/dkr-docker-go
RUN adduser --disabled-password --gecos "" --force-badname --ingroup docker --uid 500 redhatuser
RUN adduser --disabled-password --gecos "" --force-badname --ingroup docker --uid 1000 nonredhatuser
ADD docker-last-id /usr/bin/
ENTRYPOINT ["/usr/bin/docker-last-id"]
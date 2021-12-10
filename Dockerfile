# -- BUILD STAGE --

# https://hub.docker.com/_/golang/
FROM golang:1.17-buster as development

# input parameter (Makefile target)
ARG make_target

RUN mkdir /app 
# copying everything to the container as this is a BUILD ONLY IMAGE
ADD . /app/
WORKDIR /app 

RUN make $make_target

# -- RELEASE STAGE --

FROM golang:1.17-buster

RUN mkdir /app

WORKDIR /app/

# Copy the executable from the previous stage
RUN echo "copy the executable"
COPY --from=development /app/bin/k8s .

CMD ["./k8s", "-version"]
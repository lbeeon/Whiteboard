FROM alpine:latest

WORKDIR /opt/app

COPY views /opt/app/views
COPY Whiteboard /opt/app

CMD ./Whiteboard web

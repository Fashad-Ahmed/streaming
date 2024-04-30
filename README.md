# Streaming

Live Streaming server that allows you to broadcast live using OBS Studio via the RTMP protocol.

## System Design

![dsd](https://lucid.app/publicSegments/view/c2e21dba-801a-41e6-a0a5-b4aff339d4e9/image.png)

## About

- OBS Studio to carry out the transmission.
- Docker-compose to upload our entire application to make testing easier.
- We will have a database that will be used to store our transmission keys.
- We will have a shared volume between the streaming server containers (nginx) and the playback application because both need to access the same content, that is, while the streaming server is responsible for writing to the database, the playback will read from the same database.
- We will have a user in the browser watching the broadcast provided by the playback application.

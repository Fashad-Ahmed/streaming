# Streaming

Live Streaming server that allows you to broadcast live using OBS Studio via the RTMP protocol.

## System Design

![dsd](./.github/sd%20streaming.svg)

## About

- To carry out the transmission, OBS communicates with our NGINX server via the RTMP protocol in which it uses the authentication service to verify the transmission key passed.
- To expose the transmission to the user, it communicates with the Playback service via the HLS protocol, exposing the sections one after the other.

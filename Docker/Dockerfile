# syntax=docker/dockerfile:1
FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

#install app dependecies
RUN apt-get update && apt-get install -y python3.8

# Set python3.8 as default for python3 command
RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.8 1
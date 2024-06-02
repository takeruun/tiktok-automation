FROM mcr.microsoft.com/playwright/python:v1.44.0-focal

RUN apt-get update && apt-get install -y python3-pip

ENV PYTHONIOENCODING utf-8
ENV TZ="Asia/Tokyo"
ENV LANG=C.UTF-8
ENV LANGUAGE=en_US:en_US

WORKDIR /app

COPY . .

RUN pip install -r ./requirements.txt

RUN python3 -m playwright install
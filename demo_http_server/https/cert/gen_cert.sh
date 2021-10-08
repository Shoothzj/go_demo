#!/bin/bash

openssl req -new -newkey rsa:4096 -days 3650 -nodes -x509 -subj "/C=CN/ST=<State>/L=<City>/O=<Organization>/CN=<Common Name>" -keyout key.pem -out cert.pem -days 365
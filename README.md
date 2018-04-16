[![Build Status](https://travis-ci.org/mlesniak/budget-tracker.svg?branch=master)](https://travis-ci.org/mlesniak/budget-tracker)
[![Go Report Card](https://goreportcard.com/badge/github.com/mlesniak/budget-tracker)](https://goreportcard.com/report/github.com/mlesniak/budget-tracker)

# Overview

A budget tracker.

# Status

In the following we show a screenshot of the current mobile user interface with
fake data (i.e. everything is positive, no expenses, ...).

![Screenshot](current-status.png)

# Technologies

- Backend: Go(lang)
- Frontend: Vuejs + Bootstrap

# Guiding Principles

I will be guided by the following principles

- Open Source everything
- Be open to suggestions
- Write idiomatic go code
- Be pragmatic about testing, but try to test everything
- Do not over-engineer

# HTTPS Support

Since the cookie is (currently) submitted in plaintext, at least a HTTPS connection is mandatory. To
generate the necessary server key, use

    openssl genrsa -out server.key 2048

To generate the corresponding certificate, use

    openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650



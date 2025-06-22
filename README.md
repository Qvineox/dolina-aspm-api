# Domsnail Golang clean template

> Ref: https://github.com/evrone/go-clean-template

## Structure

### api

API file, protobuf files, swagger

### build

Build files, CI/CD pipelines, build configs

### cfg

Runtime configurations, ENV

### cmd

Startup scripts, project main file

### internal

#### app

Main application file, services starter

#### entities

Business logic, data models

#### repos

Data storage layer, outer data transmission to: databases, web apis, cache, etc.

#### servers

API handlers, server

#### services

Functions and data services
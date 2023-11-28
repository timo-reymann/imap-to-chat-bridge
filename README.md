imap-to-chat-bridge
===
[![LICENSE](https://img.shields.io/github/license/timo-reymann/imap-to-chat-bridge)](https://github.com/timo-reymann/imap-to-chat-bridge/blob/main/LICENSE)
[![DockerHub Pulls](https://img.shields.io/docker/pulls/timoreymann/imap-to-chat-bridge)](https://hub.docker.com/r/timoreymann/imap-to-chat-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/timo-reymann/imap-to-chat-bridge)](https://goreportcard.com/report/github.com/timo-reymann/imap-to-chat-bridge)
[![codecov](https://codecov.io/gh/timo-reymann/imap-to-chat-bridge/graph/badge.svg?token=rsQYV5lODS)](https://codecov.io/gh/timo-reymann/imap-to-chat-bridge)
[![CircleCI](https://circleci.com/gh/timo-reymann/imap-to-chat-bridge.svg?style=shield)](https://app.circleci.com/pipelines/github/timo-reymann/imap-to-chat-bridge)
[![GitHub Release](https://img.shields.io/github/v/tag/timo-reymann/imap-to-chat-bridge?label=version)](https://github.com/timo-reymann/imap-to-chat-bridge/releases)
[![Renovate](https://img.shields.io/badge/renovate-enabled-green?logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAzNjkgMzY5Ij48Y2lyY2xlIGN4PSIxODkuOSIgY3k9IjE5MC4yIiByPSIxODQuNSIgZmlsbD0iI2ZmZTQyZSIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTUgLTYpIi8+PHBhdGggZmlsbD0iIzhiYjViNSIgZD0iTTI1MSAyNTZsLTM4LTM4YTE3IDE3IDAgMDEwLTI0bDU2LTU2YzItMiAyLTYgMC03bC0yMC0yMWE1IDUgMCAwMC03IDBsLTEzIDEyLTktOCAxMy0xM2ExNyAxNyAwIDAxMjQgMGwyMSAyMWM3IDcgNyAxNyAwIDI0bC01NiA1N2E1IDUgMCAwMDAgN2wzOCAzOHoiLz48cGF0aCBmaWxsPSIjZDk1NjEyIiBkPSJNMzAwIDI4OGwtOCA4Yy00IDQtMTEgNC0xNiAwbC00Ni00NmMtNS01LTUtMTIgMC0xNmw4LThjNC00IDExLTQgMTUgMGw0NyA0N2M0IDQgNCAxMSAwIDE1eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik04MSAxODVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzI1YzRjMyIgZD0iTTIyMCAxMDBsMjMgMjNjNCA0IDQgMTEgMCAxNkwxNDIgMjQwYy00IDQtMTEgNC0xNSAwbC0yNC0yNGMtNC00LTQtMTEgMC0xNWwxMDEtMTAxYzUtNSAxMi01IDE2IDB6Ii8+PHBhdGggZmlsbD0iIzFkZGVkZCIgZD0iTTk5IDE2N2wxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMDBhZmIzIiBkPSJNMjMwIDExMGwxMyAxM2M0IDQgNCAxMSAwIDE2TDE0MiAyNDBjLTQgNC0xMSA0LTE1IDBsLTEzLTEzYzQgNCAxMSA0IDE1IDBsMTAxLTEwMWM1LTUgNS0xMSAwLTE2eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xMTYgMTQ5bDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxZGRlZGQiIGQ9Ik0xMzQgMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMxYmNmY2UiIGQ9Ik0xNTIgMTEzbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMyNGJmYmUiIGQ9Ik0xNzAgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTYzIDE2N2wxOC0xOCAxOCAxOC0xOCAxOHpNOTggMTMxbDE4LTE4IDE4IDE4LTE4IDE4eiIvPjxwYXRoIGZpbGw9IiMzNGVkZWIiIGQ9Ik0xMzQgOTVsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzFiY2ZjZSIgZD0iTTE1MyA3OGwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjMzRlZGViIiBkPSJNODAgMTEzbDE4LTE3IDE4IDE3LTE4IDE4ek0xMzUgNjBsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzk4ZWRlYiIgZD0iTTI3IDEzMWwxOC0xOCAxOCAxOC0xOCAxOHoiLz48cGF0aCBmaWxsPSIjYjUzZTAyIiBkPSJNMjg1IDI1OGw3IDdjNCA0IDQgMTEgMCAxNWwtOCA4Yy00IDQtMTEgNC0xNiAwbC02LTdjNCA1IDExIDUgMTUgMGw4LTdjNC01IDQtMTIgMC0xNnoiLz48cGF0aCBmaWxsPSIjOThlZGViIiBkPSJNODEgNzhsMTgtMTggMTggMTgtMTggMTh6Ii8+PHBhdGggZmlsbD0iIzAwYTNhMiIgZD0iTTIzNSAxMTVsOCA4YzQgNCA0IDExIDAgMTZMMTQyIDI0MGMtNCA0LTExIDQtMTUgMGwtOS05YzUgNSAxMiA1IDE2IDBsMTAxLTEwMWM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiMzOWQ5ZDgiIGQ9Ik0yMjggMTA4bC04LThjLTQtNS0xMS01LTE2IDBMMTAzIDIwMWMtNCA0LTQgMTEgMCAxNWw4IDhjLTQtNC00LTExIDAtMTVsMTAxLTEwMWM1LTQgMTItNCAxNiAweiIvPjxwYXRoIGZpbGw9IiNhMzM5MDQiIGQ9Ik0yOTEgMjY0bDggOGM0IDQgNCAxMSAwIDE2bC04IDdjLTQgNS0xMSA1LTE1IDBsLTktOGM1IDUgMTIgNSAxNiAwbDgtOGM0LTQgNC0xMSAwLTE1eiIvPjxwYXRoIGZpbGw9IiNlYjZlMmQiIGQ9Ik0yNjAgMjMzbC00LTRjLTYtNi0xNy02LTIzIDAtNyA3LTcgMTcgMCAyNGw0IDRjLTQtNS00LTExIDAtMTZsOC04YzQtNCAxMS00IDE1IDB6Ii8+PHBhdGggZmlsbD0iIzEzYWNiZCIgZD0iTTEzNCAyNDhjLTQgMC04LTItMTEtNWwtMjMtMjNhMTYgMTYgMCAwMTAtMjNMMjAxIDk2YTE2IDE2IDAgMDEyMiAwbDI0IDI0YzYgNiA2IDE2IDAgMjJMMTQ2IDI0M2MtMyAzLTcgNS0xMiA1em03OC0xNDdsLTQgMi0xMDEgMTAxYTYgNiAwIDAwMCA5bDIzIDIzYTYgNiAwIDAwOSAwbDEwMS0xMDFhNiA2IDAgMDAwLTlsLTI0LTIzLTQtMnoiLz48cGF0aCBmaWxsPSIjYmY0NDA0IiBkPSJNMjg0IDMwNGMtNCAwLTgtMS0xMS00bC00Ny00N2MtNi02LTYtMTYgMC0yMmw4LThjNi02IDE2LTYgMjIgMGw0NyA0NmM2IDcgNiAxNyAwIDIzbC04IDhjLTMgMy03IDQtMTEgNHptLTM5LTc2Yy0xIDAtMyAwLTQgMmwtOCA3Yy0yIDMtMiA3IDAgOWw0NyA0N2E2IDYgMCAwMDkgMGw3LThjMy0yIDMtNiAwLTlsLTQ2LTQ2Yy0yLTItMy0yLTUtMnoiLz48L3N2Zz4=)](https://renovatebot.com)
[![pre-commit](https://img.shields.io/badge/%E2%9A%93%20%20pre--commit-enabled-success)](https://pre-commit.com/)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_imap-to-chat-bridge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=timo-reymann_imap-to-chat-bridge)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_imap-to-chat-bridge&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=timo-reymann_imap-to-chat-bridge)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=timo-reymann_imap-to-chat-bridge&metric=bugs)](https://sonarcloud.io/summary/new_code?id=timo-reymann_imap-to-chat-bridge)

<p align="center">
	Bridge to transfer mails from IMAP account to chat apps - timo-reymann/imap-to-chat-bridge
</p>

## Features

- route IMAP messages
- supported chat apps
    - telegram

## Requirements

- IMAP account
- Account for chat app used as target
- Deletes mails from server that have been sent successfully to at least one target

## Installation

- Download [latest release](https://github.com/timo-reymann/imap-to-chat-bridge/releases/latest) and spin up the
  binary as systemd service, screen etc.
- OR use the prebuilt docker images (see [Usage](#usage) for an full example)

## Usage

### Start service with docker-compose

- Create `docker-compose.yml`
  ```yaml
  version: "3.5"
  services:
    imap-to-chat-bridge:
      image: timoreymann/imap-to-chat-bridge:latest # check for version to use if you would like to pin it
      restart: always
      environment:
        IMAP_TO_CHAT_BRIDGE__IMAP_HOST: <host>:<port>
        IMAP_TO_CHAT_BRIDGE__IMAP_TLS_ENABLED: true # Disable if your IMAP server doesnt support TLS
        IMAP_TO_CHAT_BRIDGE__IMAP_USERNAME: <username>
        IMAP_TO_CHAT_BRIDGE__IMAP_PASSWORD: <password>
        IMAP_TO_CHAT_BRIDGE__NOTIFICATION_URIS: shoutrr-itemA,shoutrr-itemB
  ```
- Run `docker compose up -d`

> For all available options just run the binary without any required variable set to get an overview, or if you are
> familiar with go code check the [config sources](./pkg/config/main.go)

### Setup notification targets

For a list of supported targets please check the [shoutrr docs](https://containrrr.dev/shoutrrr/v0.8/services/overview/).

## Motivation

E-Mails for notifications are annoying but sometimes the only messaging solution systems support.

Especially when you don't want to set up a dummy SMTP server or simply don't want to
or it could be not reliable to do so (e.g. from a home server) I find it more comfortable to have a push notification
from the comfort of my chat app.

## Contributing

I love your input! I want to make contributing to this project as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the configuration
- Submitting a fix
- Proposing new features
- Becoming a maintainer

To get started please read the [Contribution Guidelines](./CONTRIBUTING.md).

## Development

### Requirements

- [GNU make](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/get-docker/)
- [pre-commit](https://pre-commit.com/)
- Go 1.21

### Test

```shell
make test
```

### Build

```shell
make build
```

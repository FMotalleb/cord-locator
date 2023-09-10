# Cord-Locator [**UNDER-DEVELOPMENT**]

## Simple Rule-Based DNS Reverse Proxy Project

<div align="center">

![Docker Pulls<Depricated>](https://img.shields.io/docker/pulls/fmotalleb/mockery) ![Docker Pulls](https://img.shields.io/docker/pulls/fmotalleb/cord-locator) [![codecov](https://codecov.io/gh/FMotalleb/cord-locator/branch/main/graph/badge.svg?token=MPZZYK0LUJ)](https://codecov.io/gh/FMotalleb/cord-locator)

[![Publish Container to Docker](https://github.com/FMotalleb/cord-locator/actions/workflows/docker-reg.yml/badge.svg)](https://github.com/FMotalleb/cord-locator/actions/workflows/docker-reg.yml)
[![Publish Container to Github](https://github.com/FMotalleb/cord-locator/actions/workflows/github-reg.yml/badge.svg)](https://github.com/FMotalleb/cord-locator/actions/workflows/github-reg.yml)
[![Publish Container to Github at dev](https://github.com/FMotalleb/cord-locator/actions/workflows/github-reg-dev.yml/badge.svg)](https://github.com/FMotalleb/cord-locator/actions/workflows/github-reg-dev.yml)
[![Tests](https://github.com/FMotalleb/cord-locator/actions/workflows/tests.yml/badge.svg)](https://github.com/FMotalleb/cord-locator/actions/workflows/tests.yml)

</div>

## Deploy

### Using Docker hub

Located at `docker.io/fmotalleb/cord-locator:(version)`

### Using GitHub registry (since v2.0.7)

Located at `ghcr.io/fmotalleb/cord-locator:(version)`

### Development version

Located at `ghcr.io/fmotalleb/cord-locator-dev:(branch name: main,...)`

## Development contracts

* Versioning will follow the semver format.
* Versions in the format of vx.y.* will share the same configuration files without any changes relative to each other.
  * For instance, version 2.0.1 can utilize the configuration file of 2.0.0, though the reverse may not be feasible.
  * The sole exception was for v2.0.7: a keyword in the configuration file was altered (all `resolver` keys are now `resolvers`).

## Progress

### Version 2.0.x [Phase 1] (Done)

In this version, the following tasks have been completed:

* [X] Implemented Raw-Response handling.
* [X] Enabled blocking of specific requests.
* [X] Updated the request-handling flow:
  * Instead of managing requests and composing responses within the provider, requests are now solely managed within the provider, making testing more feasible.
* [X] Implemented DNS provider IP fallback, which replaced the previous random approach (this change was a side effect of altering the flow).
* [X] Modified provider parameters as needed.
* [X] Added support for multiple DNS providers for each rule.
* [X] Test the configuration package.
* [X] Consider renaming the project to a more reasonable name.

 * At this stage, the project involves a rule-based DNS server/DNS reverse proxy capable of manipulating DNS requests.
 * Currently, the server is somewhat static, but in the upcoming phases, this approach will be modified and enhanced.
 * Up to this point, the program has relied on raw DNS resolvers or raw responses. However, the plan is to expand its capabilities by integrating new data sources such as CSV files, SQL databases, Docker integration, Lua integration, and support for DOH (DNS over HTTPS) and DOT (DNS over TLS) servers.

### Version 2.1.x (In Progress)

The upcoming version 2.1.x will introduce the following features:

* [ ] Review and adjust the configuration.
* [ ] Docker Integration.
* [ ] Address and enhance log messages and levels.
* [ ] Address and enhance test messages.
* [ ] Populate the `README.md` file with relevant information.
* [ ] Improve project (code) documentation.
* [ ] Expand test coverage.
* [ ] Lua script support
* [ ] Conduct code refactoring.
* [ ] Implementation of a DOT (DNS over TLS) client. (As Resolver)
* [ ] Development of a DOH (DNS over HTTPS) client. (As Resolver)
* [ ] Introduction of rule grouping for improved organization.
* [ ] Implementation of DNS grouping to enhance management.

### Version 3.0.x (Planned)

The future version 3.0.x is expected to include the following additions:

* [ ] Integration of DOT inward capabilities. (As server)
* [ ] Implementation of DOH inward functionalities. (As server)

Please note that these tasks represent the current project status and planned developments. The project is actively evolving to incorporate these changes and improvements.

## Development

![DevFlow](https://github.com/FMotalleb/cord-locator/assets/30149519/33bfc423-cb14-48b6-ac95-120dfcf946ac)


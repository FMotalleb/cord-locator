# Mockery [**UNDER-DEVELOPMENT**]

[![ci](https://github.com/FMotalleb/mockery/actions/workflows/ci.yml/badge.svg)](https://github.com/FMotalleb/mockery/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/FMotalleb/mockery/branch/main/graph/badge.svg?token=MPZZYK0LUJ)](https://codecov.io/gh/FMotalleb/mockery)


## Simple Rule based DNS reverse proxy
### v2.0.x
* [X] Raw-Response
* [X] Block request
* [X] Update Flow
    > Instead of handling requests and writing responses in the provider
    > Just handle requests in provider dus making it testable
    * [X] DNS provider ip fallback instead of random (Done as a side effect of changing the flow)
* [X] Change provider params
* [X] DNS providers (for each rule)
* [ ] Test config methods
* [ ] Fix test-log messages/levels
* [ ] Fix Docs
* [ ] Add More Tests
* [ ] Refactor the code
* [ ] check config (maybe)
* [ ] Rename Project
* [ ] Fill `README.md`
### v2.1.x
* [ ] DOT client
* [ ] DOH client
* [ ] Rule grouping
* [ ] DNS grouping
### v3.0.x
* [ ] DOT inward
* [ ] DOH inward


## Development

```mermaid
flowchart TB
    subgraph "Initialize"
        A[Main Branch] --> B((Project Start))
    end

    subgraph "Feature Development"
        B --> C[Create Feature Branch: Feature 1]
        B --> D[Create Feature Branch: Feature 2]
        B --> E[Create Feature Branch: Feature 3]
    end

    C --> F((Feature Development))
    D --> F
    E --> F

    F --> G[Create Pull Request]
    G --> H((Actions Fail?))
    H --> |Yes| F
    H --> |NO| J[Merge into Main Branch]
    
    J --> K[Release Version]
    K --> M[Docker Image Build]
```

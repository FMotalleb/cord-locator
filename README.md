# Mockery [**UNDER-DEVELOPMENT**]

[![ci](https://github.com/FMotalleb/mockery/actions/workflows/ci.yml/badge.svg)](https://github.com/FMotalleb/mockery/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/FMotalleb/mockery/branch/main/graph/badge.svg?token=MPZZYK0LUJ)](https://codecov.io/gh/FMotalleb/mockery)


## Simple Rule based dns reverse proxy
### v2.0.x
* [X]     Raw-Response
* [X] block request
* [X] update Flow
    > instead of handling request and writing response in provider
    > just handle request in provider dus making it testable
    * [X] dns provider ip fallback instead of random (Done as a side effect of changing the flow)
* [X] change provider params
* [X] dns providers (for each rule)
* [ ] fill `README.md`
### v2.1.x
* [ ] DOT client
* [ ] DOH client
* [ ] rule grouping
* [ ] dns grouping
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

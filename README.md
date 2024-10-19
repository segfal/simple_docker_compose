# Simple Docker compose project

This is a simple project to test the capabilities of Docker compose. It consists of a Python service, a Go service, and a C++ service. All of them are designed to be run inside Docker containers.

## How to run the project
1. Run docker compose
```docker
docker compose up
```

2. Check the output of the services
```docker
docker compose logs -f
```

3. Stop the services
```docker
docker compose down
```

4. Clean up the project
```docker
docker system prune -a
```

<hr/>

## Visual explanation

```mermaid
graph TD
    subgraph Docker Compose / Kubernetes
        python_service[Python Service - Flask API]
        go_service[Go Service - HTTP Client]
        cpp_service[C++ Service - Factorial Calculator]
    end

    python_service --> |"/api/message"| go_service
    go_service --> cpp_service
    cpp_service --> |"Factorial Result"| go_service

    %% Updated Styles
    style python_service fill:#FFDDC1,stroke:#333,stroke-width:2px,color:#000;
    style go_service fill:#D4E157,stroke:#333,stroke-width:2px,color:#000;
    style cpp_service fill:#4FC3F7,stroke:#333,stroke-width:2px,color:#000;
```


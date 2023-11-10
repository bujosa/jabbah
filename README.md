# Jabbah

Jabbah is a simple project that demonstrates the use of Amazon MemoryDB for Redis.

Amazon MemoryDB for Redis is a fully managed, Redis-compatible, in-memory database that delivers ultra-fast performance and Multi-AZ durability for modern applications built using microservices architectures. MemoryDB is compatible with Redis 5.0.6, enabling you to migrate your highest-performance workloads to the cloud with minimal effort.

## Prerequisites

Before you begin, ensure you have the following installed:

* [Golang](https://golang.org/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)

You also need an [Amazon MemoryDB for Redis](https://aws.amazon.com/memorydb/) account.

## Environment Variables

This project uses the following environment variables:

* `ENVIRONMENT`: The current environment. Can be `development` or `production`. Defaults to `development` if not set.

## Running the Application

### Locally

1. Clone the repository:

    ```bash
    git clone the_repository_name
    ```

2. Set your AWS credentials. See the [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) for more information.

3. Run Redis in a Docker container:

    ```bash
    docker-compose up -d
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

## References

* [Amazon MemoryDB for Redis](https://aws.amazon.com/memorydb/)
* [Amazon MemoryDB for Redis Documentation](https://docs.aws.amazon.com/memorydb/latest/devguide/what-is-memorydb.html)
* [Amazon MemoryDB for Redis Pricing](https://aws.amazon.com/memorydb/pricing/)
* [Golang](https://golang.org/)
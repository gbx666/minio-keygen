🔑 MinIO Keygenerator (.NET)
===

This is a simple [MinIO](https://min.io) Access and Secret Key generator implemented in C#.

# Why?

MinIO stopped auto-generating these secrets, so this little tool generates compatible credentials.

# Security

The generator uses `System.Security.Cryptography.RandomNumberGenerator`, which is cryptographically secure.

# Prerequisites

- .NET 8 SDK

# Usage

```bash
dotnet run --project minio-keygen.csproj
MINIO_ACCESS_KEY=N0SDCYA4JFCBOI9BDAI4APXB
MINIO_SECRET_KEY=MU3QVVSuHHU2fRHaltxZwH2PQwi4TID38br0vVDY1GaqcJQs
```

You can also create an `.env` file from the output:

```bash
dotnet run --project minio-keygen.csproj > .env
```

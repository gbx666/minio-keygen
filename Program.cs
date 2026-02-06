using System;
using System.Buffers;
using System.Security.Cryptography;

const int KeyLength = 18;
const int SecretLength = 36;
const char PaddingCharacter = '+';

if (args.Length > 0 && (args[0] is "-h" or "--help"))
{
    PrintUsage();
    return;
}

string accessKey = GenerateKey(KeyLength, uppercase: true);
string secretKey = GenerateKey(SecretLength, uppercase: false);

Console.WriteLine($"MINIO_ACCESS_KEY={accessKey}");
Console.WriteLine($"MINIO_SECRET_KEY={secretKey}");

static string GenerateKey(int byteLength, bool uppercase)
{
    byte[] data = ArrayPool<byte>.Shared.Rent(byteLength);
    try
    {
        RandomNumberGenerator.Fill(data.AsSpan(0, byteLength));
        string encoded = Convert.ToBase64String(data, 0, byteLength)
            .Replace('+', '-')
            .Replace('/', '_')
            .Replace('=', PaddingCharacter);

        return uppercase ? encoded.ToUpperInvariant() : encoded;
    }
    finally
    {
        Array.Clear(data, 0, byteLength);
        ArrayPool<byte>.Shared.Return(data);
    }
}

static void PrintUsage()
{
    Console.WriteLine("minio-keygen (.NET)");
    Console.WriteLine("Generates a MinIO access key and secret. You can redirect the output into an .env file.");
}

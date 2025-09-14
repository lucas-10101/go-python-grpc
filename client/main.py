import grpc
import uuid_pb2_grpc
import uuid_pb2

def main():

    credentials = grpc.ssl_channel_credentials(open("server.crt", "rb").read())
    with grpc.secure_channel("localhost:50051", credentials) as channel:
        stub = uuid_pb2_grpc.UUIDServiceStub(channel)
        print("Response received (v4):", stub.GetUUID(uuid_pb2.UUIDRequest(version=4)))
        print("Response received (v5):", stub.GetUUID(uuid_pb2.UUIDRequest(version=5, namespace="b903f47e-7d15-40f9-b217-a9ced690f955", valueToHash="hello world")))
        print("Response received (v6):", stub.GetUUID(uuid_pb2.UUIDRequest(version=6)))
        print("Response received (v7):", stub.GetUUID(uuid_pb2.UUIDRequest(version=7)))

if __name__ == "__main__":
    main()
import grpc
import uuid_pb2_grpc
import uuid_pb2

def main():

    credentials = grpc.ssl_channel_credentials(open("server.crt", "rb").read())
    with grpc.secure_channel("localhost:50051", credentials) as channel:
        stub = uuid_pb2_grpc.UUIDServiceStub(channel)
        print("Response received:", stub.GetUUID(uuid_pb2.UUIDRequest(version=4)))

if __name__ == "__main__":
    main()
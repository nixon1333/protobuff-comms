import users_pb2
from google.protobuf.json_format import MessageToJson
import random


class UserHandler:

    def __init__(self):
        self.__users = users_pb2.UsersList()
        self.__user_size = 10
        self.__seed_users()

    def __seed_users(self):
        for i in range(self.__user_size):
            user = users_pb2.User()
            user.id = i
            user.name = f"nixon {i}"
            user.phone = f"01672323 {i}"
            self.__users.users.append(user)

    def get_single_user(self, to_string=False):
        single_user = self.__users.users[random.randint(0, self.__user_size)]
        if to_string:
            return single_user.SerializeToString()
        return single_user

    def add_user(self, request_data):
        user = users_pb2.User()
        user.ParseFromString(request_data)
        self.__users.users.append(user)

    def get_all_users(self, to_string=False, to_json=False):
        if to_string:
            return self.__users.SerializeToString()
        elif to_json:
            return MessageToJson(self.__users)
        return self.__users


from flask import Flask, send_file, jsonify, request
import io

from user_handle import UserHandler

app = Flask(__name__)
file_size = 10

user = UserHandler()


@app.route('/proto', methods=['GET'])
def get_proto_users():
    return send_file(
        io.BytesIO(user.get_all_users(to_string=True)),
        mimetype='attachment/x-protobuf'
    )


@app.route('/proto-solo', methods=['GET'])
def get_proto_solo_user():
    return send_file(
        io.BytesIO(user.get_single_user(to_string=True)),
        mimetype='attachment/x-protobuf'
    )


@app.route('/proto_create', methods=['POST'])
def create_proto_users():
    user.add_user(request.data)
    return "", 204


@app.route('/json', methods=['GET'])
def get_json_users():
    return jsonify({
        "api_stuff": user.get_all_users(to_json=True),
    }), 200

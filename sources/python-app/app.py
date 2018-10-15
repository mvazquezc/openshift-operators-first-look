#!/usr/bin/python

from flask import Flask
from flask import jsonify
from flask import request
import socket

app = Flask(__name__)

version=0.1

@app.route('/version', methods=['GET'])
def show_version():
    """Endpoint to return app version"""
    response = "App Version: " + str(version)
    return response

@app.route('/', methods=['GET'])
def root():
    """Endpoint that returns a hello world message"""
    hostname = socket.gethostname()
    response = "Hello World from: " + str(hostname)
    return response


if __name__ == "__main__":
    app.run(host='0.0.0.0')

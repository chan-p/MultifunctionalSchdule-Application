#!/usr/bin/env python
# -*- coding: utf-8 -*-
from flask import Flask
from jinja2 import Environment, FileSystemLoader
from flask import render_template
import commands
import os

app = Flask(__name__)
app.config['DEBUG'] = True

@app.route('/pull',methods=['POST'])
def asd():
    print(os.system("git status"))
    print(os.system("sh tool/test.sh"))
    print(os.system("./application"))
    a = "HELLO"
    return a

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=1000,threaded=True)

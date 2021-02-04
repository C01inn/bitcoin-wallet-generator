from flask import Flask, request, jsonify
from flask_cors import CORS
import time
import mysql.connector

# create app
app = Flask(__name__)
CORS(app)

dbConn = mysql.connector.connect(
    host="localhost",
    port=1800,
    database="bitcoinWallets",
    user='root',
    password="8268Wrenfield"
)

@app.route('/')
def index():
    return '😄'

@app.route("/new", methods=["POST"])
def newWallet():
    privateKey = request.json.get("privateKey")
    publicKey = request.json.get("publicKey")

    # insert into database
    cur = dbConn.cursor()
    sql = "INSERT INTO wallets (publicKey, privateKey, createdAt) VALUES (%s, %s, %s)"
    sqlParams = (publicKey, privateKey, int(time.time()))
    cur.execute(sql, sqlParams)
    dbConn.commit()

    # done
    return ''' done '''

if __name__ == "__main__":
    app.run(debug=True, port=3737)
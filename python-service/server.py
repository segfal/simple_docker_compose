from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

# Add the /api/message route
@app.route('/api/message', methods=['GET'])
def get_message():
    return jsonify({'message': 'Hello from Python service!'})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
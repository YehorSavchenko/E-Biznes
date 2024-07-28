from flask import Flask, request, jsonify
import requests
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

OLLAMA_URL = 'http://localhost:11434/v1/completions'

@app.route('/chat', methods=['POST'])
def chat():
    user_message = request.json.get('message')
    if not user_message:
        return jsonify({"error": "No message provided"}), 400

    headers = {
        'Content-Type': 'application/json'
    }

    data = {
        "model": "llama3",
        "prompt": user_message
    }

    try:
        response = requests.post(OLLAMA_URL, headers=headers, json=data)
        if response.status_code != 200:
            return jsonify({"error": "Failed to connect to Ollama"}), 500

        response_data = response.json()
        return jsonify({"response": response_data['choices'][0]['text']})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
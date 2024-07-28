from flask import Flask, request, jsonify
import requests
import random
from flask_cors import CORS
from textblob import TextBlob

app = Flask(__name__)
CORS(app)

OLLAMA_URL = 'http://localhost:11434/v1/completions'

openings = [
    "Hello, how can I assist you today?",
    "Hi there, what can I help you with today?",
    "Greetings, how can I be of service today?",
    "Good day, what information are you looking for?",
    "Hi, how can I help you today?"
]

closings = [
    "Thank you for your time. Have a great day!",
    "I'm glad I could help. Goodbye!",
    "It was a pleasure assisting you. Take care!",
    "Thank you for your questions. Have a wonderful day!",
    "Feel free to reach out if you need more help. Goodbye!"
]

allowed_keywords = [
    "clothes", "clothing", "apparel", "shirt", "pants", "dress", "shoes",
    "shop", "store", "purchase", "buy", "return", "exchange"
]

def get_random_opening():
    return random.choice(openings)

def get_random_closing():
    return random.choice(closings)

def is_allowed_topic(message):
    message = message.lower()
    return any(keyword in message for keyword in allowed_keywords)

def analyze_sentiment(text):
    blob = TextBlob(text)
    return blob.sentiment.polarity

@app.route('/chat', methods=['POST'])
def chat():
    user_message = request.json.get('message')
    if not user_message:
        return jsonify({"error": "No message provided"}), 400

    greetings_keywords = ["hello", "hi", "greetings", "good day", "hey"]
    farewells_keywords = ["thank you", "goodbye", "take care", "see you", "have a great day"]

    if any(keyword in user_message.lower() for keyword in greetings_keywords):
        return jsonify({"response": get_random_opening()})

    if any(keyword in user_message.lower() for keyword in farewells_keywords):
        return jsonify({"response": get_random_closing()})

    if not is_allowed_topic(user_message):
        return jsonify({"response": "I'm sorry, but I can only assist with questions related to clothing and our store."}), 400

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

        try:
            response_data = response.json()
            generated_text = response_data['choices'][0]['text']
            sentiment = analyze_sentiment(generated_text)
            if sentiment >= 0:
                return jsonify({"response": generated_text})
            else:
                return jsonify({"response": "The response generated had a negative sentiment. Please try again."})
        except ValueError as e:
            return jsonify({"error": "Invalid JSON response from Ollama", "details": response.text}), 500
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True)
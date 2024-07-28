document.getElementById('send-button').addEventListener('click', sendMessage);

async function sendMessage() {
    const messageInput = document.getElementById('message-input');
    const message = messageInput.value;

    if (!message) return;

    try {
        const response = await fetch('http://localhost:5000/chat', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ message: message }),
        });

        const data = await response.json();

        if (data.error) {
            alert(data.error);
            return;
        }

        const messagesDiv = document.getElementById('messages');
        messagesDiv.innerHTML += `<div class="message user-message">${message}</div>`;
        messagesDiv.innerHTML += `<div class="message bot-response">${data.response}</div>`;

        messageInput.value = '';
    } catch (error) {
        console.error('Error:', error);
        alert('An error occurred. Please try again.');
    }
}
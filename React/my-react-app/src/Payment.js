import React, { useState } from 'react';
import axios from 'axios';

function Payment() {
    const [amount, setAmount] = useState('');
    const [currency, setCurrency] = useState('USD');

    const handleSubmit = (event) => {
        event.preventDefault();
        const paymentData = {
            Amount: parseFloat(amount),
            Currency: currency,
        };
        console.log('Sending payment data:', paymentData);

        axios.post('http://localhost:8080/payment', paymentData)
            .then(response => {
                alert('Payment successful!');
            })
            .catch(error => {
                console.error('Error processing payment:', error.response.data.error);
                alert('Payment failed: ' + error.response.data.error); // Pokaż użytkownikowi powód błędu
            });
    };


    return (
        <div style={{ marginTop: '30px', padding: '20px', border: '1px solid #ccc', borderRadius: '5px' }}>
            <h1>Płatność</h1>
            <form onSubmit={handleSubmit}>
                <div style={{ margin: '10px 0' }}>
                    <label>Kwota:</label>
                    <input type="number" value={amount} onChange={e => setAmount(e.target.value)} />
                </div>
                <div style={{ margin: '10px 0' }}>
                    <label>Waluta:</label>
                    <select value={currency} onChange={e => setCurrency(e.target.value)}>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                </div>
                <button type="submit">Zapłać</button>
            </form>
        </div>
    );
}

export default Payment;
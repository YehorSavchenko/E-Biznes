import React, { useState } from 'react';
import { useCart } from './Context';

function Payment() {
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');
  const { processPayment, paymentStatus } = useCart();

  const handleSubmit = (event) => {
    event.preventDefault();
    processPayment({ Amount: parseFloat(amount), Currency: currency });
  };

  return (
    <div>
      <h1>Płatność</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor='amount'>Kwota:</label>
        <input
          id='amount'
          type='number'
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />
        <label htmlFor='currency'>Waluta:</label>
        <select
          id='currency'
          value={currency}
          onChange={(e) => setCurrency(e.target.value)}
        >
          <option value='USD'>USD</option>
          <option value='EUR'>EUR</option>
        </select>
        <button type='submit'>Zapłać</button>
      </form>
      {paymentStatus && <p>Status płatności: {paymentStatus}</p>}
    </div>
  );
}

export default Payment;

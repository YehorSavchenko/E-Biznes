import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Cart() {
    const [items, setItems] = useState([]);

    const fetchCartItems = () => {
        axios.get('http://localhost:8080/carts/1')
            .then(response => {
                setItems(response.data.Items);
            })
            .catch(error => {
                console.error('Error fetching cart items:', error);
            });
    };

    const removeItemFromCart = (itemId) => {
        axios.delete(`http://localhost:8080/carts/1/items/${itemId}`)
            .then(() => {
                setItems(currentItems => currentItems.filter(item => item.ID !== itemId));
            })
            .catch(error => {
                console.error('Error removing item from cart:', error);
            });
    };

    useEffect(() => {
        fetchCartItems();
    }, []);

    return (
        <div>
            <h1>Koszyk</h1>
            {items.length === 0 ? (
                <p>Koszyk jest pusty</p>
            ) : (
                <ul>
                    {items.map(item => (
                        <li key={item.ID}>
                            {item.Product.Name} - {item.Quantity} szt.
                            <button onClick={() => removeItemFromCart(item.ID)}>Usuń</button>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}

export default Cart;
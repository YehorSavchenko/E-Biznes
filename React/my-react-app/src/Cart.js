import React from 'react';
import { useCart } from './Context';

function Cart() {
    const { items, removeItemFromCart } = useCart();

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
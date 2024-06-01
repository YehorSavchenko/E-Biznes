import React, { useState } from 'react';
import { useCart } from './Context';

function Products() {
  const { products, addItemToCart } = useCart();
  const [quantity, setQuantity] = useState(1);

  return (
    <div>
      <h1>Produkty</h1>
      {products.map((product) => (
        <div
          key={product.ID}
          style={{ margin: '20px', padding: '10px', border: '1px solid #ccc' }}
        >
          <h2>{product.Name}</h2>
          <p>
            <strong>Opis:</strong> {product.Description}
          </p>
          <p>
            <strong>Cena:</strong> {product.Price} zł
          </p>
          <input
            type='number'
            value={quantity}
            onChange={(e) => setQuantity(Math.max(1, parseInt(e.target.value)))}
            min='1'
          />
          <button onClick={() => addItemToCart(product, quantity)}>
            Dodaj do koszyka
          </button>
        </div>
      ))}
    </div>
  );
}

export default Products;

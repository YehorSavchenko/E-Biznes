import React, { useEffect, useState } from 'react';
import axios from 'axios';

function Products() {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/products')
            .then(response => {
                console.log("Received data:", response.data);
                setProducts(response.data);
            })
            .catch(error => {
                console.error('Error fetching products:', error);
            });
    }, []);

    return (
        <div>
            <h1>Produkty</h1>
            {products.map(product => (
                <div key={product.ID} style={{ margin: '20px', padding: '10px', border: '1px solid #ccc' }}>
                    <h2>{product.Name}</h2>
                    <p><strong>Opis:</strong> {product.Description}</p>
                    <p><strong>Cena:</strong> {product.Price} zł</p>
                    <p><strong>Kategoria:</strong> {product.Category.Name}</p>
                </div>
            ))}
        </div>
    );
}

export default Products;
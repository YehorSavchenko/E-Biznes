import React, { createContext, useState, useContext, useEffect } from 'react';
import axios from 'axios';

const Context = createContext();

export const useCart = () => useContext(Context);

export const Provider = ({ children }) => {
    const [items, setItems] = useState([]);
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [paymentStatus, setPaymentStatus] = useState(null);

    const fetchProducts = () => {
        axios.get('http://localhost:8080/products')
            .then(response => {
                setProducts(response.data);
                setLoading(false);
            })
            .catch(error => {
                console.error('Error fetching products:', error);
                setLoading(false);
            });
    };

    const fetchCartItems = () => {
        axios.get('http://localhost:8080/carts/1')
            .then(response => {
                setItems(response.data.Items);
                setLoading(false);
            })
            .catch(error => {
                console.error('Error fetching cart items:', error);
                setLoading(false);
            });
    };

    const addItemToCart = (product, quantity = 1) => {
        const item = {
            ProductID: product.ID,
            Quantity: quantity
        };

        axios.post(`http://localhost:8080/carts/1/items`, item)
            .then(() => {
                fetchCartItems();
            })
            .catch(error => {
                console.error('Error adding item to cart:', error);
            });
    };

    const removeItemFromCart = (itemId) => {
        axios.delete(`http://localhost:8080/carts/1/items/${itemId}`)
            .then(() => {
                fetchCartItems();
            })
            .catch(error => {
                console.error('Error removing item from cart:', error);
            });
    };

    const processPayment = (paymentData) => {
        axios.post('http://localhost:8080/payment', paymentData)
            .then(response => {
                setPaymentStatus('Success');
            })
            .catch(error => {
                console.error('Error processing payment:', error);
                setPaymentStatus('Failed');
            });
    };

    useEffect(() => {
        fetchProducts();
        fetchCartItems();
    }, []);

    const value = {
        items,
        products,
        addItemToCart,
        removeItemFromCart,
        processPayment,
        paymentStatus,
        loading
    };

    return (
        <Context.Provider value={value}>
            {children}
        </Context.Provider>
    );
};
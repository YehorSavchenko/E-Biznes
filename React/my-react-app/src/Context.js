import React, {
  createContext,
  useState,
  useContext,
  useEffect,
  useMemo,
} from 'react';
import PropTypes from 'prop-types';
import api from './Api';

const Context = createContext();

export const useCart = () => useContext(Context);

export const Provider = ({ children }) => {
  const [items, setItems] = useState([]);
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [paymentStatus, setPaymentStatus] = useState(null);

  const fetchProducts = () => {
    api
      .get('http://localhost:8080/products')
      .then((response) => {
        setProducts(response.data);
        setLoading(false);
      })
      .catch((error) => {
        console.error('Error fetching products:', error);
        setLoading(false);
      });
  };

  const fetchCartItems = () => {
    api
      .get('http://localhost:8080/carts/1')
      .then((response) => {
        setItems(response.data.Items);
        setLoading(false);
      })
      .catch((error) => {
        console.error('Error fetching cart items:', error);
        setLoading(false);
      });
  };

  const addItemToCart = (product, quantity = 1) => {
    const item = {
      ProductID: product.ID,
      Quantity: quantity,
    };

    api
      .post(`http://localhost:8080/carts/1/items`, item)
      .then(() => {
        fetchCartItems();
      })
      .catch((error) => {
        console.error('Error adding item to cart:', error);
      });
  };

  const removeItemFromCart = (itemId) => {
    api
      .delete(`http://localhost:8080/carts/1/items/${itemId}`)
      .then(() => {
        fetchCartItems();
      })
      .catch((error) => {
        console.error('Error removing item from cart:', error);
      });
  };

  const processPayment = (paymentData) => {
    api
      .post('http://localhost:8080/payment', paymentData)
      .then(() => {
        setPaymentStatus('Success');
      })
      .catch((error) => {
        console.error('Error processing payment:', error);
        setPaymentStatus('Failed');
      });
  };

  useEffect(() => {
    fetchProducts();
    fetchCartItems();
  }, []);

  const value = useMemo(
    () => ({
      items,
      products,
      addItemToCart,
      removeItemFromCart,
      processPayment,
      paymentStatus,
      loading,
    }),
    [items, products, paymentStatus, loading],
  );

  return <Context.Provider value={value}>{children}</Context.Provider>;
};

Provider.propTypes = {
  children: PropTypes.node.isRequired,
};

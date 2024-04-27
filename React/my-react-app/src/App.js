import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { Provider } from './Context';
import Products from './Products';
import Payment from './Payment';
import Cart from './Cart';

function App() {
    return (
        <Provider>
            <Router>
                <div className="App">
                    <nav>
                        <Link to="/products">Produkty</Link> |
                        <Link to="/cart">Koszyk</Link> |
                        <Link to="/payment">Płatność</Link>
                    </nav>
                    <Routes>
                        <Route path="/products" element={<Products />} />
                        <Route path="/cart" element={<Cart />} />
                        <Route path="/payment" element={<Payment />} />
                        <Route path="/" element={<Products />} />
                    </Routes>
                </div>
            </Router>
        </Provider>
    );
}

export default App;
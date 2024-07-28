import React, { useState } from 'react';
import axios from 'axios';

function App() {
    // Stany dla rejestracji
    const [registerUsername, setRegisterUsername] = useState('');
    const [registerPassword, setRegisterPassword] = useState('');

    // Stany dla logowania
    const [loginUsername, setLoginUsername] = useState('');
    const [loginPassword, setLoginPassword] = useState('');

    const register = async () => {
        try {
            await axios.post('http://localhost:5000/register', { username: registerUsername, password: registerPassword });
            alert('User registered successfully');
            setRegisterUsername('');
            setRegisterPassword('');
        } catch (error) {
            alert('Error registering user');
        }
    };

    const login = async () => {
        try {
            const response = await axios.post('http://localhost:5000/login', { username: loginUsername, password: loginPassword });
            const token = response.data.token;
            // Zapisanie tokenu w localStorage lub w stanie aplikacji
            localStorage.setItem('token', token);
            alert('Logged in successfully');
            setLoginUsername('');
            setLoginPassword('');
        } catch (error) {
            alert('Invalid credentials');
        }
    };

    const loginWithGoogle = () => {
        window.location.href = 'http://localhost:5000/auth/google';
    };

    const loginWithGitHub = () => {
        window.location.href = 'http://localhost:5000/auth/github';
    };

    return (
        <div>
            <h1>Register</h1>
            <input
                type="text"
                placeholder="Username"
                value={registerUsername}
                onChange={(e) => setRegisterUsername(e.target.value)}
            />
            <input
                type="password"
                placeholder="Password"
                value={registerPassword}
                onChange={(e) => setRegisterPassword(e.target.value)}
            />
            <button onClick={register}>Register</button>

            <h1>Login</h1>
            <input
                type="text"
                placeholder="Username"
                value={loginUsername}
                onChange={(e) => setLoginUsername(e.target.value)}
            />
            <input
                type="password"
                placeholder="Password"
                value={loginPassword}
                onChange={(e) => setLoginPassword(e.target.value)}
            />
            <button onClick={login}>Login</button>

            <h1>Or</h1>
            <button onClick={loginWithGoogle}>Login with Google</button>
            <button onClick={loginWithGitHub}>Login with GitHub</button>
        </div>
    );
}

export default App;
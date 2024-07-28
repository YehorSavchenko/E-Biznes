require('dotenv').config();
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const sqlite3 = require('sqlite3').verbose();
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');

const app = express();
app.use(bodyParser.json());
app.use(cors());

const PORT = 5000;
const JWT_SECRET = process.env.JWT_SECRET;

// Initialize SQLite database
const db = new sqlite3.Database(':memory:');

db.serialize(() => {
    db.run("CREATE TABLE users (id INTEGER PRIMARY KEY, username TEXT UNIQUE, password TEXT, token TEXT)");
});

// Registration endpoint
app.post('/register', (req, res) => {
    const { username, password } = req.body;

    bcrypt.hash(password, 10, (err, hashedPassword) => {
        if (err) return res.status(500).send('Error hashing password');

        const stmt = db.prepare("INSERT INTO users (username, password) VALUES (?, ?)");
        stmt.run(username, hashedPassword, (err) => {
            if (err) return res.status(400).send('Error registering user');
            res.status(201).send('User registered successfully');
        });
        stmt.finalize();
    });
});

// Login endpoint
app.post('/login', (req, res) => {
    const { username, password } = req.body;

    db.get("SELECT * FROM users WHERE username = ?", [username], (err, user) => {
        if (err) return res.status(500).send('Error querying database');
        if (!user) return res.status(400).send('Invalid credentials');

        bcrypt.compare(password, user.password, (err, isPasswordValid) => {
            if (err) return res.status(500).send('Error comparing passwords');
            if (!isPasswordValid) return res.status(400).send('Invalid credentials');

            const token = jwt.sign({ userId: user.id, username: user.username }, JWT_SECRET, { expiresIn: '1h' });

            db.run("UPDATE users SET token = ? WHERE id = ?", [token, user.id], (err) => {
                if (err) return res.status(500).send('Error updating token');
                res.status(200).json({ token });
            });
        });
    });
});

// Endpoint do sprawdzania danych użytkownika (do testów)
app.get('/users', (req, res) => {
    db.all("SELECT id, username, password, token FROM users", [], (err, rows) => {
        if (err) return res.status(500).send('Error querying database');
        res.status(200).json(rows);
    });
});

app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});
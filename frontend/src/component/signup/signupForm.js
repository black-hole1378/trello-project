"use client";

import React, { useState } from 'react';
import { Box, Button, TextField, Typography, Link } from '@mui/material';
import axios from 'axios';
import { useRouter } from 'next/navigation';

const SignUpForm = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [retypePassword, setRetypePassword] = useState('');
    const [error, setError] = useState(null);
    const router = useRouter()

    const handleSubmit = async (event) => {
        event.preventDefault();

        const passwordRegex = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*(),.?":{}|<>]).{8,}$/;

        if (password !== retypePassword) {
            setError('Passwords do not match.');
            return;
        }

        if (!passwordRegex.test(password)) {
            setError('Password must be at least 8 characters long and contain at least one number, one uppercase letter, one lowercase letter, and one special character.');
            return;
        }

        try {
            const response = await axios.post('http://localhost:8000/auth/signup', {
                username,
                email,
                password,
            });
            const { accessToken, refreshToken, message, user } = response.data;
            // Store tokens in localStorage or sessionStorage for further usage
            localStorage.setItem('accessToken', accessToken);
            localStorage.setItem('refreshToken', refreshToken);
            localStorage.setItem('userID', user.ID);
            router.push("/");
            console.log(response.data);
        } catch (error) {
            setError(error.response ? error.response.data.message : error.message);
        }
    };

    return (
        <Box
            component="form"
            onSubmit={handleSubmit}
        >
            <Typography component="h1" textAlign={"center"} variant="h5">
                Sign Up
            </Typography>
            <TextField
                margin="normal"
                required
                fullWidth
                id="username"
                label="Username"
                name="username"
                autoComplete="username"
                autoFocus
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="new-password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                name="retype-password"
                label="Retype Password"
                type="password"
                id="retype-password"
                autoComplete="new-password"
                value={retypePassword}
                onChange={(e) => setRetypePassword(e.target.value)}
            />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
            >
                Create Account
            </Button>
            {error && <Typography variant="body2" color="error">{error}</Typography>}
            <Link href="/login" variant="body2" style={{ textDecoration: "none", outline: "none" }}>
                <Typography color={"text.info"} textAlign={"center"}>
                    Already have an account? Log in
                </Typography>
            </Link>
        </Box>
    );
};

export default SignUpForm;

// components/LoginForm.js
"use client"
import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Box, Button, TextField, Typography } from '@mui/material';
import axios from 'axios';

const LoginForm = () => {
    const [userName, setUserName] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const router = useRouter(); // Get router object from Next.js


    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsLoading(true);
        setError(null);

        try {
            const response = await axios.post('http://localhost:8000/auth/login', {
                userName,
                password,
            });
            const { accessToken, refreshToken, userID } = response.data;
            // Store tokens in localStorage or sessionStorage for further usage
            localStorage.setItem('accessToken', accessToken);
            localStorage.setItem('refreshToken', refreshToken);
            localStorage.setItem('userID', userID);
            router.push("/")
        } catch (err) {
            setError(err.response ? err.response.data.message : err.message);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <Box
            component="form"
            onSubmit={handleSubmit}
        >
            <Typography component="h1" variant="h5">
                Log in to continue
            </Typography>
            <TextField
                margin="normal"
                required
                fullWidth
                id="userName"
                label="Enter your username"
                name="userName"
                autoFocus
                value={userName}
                onChange={(e) => setUserName(e.target.value)}
            />
            <TextField
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                disabled={isLoading}
            >
                {isLoading ? 'Loading...' : 'Continue'}
            </Button>
            {error && <Typography variant="body1" textAlign={"center"} color="error">{error}</Typography>}
        </Box>
    );
};

export default LoginForm;

"use client"
import React, { useState } from 'react';
import { Box, Button, TextField, Typography } from '@mui/material';
import useFetch from '../../useFetch/useFetch';

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [submitUrl, setSubmitUrl] = useState(null);
    const [requestOptions, setRequestOptions] = useState(undefined);

    const { data, isLoading, error } = useFetch(submitUrl, requestOptions);

    const handleSubmit = (event) => {
        event.preventDefault();
        setSubmitUrl('/api/login');
        setRequestOptions({
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            data: JSON.stringify({ email, password }),
        });
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
                id="email"
                label="Enter your email"
                name="email"
                autoComplete="email"
                autoFocus
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
                autoComplete="current-password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
            >
                Continue
            </Button>
            {isLoading && <Typography variant="body2">Loading...</Typography>}
            {error && <Typography variant="body2" color="error">{error.message}</Typography>}
            {data && <Typography variant="body2">Login Successful!</Typography>}
        </Box>
    );
};

export default LoginForm;

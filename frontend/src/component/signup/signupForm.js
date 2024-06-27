"use client";

import React, { useState, useEffect } from 'react';
import { Box, Button, TextField, Typography, Link } from '@mui/material';
import useFetch from '../../useFetch/useFetch';

const SignUpForm = () => {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [retypePassword, setRetypePassword] = useState('');
    const [submit, setSubmit] = useState(false);
    const [error, setError] = useState(null);
    const [success, setSuccess] = useState(null);

    const { data, error: fetchError, isLoading } = useFetch(
        submit ? '/api/signup' : null,
        submit ? {
            method: 'POST',
            data: { username, email, password },
        } : null
    );

    useEffect(() => {
        if (fetchError) {
            setError(fetchError.response ? fetchError.response.data.message : fetchError.message);
        }
        if (data) {
            setSuccess('Account created successfully!');
            resetForm();
        }
        setSubmit(false);
    }, [data, fetchError]);

    const handleSubmit = (event) => {
        event.preventDefault();
        if (password !== retypePassword) {
            setError('Passwords do not match.');
            return;
        }
        setSubmit(true);
        setError(null);
        setSuccess(null);
    };

    const resetForm = () => {
        setUsername('');
        setEmail('');
        setPassword('');
        setRetypePassword('');
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
                disabled={isLoading}
            >
                {isLoading ? 'Loading...' : 'Create Account'}
            </Button>
            {error && <Typography variant="body2" color="error">{error}</Typography>}
            {success && <Typography variant="body2">{success}</Typography>}
            <Link href="/login" variant="body2" style={{ textDecoration: "none", outline: "none" }}>
                <Typography color={"text.info"} textAlign={"center"}>
                    Already have an account? Log in
                </Typography>
            </Link>
        </Box>
    );
};

export default SignUpForm;

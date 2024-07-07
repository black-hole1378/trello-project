// components/Dashboard.js
"use client";

import React from 'react';
import { Container, Typography, List, ListItem, ListItemText, Divider } from '@mui/material';
import { useRouter } from 'next/navigation';

const Dashboard = () => {
    const router = useRouter();

    const handleNavigation = (path) => {
        router.push(path);
    };

    return (
        <Container maxWidth="lg" sx={{ mt: 4 }}>
            <Typography variant="h4" gutterBottom>
                Dashboard
            </Typography>
            <Divider sx={{ mb: 2 }} />
            <List>
                <ListItem button onClick={() => handleNavigation('/')}>
                    <ListItemText primary="Boards" />
                </ListItem>
                <ListItem button onClick={() => handleNavigation('/dashboard')}>
                    <ListItemText primary="Tasks" />
                </ListItem>
            </List>
        </Container>
    );
};

export default Dashboard;

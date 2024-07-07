// app/page.js
"use client";

import React, { useState, useEffect } from 'react';
import Dashboard from '../component/dashboard/dashboard';
import { Container, Grid, Card, CardContent, Typography, Box } from '@mui/material';
import axios from 'axios';

const HomePage = () => {
    const [workSpaces, setWorkSpaces] = useState([]);

    useEffect(() => {
        const fetchWorkSpaces = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/workspaces`, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });
                console.log("console12", response.data)
                setWorkSpaces(response.data.workSpaces);
            } catch (error) {
                console.error("Error fetching workspaces:", error);
            }
        };

        fetchWorkSpaces();
    }, []);

    return (
        <Container sx={{ mt: 2 }}>
            <Grid container spacing={2}>
                <Grid item md={3} bgcolor={"#F5F7F8"} height={"100vh"}>
                    <Dashboard />
                </Grid>
                <Grid item md={9}>
                    <Box display="flex" flexWrap="wrap" gap={2}>
                        {workSpaces.length ? workSpaces.map((workSpace) => (
                            <Card key={workSpace.ID} sx={{ width: '30%', minWidth: '250px' }}>
                                <CardContent>
                                    <Typography variant="h6" component="div">
                                        {workSpace.Name}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        {workSpace.Description}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                                        Created at: {new Date(workSpace.CreatedAt).toLocaleDateString()}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Updated at: {new Date(workSpace.UpdatedAt).toLocaleDateString()}
                                    </Typography>
                                </CardContent>
                            </Card>
                        )) : <div></div>}
                    </Box>
                </Grid>
            </Grid>
        </Container>
    );
};

export default HomePage;

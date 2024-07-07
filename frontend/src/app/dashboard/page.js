// app/page.js
"use client";

import React, { useState, useEffect } from 'react';
import Dashboard from '../../component/dashboard/dashboard';
import { Container, Grid, Box } from '@mui/material';
import axios from 'axios';
import Task from '@/component/dashboard/Task';

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
                        {workSpaces.length ? workSpaces.map((workSpace, index) => (
                            <Task key={index} workSpace={workSpace} />
                        )) : <div></div>}
                    </Box>
                </Grid>
            </Grid>
        </Container>
    );
};

export default HomePage;

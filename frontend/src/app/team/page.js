"use client"
import React, { useState, useEffect } from "react";
import axios from "axios";
import { Box, Paper, Container, List, ListItem, ListItemText, Typography } from "@mui/material";

export default function Team() {
    const [users, setUsers] = useState([]);
    const workSpaceID = localStorage.getItem("workSpaceID");

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/workspaces/${workSpaceID}/users`);
                setUsers(response.data.users);
                console.log(response.data)
            } catch (error) {
                console.error("Error fetching users:", error);
            }
        };

        if (workSpaceID) {
            fetchUsers();
        }
    }, [workSpaceID]);

    return (
        <Container maxWidth="sm">
            <Box mt={4} p={3} component={Paper} elevation={3}>
                <Typography variant="h5" gutterBottom>
                    Team Members
                </Typography>
                <List>
                    {users.map(user => (
                        <ListItem key={user.ID} divider>
                            <ListItemText primary={user.Username} secondary={user.Email} />
                        </ListItem>
                    ))}
                </List>
            </Box>
        </Container>
    );
}

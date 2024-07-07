"use client"
import React, { useState } from "react";
import { Container, TextField, Button, Typography, Box } from "@mui/material";
import axios from "axios";

export default function SubTaskForm() {
    const [userName, setUserName] = useState("");
    const [title, setTitle] = useState("");
    const [isComplete, setIsComplete] = useState("");
    const [responseMessage, setResponseMessage] = useState("");
    const taskID = localStorage.getItem("taskID");

    const handleSubmit = async (event) => {
        event.preventDefault();
        setResponseMessage("");

        try {
            const response = await axios.post(`http://localhost:8000/tasks/${taskID}/subtasks`, {
                userName,
                title,
                status: isComplete,
            }, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
                    'Content-Type': 'application/json',
                },
            }
            );

            if (response.status == 200) {
                setResponseMessage("Subtask created successfully!");
                console.log("Subtask created:", response.data);
            }
            else {
                setResponseMessage("An error Occured")
            }

        } catch (error) {
            setResponseMessage("Error creating subtask!");
            console.error("Error creating subtask:", error);
        }
    };

    return (
        <Container maxWidth="sm">
            <Box mt={5}>
                <Typography variant="h4" gutterBottom>
                    Create SubTask
                </Typography>
                <form onSubmit={handleSubmit}>
                    <TextField
                        label="User Name"
                        value={userName}
                        onChange={(e) => setUserName(e.target.value)}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <TextField
                        label="Title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <TextField
                        label="Is completed"
                        value={isComplete}
                        onChange={(e) => setIsComplete(e.target.value)}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <Box mt={2}>
                        <Button type="submit" variant="contained" color="primary" fullWidth>
                            Create SubTask
                        </Button>
                    </Box>
                </form>
                {responseMessage && (
                    <Typography variant="body2" color="textSecondary" mt={2}>
                        {responseMessage}
                    </Typography>
                )}
            </Box>
        </Container>
    );
}

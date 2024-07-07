"use client"
import React, { useEffect, useState } from "react";
import { Container, Paper, Box, TextField, Stack, Typography, Button } from "@mui/material";
import { useRouter } from "next/navigation";

export default function AddUser() {
    const [user, setUser] = useState("");
    const [error, setError] = useState("");
    const router = useRouter();

    const handleSubmit = async (event) => {
        event.preventDefault();

        try {
            const response = await fetch(`http://localhost:8000/workspaces/${localStorage.getItem("workSpaceID")}/users`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                },
                body: JSON.stringify({ role: "Standard User", userName: user })
            });
            if (response.ok) {
                alert(user + " Successfully Added to WorkSpace!")
                setUser("")
            }

        } catch (error) {
            setError(error.response ? error.response.data.message : error.message);
        }
    };

    const handleNext = () => {
        router.push(`/board?workSpaceID=${localStorage.getItem("workSpaceID")}`)
    }

    return (
        <Box
            component="form"
            onSubmit={handleSubmit}
            display={"flex"}
            justifyContent={"center"}
            alignItems={"center"} m={2}>
            <Paper sx={{ width: 300, height: 400 }}>
                <Container>
                    <Stack spacing={2}>
                        <Typography variant="h6" textAlign={"center"}>Add User to Board</Typography>
                        <TextField
                            margin="normal"
                            fullWidth
                            size="small"
                            label="User"
                            name="user"
                            autoFocus
                            value={user}
                            onChange={(e) => setUser(e.target.value)}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Add
                        </Button>
                        <Button
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                            onClick={handleNext}
                        >
                            Next
                        </Button>
                    </Stack>
                </Container>
            </Paper>
        </Box>

    );
};


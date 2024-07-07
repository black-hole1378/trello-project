"use client"
import { Container, Paper, Box, TextField, Stack, Typography, Button } from "@mui/material";
import React, { useState } from "react";
import { useRouter } from "next/navigation";

export default function Create() {
    const [name, setName] = React.useState("");
    const [description, setDescription] = React.useState("");
    const [error, setError] = useState(null);
    const router = useRouter()

    const handleSubmit = async (event) => {
        event.preventDefault();

        try {
            const response = await fetch(`http://localhost:8000/workspaces`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                },
                body: JSON.stringify({ name, description })
            });
            if (response.ok) {
                const data = await response.json();
                router.push("/create/user");
                localStorage.setItem("workSpaceID", data.workSpace.ID)
            }

        } catch (error) {
            setError(error.response ? error.response.data.message : error.message);
        }
    };


    return (
        <Box
            component="form"
            onSubmit={handleSubmit}
            display={"flex"}
            justifyContent={"center"}
            alignItems={"center"} m={2}>
            <Paper sx={{ width: 500, height: 500 }}>
                <Container>
                    <Stack spacing={2}>
                        <Typography variant="h6" textAlign={"center"}>New Board</Typography>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            size="small"
                            label="Name"
                            name="name"
                            autoComplete="name"
                            autoFocus
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                        <TextField
                            fullWidth
                            rows={5}
                            multiline
                            label="Description"
                            name="description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Create Board
                        </Button>
                    </Stack>
                </Container>
            </Paper>
        </Box>
    )
}
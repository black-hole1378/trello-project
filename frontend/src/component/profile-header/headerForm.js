import React, { useState } from 'react';
import { TextField, Button, Grid, Box } from '@mui/material';
import { useRouter } from 'next/navigation';

export default function HeaderForm({ user }) {
    console.log("user", user)
    const [username, setUsername] = useState(user.Username);
    const [bio, setBio] = useState("My name is Ahamad");
    const [email, setEmail] = useState(user.Email);
    const [password, setPassword] = useState('');
    const router = useRouter();

    const handleSubmit = async (event) => {
        event.preventDefault();

        const requestBody = {
            username,
            email,
            password
        };

        try {
            const response = await fetch(`http://localhost:8000/users/${localStorage.getItem("userID")}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                },
                body: JSON.stringify(requestBody)
            });

            if (!response.ok) {
                const errorMessage = await response.text();
                throw new Error(`HTTP error! Status: ${response.status}, Message: ${errorMessage}`);
            }

            const data = await response.json();
            localStorage.setItem('accessToken', data.accessToken);
            localStorage.setItem('refreshToken', data.refreshToken);
            alert("Successfully updated!")
            // Handle success response (e.g., show success message to user)
        } catch (error) {
            console.error('Error:', error);
            // Handle error (e.g., show error message to user)
        }
    };

    const onDelete = async (event) => {
        const answer = window.confirm("Are you sure?");
        event.preventDefault()
        if (answer)
            try {
                const response = await fetch(`http://localhost:8000/users/${localStorage.getItem("userID")}`, {
                    method: 'DELETE',
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });

                if (!response.ok) {
                    const errorMessage = await response.text();
                    throw new Error(`HTTP error! Status: ${response.status}, Message: ${errorMessage}`);
                }
                alert("Successfully deleted")
                router.push("/signup")
            } catch (error) {
                console.error('Error:', error);
                // Handle error (e.g., show error message to user)
            }
    }

    return (
        <Box>
            <form onSubmit={handleSubmit}>
                <Grid container spacing={2}>
                    <Grid item xs={12}>
                        <TextField
                            label="Username"
                            fullWidth
                            value={username}
                            onChange={(event) => setUsername(event.target.value)}
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <TextField
                            label="Biography"
                            fullWidth
                            multiline
                            rows={4}
                            value={bio}
                            onChange={(event) => setBio(event.target.value)}
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <TextField
                            label="Email"
                            fullWidth
                            value={email}
                            onChange={(event) => setEmail(event.target.value)}
                        />
                    </Grid>
                    <Grid item xs={12} display={"flex"} justifyContent={"space-around"}>
                        <Button type="submit" variant="contained" color="success">
                            Save
                        </Button>
                        <Button onClick={onDelete} variant="contained" color="error">
                            Delete
                        </Button>
                    </Grid>
                </Grid>
            </form>
        </Box>
    );
};


import React, { useEffect, useState } from "react";
import { Box, Container, Typography, Button, Divider } from "@mui/material";
import axios from "axios";
import GroupWorkIcon from '@mui/icons-material/GroupWork';
import GroupsIcon from '@mui/icons-material/Groups';
import Link from "next/link";

export default function Header() {
    const [workSpace, setWorkSpace] = useState(null);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                console.log("hello", localStorage.getItem("workSpaceID"))
                const response = await axios.get(`http://localhost:8000/workspaces/${localStorage.getItem("workSpaceID")}`, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });
                console.log(response.data)
                setWorkSpace(response.data.workSpace); // Assuming the response structure matches { user: ... }
            } catch (error) {
                setError(error.message);
            }
        };

        fetchData();
    }, []);

    if (workSpace != null)
        return (
            <Box>
                <Container>
                    <Box display={"flex"} alignItems={"center"} gap={3}>
                        <Box display={"flex"} alignItems={"center"} gap={1}>
                            <GroupWorkIcon fontSize="medium" />
                            <Typography variant="body1">{workSpace.Name}</Typography>
                        </Box>
                        <Button component={Link} href="/team" variant="contained" sx={{ bgcolor: 'text.success', p: 1 }}>
                            <GroupsIcon fontSize="small" />
                            Teams
                        </Button>
                        <Button component={Link} href="/" variant="contained" color="error">
                            Exit
                        </Button>
                    </Box>
                    <Divider sx={{ mt: 1, color: "ActiveCaption" }} orientation="horizontal" />
                </Container>
            </Box>
        )
}


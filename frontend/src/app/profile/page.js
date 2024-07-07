"use client"
import { useUser } from "@/component/profile-header/user-cotext";
import HeaderForm from "../../component/profile-header/headerForm";
import { Container, Divider, Paper, Typography } from "@mui/material";
import { Stack, Box } from "@mui/system";
import React, { useState, useEffect } from "react";

const Profile = () => {
    const u = useUser()
    console.log(u.user)
    const [user, setUser] = useState(null);

    if (!u.user) {
        return <div>Error while fetching!</div>
    }
    else {
        return (
            <Container sx={{ display: "flex", justifyContent: "center", mt: 2 }}>
                <Paper sx={{ width: 500, p: 2 }}>
                    <Stack spacing={4}>
                        <Typography variant="body1" fontWeight={"600"}>
                            Manage your personal information
                        </Typography>
                        <Stack spacing={2}>
                            <Stack spacing={0.6}>
                                <Typography variant="body2" fontWeight={"600"}>
                                    About
                                </Typography>
                                <Divider orientation="horizontal" />
                            </Stack>
                            <HeaderForm user={u.user} />
                        </Stack>
                    </Stack>
                </Paper>
            </Container>
        )
    }
}

export default Profile
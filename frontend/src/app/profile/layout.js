"use client";
import { Box, Stack } from "@mui/material";
import React from "react";
import ProfileContainer from "@/component/profile-header/Header";
import withAuth from "../hoc/withAuth";
import { UserProvider } from "../../component/profile-header/user-cotext";

const Layout = ({ children }) => {
    return (
        <UserProvider>
            <Stack spacing={2} component={"main"} m={2}>
                <Box>
                    <ProfileContainer />
                </Box>
                <Box>
                    {children}
                </Box>
            </Stack>
        </UserProvider>
    );
};

export default withAuth(Layout);

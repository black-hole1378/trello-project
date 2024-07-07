"use client";
import { Box, Stack } from "@mui/material";
import React from "react";
import withAuth from "../hoc/withAuth";
import Header from "@/component/board/Header";

const Layout = ({ children }, props) => {
    return (
        <Stack spacing={2} component={"main"} m={2}>
            <Box>
                <Header />
            </Box>
            <Box>
                {children}
            </Box>
        </Stack>
    );
};

export default withAuth(Layout);

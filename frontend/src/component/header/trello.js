import { Box, Typography } from "@mui/material";
import React from "react";
import MenuIcon from '@mui/icons-material/Menu';

export default function Trello() {
    return (
        <Box display={"flex"} flexDirection={"row"} alignItems={"center"} gap={1}>
            <MenuIcon fontSize="medium" />
            <Typography variant="body1">
                Trello
            </Typography>
        </Box>
    )
}
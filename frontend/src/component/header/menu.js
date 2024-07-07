import React from "react";
import { Box, Typography } from "@mui/material";
import Link from "next/link";

export default function Menu({ url, name }) {
    return (
        <Box>
            <Link style={{ textDecoration: "none", outline: "none" }} href={url}>
                <Typography variant='body1' color={"text.primary"}>
                    {name}
                </Typography>
            </Link>
        </Box>
    )
}
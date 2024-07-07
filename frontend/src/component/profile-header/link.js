import React from "react";

import { Divider, Stack, Typography } from "@mui/material";

import Link from "next/link";

const MyLink = ({ url, name, selected, index, changeSelected }) => {
    return (
        <Stack>
            <Link onClick={() => changeSelected(index)} style={{ textDecoration: "none", outline: "none" }} href={url}>
                <Typography variant='body1' color={selected ? "text.info" : "text.primary"}>
                    {name}
                </Typography>
            </Link>
            <Divider sx={{ display: selected ? "flex" : "none", color: "blue" }} />
        </Stack>
    )
}

export default MyLink
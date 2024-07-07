import { Avatar, Box, Divider, Stack, Typography } from "@mui/material";
import React from "react";
import Link from "./link";
import { useUser } from "./user-cotext";

const datas = [
    {
        name: "Profile and Visibility",
        url: "/profile",
        selected: true
    },
    {
        name: "Activity",
        url: "/profile/activity",
        selected: false
    },
    {
        name: "Cards",
        url: "/profile/cards",
        selected: false
    },
    {
        name: "Setting",
        url: "/profile/setting",
        selected: false
    },
]

export default function ProfileContainer() {
    const [data, setData] = React.useState(datas)

    const changeSelected = (index) => {
        let new_data = data
        for (let i = 0; i < new_data.length; i++) {
            new_data[i].selected = false
        }
        new_data[index].selected = true
        setData(new_data)
    }

    const u = useUser()

    if (!u.user)
        return <div>Loading...</div>
    else
        return (
            <Stack spacing={2}>
                <Box display={"flex"} flexDirection={"row"} alignItems={"center"} gap={1}>
                    <Avatar alt="Not found" src={"https://picsum.photos/200/300"} sx={{ width: 50, height: 50 }} />
                    <Typography variant="body1">{u.user.Username}</Typography>
                </Box>
                <Stack spacing={1}>
                    <Box display={"flex"} flexDirection={"row"} alignItems={"center"} gap={3} justifyContent={"start"}>
                        {
                            data.map((data, index) => (
                                <Link key={index} name={data.name} changeSelected={changeSelected} index={index} url={data.url} selected={data.selected} />
                            ))
                        }
                    </Box>
                    <Divider orientation="horizontal" />
                </Stack>
            </Stack>
        )
}
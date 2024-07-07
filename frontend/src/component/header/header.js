// components/Header.js

import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, Box, TextField, Menu, MenuItem, Avatar, IconButton } from '@mui/material';
import Link from 'next/link';
import Trello from './trello';
import Menus from './menu';
import NotificationsActiveIcon from '@mui/icons-material/NotificationsActive';
import HelpIcon from '@mui/icons-material/Help';
import { useRouter } from 'next/navigation';

const datas = [
    {
        name: "Workspaces",
        url: "/workspaces",
    },
    {
        name: "Recent",
        url: "/recent",
    },
    {
        name: "Starred",
        url: "/starred",
    },
    {
        name: "Templates",
        url: "/template"
    },
]

const Header = () => {
    const [anchorEl, setAnchorEl] = useState(null);
    const router = useRouter()

    const handleAvatarClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleProfile = () => {
        setAnchorEl(null);
        router.push("/profile");
    };

    const handleLogOut = () => {
        setAnchorEl(null);
        router.push("/login");
        localStorage.clear()
    }

    return (
        <AppBar position="static" sx={{ bgcolor: 'background.paper', color: 'text.primary' }}>
            <Toolbar sx={{ justifyContent: 'space-between' }}>
                <Box display={"flex"} alignItems={"center"} gap={2}>
                    <Trello />
                    {
                        datas.map((data, index) => (
                            <Menus key={index} name={data.name} url={data.url} />
                        ))
                    }
                    <Button component={Link} href="/create" variant="contained" sx={{ bgcolor: 'primary.main' }}>
                        Create
                    </Button>
                </Box>
                <Box display={"flex"} alignItems={"center"} gap={1}>
                    <TextField fullWidth label="Search" maxRows={1} size='small' />
                    <NotificationsActiveIcon fontSize='small' />
                    <HelpIcon fontSize='small' />
                    <IconButton onClick={handleAvatarClick}>
                        <Avatar alt={"Not found"} src={"https://picsum.photos/80/80"} />
                    </IconButton>
                    <Menu
                        anchorEl={anchorEl}
                        open={Boolean(anchorEl)}
                        onClose={() => setAnchorEl(null)}
                    >
                        <MenuItem onClick={handleProfile}>Profile</MenuItem>
                        <MenuItem onClick={() => { setAnchorEl(null) }}>Settings</MenuItem>
                        <MenuItem onClick={handleLogOut}>Logout</MenuItem>
                    </Menu>
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default Header;

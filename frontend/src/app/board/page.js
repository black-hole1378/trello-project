"use client"
import React, { useState, useEffect } from "react";
import { Box, Container, Grid, Paper, Stack, Typography, IconButton, Divider, Button } from "@mui/material";
import MoreHorizIcon from '@mui/icons-material/MoreHoriz';
import { useRouter } from "next/navigation";
import Task from "@/component/board/Task";
import axios from "axios";

export default function () {
    const [tasks, setTasks] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/workspaces/${localStorage.getItem("workSpaceID")}/tasks`, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });

                if (response.status === 200) {
                    console.log("data", response.data)
                    setTasks(response.data.tasks);  // Corrected to setTasks
                }
            } catch (error) {
                console.log("error", error)
                setError(error.message);
            }
        };

        fetchData();
    }, []);


    const router = useRouter();
    const handleClick = (status) => {
        localStorage.setItem("status", status);
        router.push("/create_task")
    }


    return (
        <Box width={"100vw"} mt={2} height={"100vh"}>
            <Container >
                <Grid container spacing={1}>
                    <Grid md={4}>
                        <Box width={350} sx={{ height: "max-content" }}>
                            <Paper sx={{ p: 1 }}>
                                <Box display={"flex"} flexDirection={"column"} justifyContent={"space-around"}>
                                    <Box display={"flex"} alignItems={"center"} justifyContent={"space-between"}>
                                        <Typography variant="body1">To Do</Typography>
                                        <IconButton>
                                            <MoreHorizIcon fontSize="medium" />
                                        </IconButton>
                                    </Box>

                                    {
                                        tasks.length ? (
                                            tasks.filter(task => task.Status == "Planned").map((task, index) => (
                                                <Task key={index} task={task} />
                                            ))
                                        ) : <div></div>
                                    }

                                    <Button variant="contained" onClick={() => handleClick("Planned")} color="info">
                                        Add to Cart
                                    </Button>
                                </Box>
                            </Paper>
                        </Box>
                    </Grid>
                    <Grid md={4}>
                        <Box width={350} sx={{ height: "max-content" }}>
                            <Paper sx={{ p: 1 }} >
                                <Box display={"flex"} flexDirection={"column"} justifyContent={"space-around"}>
                                    <Box display={"flex"} alignItems={"center"} justifyContent={"space-between"}>
                                        <Typography variant="body1">Doing</Typography>
                                        <IconButton>
                                            <MoreHorizIcon fontSize="medium" />
                                        </IconButton>
                                    </Box>
                                    {
                                        tasks.length ? (
                                            tasks.filter(task => task.Status == "In Progress").map((task, index) => (
                                                <Task key={index} task={task} />
                                            ))
                                        ) : <div></div>
                                    }
                                    <Button variant="contained" onClick={() => handleClick("In Progress")} color="info">
                                        Add to Cart
                                    </Button>
                                </Box>
                            </Paper>
                        </Box>
                    </Grid>
                    <Grid md={4}>
                        <Box width={350} sx={{ height: "max-content" }}>
                            <Paper sx={{ p: 1 }}>
                                <Box display={"flex"} flexDirection={"column"} justifyContent={"space-around"}>
                                    <Box display={"flex"} alignItems={"center"} justifyContent={"space-between"}>
                                        <Typography variant="body1">Done</Typography>
                                        <IconButton>
                                            <MoreHorizIcon fontSize="medium" />
                                        </IconButton>
                                    </Box>
                                    {
                                        tasks.length ? (
                                            tasks.filter(task => task.Status == "Completed").map((task, index) => (
                                                <Task key={index} task={task} />
                                            ))
                                        ) : <div></div>
                                    }
                                    <Button variant="contained" onClick={() => handleClick("Completed")} color="info">
                                        Add to Cart
                                    </Button>
                                </Box>
                            </Paper>
                        </Box>
                    </Grid>
                </Grid>
            </Container >
        </Box >
    )
}
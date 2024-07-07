"use client";
import React, { useState, useEffect } from "react";
import axios from "axios";
import { Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";

const Profile = () => {
    const [workSpaces, setWorkSpaces] = useState([]);
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        const fetchWorkSpacesAndTasks = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/workspaces`, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });
                console.log("hi", response.data)
                const fetchedWorkSpaces = response.data.workSpaces;
                setWorkSpaces(fetchedWorkSpaces);

                const allTasks = await Promise.all(
                    fetchedWorkSpaces.map(async (workSpace) => {
                        const tasksResponse = await axios.get(`http://localhost:8000/workspaces/${workSpace.ID}/tasks`, {
                            headers: {
                                Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                            }
                        });
                        return tasksResponse.data.tasks.map(task => ({
                            ...task,
                            WorkSpaceName: workSpace.Name
                        }));
                    })
                );

                setTasks(allTasks.flat());
            } catch (error) {
                console.error("Error fetching workspaces and tasks:", error);
            }
        };

        fetchWorkSpacesAndTasks();
    }, []);

    return (
        <Container sx={{ mt: 2 }}>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Board</TableCell>
                            <TableCell>Due Date</TableCell>
                            <TableCell>Title</TableCell>
                            <TableCell>Description</TableCell>
                            <TableCell>Status</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {tasks.map((task) => (
                            <TableRow key={task.ID}>
                                <TableCell>{task.WorkSpaceName}</TableCell>
                                <TableCell>{new Date(task.DueDate).toLocaleDateString()}</TableCell>
                                <TableCell>{task.Title}</TableCell>
                                <TableCell>{task.Description}</TableCell>
                                <TableCell>{task.Status}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Container>
    );
};

export default Profile;

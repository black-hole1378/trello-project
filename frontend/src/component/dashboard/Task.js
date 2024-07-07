// component/dashboard/Task.js
import { Card, CardContent, Typography, Box } from "@mui/material";
import React, { useState, useEffect } from "react";
import axios from "axios";

export default function Task({ workSpace }) {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        const fetchTasks = async () => {
            try {
                const response = await axios.get(`http://localhost:8000/workspaces/${workSpace.ID}/tasks`, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                });
                setTasks(response.data.tasks);
            } catch (error) {
                console.error(`Error fetching tasks for workspace ${workSpace.ID}:`, error);
            }
        };

        fetchTasks();
    }, [workSpace.ID]);

    return (
        <Card sx={{ width: '100%', minWidth: '300px', mb: 2 }}>
            <CardContent>
                <Typography variant="h6" component="div">
                    {workSpace.Name}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    {workSpace.Description}
                </Typography>
                <Box mt={2}>
                    {tasks.length > 0 ? (
                        tasks.map((task) => (
                            <Box key={task.ID} sx={{ mb: 2 }}>
                                <Typography variant="body1" component="div">
                                    {task.Title}
                                </Typography>
                                <Typography variant="body2" color="text.secondary">
                                    {task.Description}
                                </Typography>
                            </Box>
                        ))
                    ) : (
                        <Typography variant="body2" color="text.secondary">
                            No tasks available.
                        </Typography>
                    )}
                </Box>
            </CardContent>
        </Card>
    )
}

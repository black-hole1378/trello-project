import React, { useState } from "react";
import { Card, CardContent, Typography, CardMedia, IconButton, Menu, MenuItem, Chip, Box, Modal, TextField, Button } from "@mui/material";
import MoreVertIcon from '@mui/icons-material/MoreVert';
import { useRouter } from "next/navigation";
import axios from "axios";

export default function Task({ task, onDelete }) {
    const [anchorEl, setAnchorEl] = useState(null);
    const [openCommentModal, setOpenCommentModal] = useState(false);
    const [comment, setComment] = useState("");
    const router = useRouter();

    const handleMenuOpen = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleMenuClose = () => {
        setAnchorEl(null);
    };

    const handleDeleteTask = async () => {
        try {
            const response = await axios.delete(
                `http://localhost:8000/workspaces/${localStorage.getItem("workSpaceID")}/tasks/${task.ID}`,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                }
            );
            if (response.status === 200) {
                onDelete(task.ID); // Call onDelete to remove the task from the parent component's state
                setOpenDialog(false);
            }
        } catch (error) {
            console.error("Error deleting task:", error);
        }
    };

    const handleOpenCommentModal = () => {
        setOpenCommentModal(true);
        handleMenuClose();
    };

    const handleCloseCommentModal = () => {
        setOpenCommentModal(false);
    };

    const handleCommentChange = (event) => {
        setComment(event.target.value);
    };

    const handleCommentSubmit = async () => {
        try {
            const response = await axios.post(
                `http://localhost:8000/tasks/${task.ID}/comments`,
                {
                    userID: localStorage.getItem("userID"),
                    content: comment
                },
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
                    }
                }
            );
            if (response.status === 200) {
                setComment("");
                handleCloseCommentModal();
                console.log("data", response.data)
            }
        } catch (error) {
            console.error("Error creating comment:", error);
        }
    };

    const getStatusColor = (status) => {
        switch (status) {
            case "Planned":
                return "info";
            case "In Progress":
                return "warning";
            case "Completed":
                return "success";
            default:
                return "default";
        }
    };

    return (
        <>
            <Card sx={{ maxWidth: 345, mb: 2, height: 200 }}>
                <CardMedia
                    component="img"
                    height="80"
                    image={task.ImageUrl}
                    alt="task image"
                />
                <CardContent sx={{ p: 1 }}>
                    <Box display="flex" justifyContent="space-between" alignItems="center">
                        <Typography variant="h6" component="div" noWrap sx={{ fontSize: '1rem' }}>
                            {task.Title}
                        </Typography>
                        <IconButton onClick={handleMenuOpen} size="small">
                            <MoreVertIcon fontSize="small" />
                        </IconButton>
                    </Box>
                    <Menu
                        anchorEl={anchorEl}
                        open={Boolean(anchorEl)}
                        onClose={handleMenuClose}
                    >
                        <MenuItem onClick={handleMenuClose}>Edit</MenuItem>
                        <MenuItem onClick={handleDeleteTask}>Delete</MenuItem>
                        <MenuItem onClick={() => {
                            localStorage.setItem("taskID", task.ID)
                            router.push("/create_subtask")
                        }}>Create SubTask</MenuItem>
                        <MenuItem onClick={handleOpenCommentModal}>Comment</MenuItem>
                    </Menu>
                    <Typography variant="body2" color="textSecondary" noWrap sx={{ fontSize: '0.8rem' }}>
                        {task.Description}
                    </Typography>
                    <Box mt={1}>
                        <Chip label={task.Status} color={getStatusColor(task.Status)} size="small" />
                    </Box>
                    <Typography variant="caption" color="textSecondary" noWrap sx={{ fontSize: '0.7rem' }}>
                        Due: {task.DueDate}
                    </Typography>
                </CardContent>
            </Card>

            <Modal open={openCommentModal} onClose={handleCloseCommentModal}>
                <Box sx={{
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    transform: 'translate(-50%, -50%)',
                    width: 400,
                    bgcolor: 'background.paper',
                    border: '2px solid #000',
                    boxShadow: 24,
                    p: 4,
                }}>
                    <Typography variant="h6" component="h2">
                        Add a Comment
                    </Typography>
                    <TextField
                        label="Comment"
                        multiline
                        rows={4}
                        fullWidth
                        variant="outlined"
                        value={comment}
                        onChange={handleCommentChange}
                        sx={{ mt: 2, mb: 2 }}
                    />
                    <Button onClick={handleCommentSubmit} variant="contained" color="primary">
                        Submit
                    </Button>
                </Box>
            </Modal>
        </>
    );
}

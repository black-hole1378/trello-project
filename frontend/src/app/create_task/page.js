"use client";
import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import {
    TextField,
    Button,
    Grid,
    Typography,
    Container,
    Paper,
    Modal,
    IconButton,
    CircularProgress,
    Alert,
} from '@mui/material';
import { AddCircleOutline } from '@mui/icons-material';
import axios from 'axios';

const CreateTask = ({ status }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [estimatedTime, setEstimatedTime] = useState('');
    const [dueDate, setDueDate] = useState('');
    const [imageUrl, setImageUrl] = useState('');
    const [priority, setPriority] = useState('');
    const [assignee, setAssignee] = useState('');
    const [openModal, setOpenModal] = useState(false);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const router = useRouter()
    const handleFormSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        setSuccess('');

        try {
            // Format the dueDate to 'YYYY-MM-DD'
            const formattedDueDate = new Date(dueDate).toISOString().split('T')[0];

            const taskData = {
                title,
                description,
                estimatedTime,
                dueDate: formattedDueDate,
                imageUrl,
                priority,
                status,
                userName: assignee,
            };

            const response = await axios.post(
                `http://localhost:8000/workspaces/${localStorage.getItem("workSpaceID")}/tasks`,
                taskData,
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
                        'Content-Type': 'application/json',
                    },
                }
            );

            if (response.status === 200) {
                console.log("Task created:", response.data);
                setSuccess('Task created successfully');
                router.push("/board");
            }
        } catch (error) {
            console.error('Error creating task:', error);
            setError('Error creating task');
        } finally {
            setLoading(false);
        }
    };

    const handleAssigneeClick = () => {
        setOpenModal(true);
    };

    const handleModalClose = () => {
        setOpenModal(false);
    };

    const handleAssigneeModalSubmit = (e) => {
        e.preventDefault();
        setOpenModal(false);
    };

    return (
        <Container maxWidth="sm" style={{ marginTop: '50px' }}>
            <Paper elevation={3} style={{ padding: '20px' }}>
                <Typography variant="h5" gutterBottom>
                    Create New Task
                </Typography>
                {error && <Alert severity="error">{error}</Alert>}
                {success && <Alert severity="success">{success}</Alert>}
                <form onSubmit={handleFormSubmit}>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                required
                                fullWidth
                                label="Title"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                multiline
                                rows={3}
                                label="Description"
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                fullWidth
                                label="Estimated Time (in hours)"
                                value={estimatedTime}
                                onChange={(e) => setEstimatedTime(e.target.value)}
                            />
                        </Grid>
                        <Grid item xs={6}>
                            <TextField
                                fullWidth
                                type="date"
                                label="Due Date"
                                value={dueDate}
                                onChange={(e) => setDueDate(e.target.value)}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                label="Image URL"
                                value={imageUrl}
                                onChange={(e) => setImageUrl(e.target.value)}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                label="Priority"
                                value={priority}
                                onChange={(e) => setPriority(e.target.value)}
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                fullWidth
                                label="Assignee"
                                value={assignee}
                                InputProps={{
                                    endAdornment: (
                                        <IconButton onClick={handleAssigneeClick}>
                                            <AddCircleOutline />
                                        </IconButton>
                                    ),
                                }}
                                disabled
                            />
                        </Grid>
                        <Grid item xs={12}>
                            <Button type="submit" variant="contained" color="primary" disabled={loading}>
                                {loading ? <CircularProgress size={24} /> : 'Create Task'}
                            </Button>
                        </Grid>
                    </Grid>
                </form>
            </Paper>

            {/* Assignee Modal */}
            <Modal
                open={openModal}
                onClose={handleModalClose}
                aria-labelledby="assignee-modal-title"
                aria-describedby="assignee-modal-description"
                style={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                }}
            >
                <Paper elevation={3} style={{ padding: '20px', maxWidth: '500px' }}>
                    <Typography variant="h6" id="assignee-modal-title" gutterBottom>
                        Assign Task
                    </Typography>
                    <form onSubmit={handleAssigneeModalSubmit}>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <TextField
                                    required
                                    fullWidth
                                    label="Assignee Name"
                                    value={assignee}
                                    onChange={(e) => setAssignee(e.target.value)}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Button type="submit" variant="contained" color="primary">
                                    Assign
                                </Button>
                            </Grid>
                        </Grid>
                    </form>
                </Paper>
            </Modal>
        </Container>
    );
};

export default CreateTask;

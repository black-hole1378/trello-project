import { Container, Grid, Box, Button, Typography } from '@mui/material';
import LoginForm from '../../component/login/LoginForm';
import SocialLoginButtons from '../../component/login/SocialLogin';
import React from 'react';
import Link from 'next/link';

const LoginPage = () => {
    return (
        <Container
            component="main"
            maxWidth="xs"
            sx={{
                backgroundImage: 'url(https://source.unsplash.com/random)',
                backgroundSize: 'cover',
                backgroundPosition: 'center',
                minHeight: '100vh',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
            }}
        >
            <Box sx={{
                backgroundColor: 'rgba(255, 255, 255, 0.8)',
                padding: 4,
                borderRadius: 2,
                boxShadow: 3,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
            }}>
                <LoginForm />
                <SocialLoginButtons />
                <Grid container justifyContent="space-between" sx={{ mt: 2 }}>
                    <Grid item>
                        <Button href="#" variant="text">
                            Can't log in?
                        </Button>
                    </Grid>
                    <Grid item>
                        <Link style={{ textDecoration: "none", outline: "none" }} href="/signup">
                            <Typography variant='body1' mt={0.8} color={"text.information"}>
                                Create an account
                            </Typography>
                        </Link>
                    </Grid>
                </Grid>
            </Box>
        </Container>
    );
};

export default LoginPage;

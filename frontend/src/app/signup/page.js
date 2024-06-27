import { Container, Box } from '@mui/material';
import SignUpForm from '../../component/signup/signupForm';

const SignupPage = () => {
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
            <Box
                sx={{
                    backgroundColor: 'rgba(255, 255, 255, 0.8)',
                    padding: 4,
                    borderRadius: 2,
                    boxShadow: 3,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <SignUpForm />
            </Box>
        </Container>
    );
};

export default SignupPage;

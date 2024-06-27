import { Button, Grid, Typography } from '@mui/material';
import { Google as GoogleIcon } from '@mui/icons-material';
import React from 'react';

const SocialLoginButtons = () => {
    return (
        <>
            <Typography variant="body2" align="center" sx={{ mt: 2 }}>
                Or continue with:
            </Typography>
            <Grid container spacing={2} sx={{ mt: 1 }}>
                <Grid item xs={12}>
                    <Button
                        fullWidth
                        variant="outlined"
                        startIcon={<GoogleIcon />}
                    >
                        Google
                    </Button>
                </Grid>
            </Grid>
        </>
    );
};

export default SocialLoginButtons;

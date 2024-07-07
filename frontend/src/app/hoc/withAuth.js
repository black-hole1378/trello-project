// app/hoc/withAuth.js
"use client"
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

const withAuth = (WrappedComponent) => {
    return (props) => {
        const router = useRouter();

        useEffect(() => {
            const accessToken = localStorage.getItem('accessToken');
            console.log("access", accessToken)
            if (!accessToken) {
                router.push('/login');
            }
        }, []);

        return <WrappedComponent {...props} />;
    };
};

export default withAuth;

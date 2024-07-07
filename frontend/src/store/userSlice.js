// src/store/userSlice.js
import { createSlice } from '@reduxjs/toolkit';

const userSlice = createSlice({
    name: 'user',
    initialState: {
        userID: null,
        userName: null,
    },
    reducers: {
        setUser(state, action) {
            state.userID = action.payload.userID;
            state.userName = action.payload.userName;
        },
    },
});

export const { setUser } = userSlice.actions;
export default userSlice.reducer;

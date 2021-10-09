import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

import { RootState } from './store';
import Api from '../api/api';


const initialState = {
  isLogin: false,
  accessToken: '',
};

const appSlice = createSlice({
  name: 'app',
  initialState,
  reducers: {
    f(state, action) {
      state.isLogin = action.payload;
    }
  },
  extraReducers: {},
});

// export const { } = appSlice.actions;

export default appSlice.reducer;

// export const selectLoginStatus = (state) => state.app.isLogin;
// export const selectAccessToken = (state) => state.app.accessToken;
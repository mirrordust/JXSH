import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

import {
  Api,
  ApiStatus,
  initialApiStatus,
  handleValidateResponse
} from '../api/api';
import { Credential, User } from '../api/model';
import { RootState } from './store';


type AppState = ApiStatus & {
  isLogin: boolean;
  credential: Credential;
};

const initialState: AppState = {
  ...initialApiStatus,
  isLogin: false,
  credential: { access_token: '' }
};

export const login = createAsyncThunk(
  'app/login',
  async (user: User) => {
    const response = await Api.login(user);
    return handleValidateResponse(200, response);
  }
);

export const logout = createAsyncThunk(
  'app/logout',
  async (credential: Credential) => {
    const response = await Api.logout(credential);
    return handleValidateResponse(204, response);
  }
);

const appSlice = createSlice(
  {
    name: 'app',
    initialState,
    reducers: {},
    extraReducers(builder) {
      builder
        // login
        .addCase(login.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.error = undefined;
          state.isLogin = true;
          state.credential = { access_token: action.payload.data.access_token };
        })
        .addCase(login.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        // logout
        .addCase(logout.fulfilled, (state, _action) => {
          state.status = 'succeeded';
          state.error = undefined;
          state.isLogin = false;
          state.credential = { access_token: '' };
        })
        .addCase(logout.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
    }
  }
);

export const selectAppisLogin = (state: RootState) => state.app.isLogin;
export const selectCredential = (state: RootState) => state.app.credential;
export const selectAppStatus = (state: RootState) => state.app.status;
export const selectAppError = (state: RootState) => state.app.error;

export const appReducer = appSlice.reducer;

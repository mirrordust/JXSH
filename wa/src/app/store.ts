import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';

import { appReducer } from './appSlice';
import { tagsReducer } from '../features/tags/tagsSlice';


const store = configureStore({
  reducer: {
    app: appReducer,
    tags: tagsReducer,
  },
});

export default store;

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;

import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';

import { appReducer } from './appSlice';
import { postsReducer } from '../features/posts/postsSlice';
import { tagsReducer } from '../features/tags/tagsSlice';


const store = configureStore({
  reducer: {
    app: appReducer,
    posts: postsReducer,
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

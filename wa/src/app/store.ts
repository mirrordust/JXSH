import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';

// import appReducer from './appSlice';
import postsReducer from '../features/posts/postsSlice';
import usersReducer from '../features/users/usersSlice';
import notificationsReducer from '../features/notifications/notificationsSlice';


const store = configureStore({
  reducer: {
    // app: appReducer,
    posts: postsReducer,
    users: usersReducer,
    notifications: notificationsReducer,
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

import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice
} from '@reduxjs/toolkit';

import { Api, initialApiStatus, handleValidateResponse } from '../../api/api';
import { Credential, Post } from '../../api/model';
import { RootState } from '../../app/store';


const postAdapter = createEntityAdapter<Post>({
  selectId: post => post.id,
  sortComparer: (a, b) => {
    if (a.id > b.id) {
      return -1;
    } else if (a.id < b.id) {
      return 1;
    } else {
      return 0;
    }
  }
});

const initialState = postAdapter.getInitialState(initialApiStatus);

export const fetchPosts = createAsyncThunk(
  'posts/fetchPosts',
  async (credential: Credential) => {
    const response = await Api.getAll(credential, 'posts');
    return handleValidateResponse(200, response);
  }
);

export const createPost = createAsyncThunk(
  'posts/addPost',
  async (params: { credential: Credential, post: Post }) => {
    const { credential, post } = params;
    const response = await Api.create(credential, 'posts', post);
    return handleValidateResponse(201, response);
  }
);

export const updatePost = createAsyncThunk(
  'posts/updatePost',
  async (params: { credential: Credential, id: number, post: Post }) => {
    const { credential, id, post } = params;
    const response = await Api.updateById(credential, 'posts', id, post);
    return handleValidateResponse(200, response);
  }
);

export const deletePost = createAsyncThunk(
  'posts/deletePost',
  async (params: { credential: Credential, id: number }) => {
    const { credential, id } = params;
    const response = await Api.deleteById(credential, 'posts', id);
    response.data = { id };
    return handleValidateResponse(204, response);
  }
);

const postsSlice = createSlice(
  {
    name: 'posts',
    initialState,
    reducers: {},
    extraReducers(builder) {
      builder
        // fetch
        .addCase(fetchPosts.pending, (state, action) => {
          state.status = 'loading';
        })
        .addCase(fetchPosts.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.error = undefined;
          postAdapter.upsertMany(state, action.payload.data);
        })
        .addCase(fetchPosts.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        // create
        .addCase(createPost.pending, (state, action) => {
          state.status = 'loading';
        })
        .addCase(createPost.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.error = undefined;
          postAdapter.addOne(state, action.payload.data);
        })
        .addCase(createPost.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        // update
        .addCase(updatePost.pending, (state, action) => {
          state.status = 'loading';
        })
        .addCase(updatePost.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.error = undefined;
          postAdapter.upsertOne(state, action.payload.data);
        })
        .addCase(updatePost.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
        // delete
        .addCase(deletePost.pending, (state, action) => {
          state.status = 'loading';
        })
        .addCase(deletePost.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.error = undefined;
          postAdapter.removeOne(state, action.payload.id);
        })
        .addCase(deletePost.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        })
    }
  }
);

export const {
  selectAll: selectAllPosts,
  selectById: selectPostById,
  selectIds: selectPostsIds
} = postAdapter.getSelectors((state: RootState) => state.posts);

// export const selectPostsByTag =createSelector(
//   [selectAllPosts, (state, tagId)=>tagId],
//   (posts, tagId) => 
// );

export const selectPostsStatus = (state: RootState) => state.posts.status;
export const selectPostsError = (state: RootState) => state.posts.error;

export const postsReducer = postsSlice.reducer;

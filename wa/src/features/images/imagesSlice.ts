import { createSlice } from '@reduxjs/toolkit';

import { RootState } from '../../app/store';


type ImageState = {
  url1: string;
  url2: string;
};

const initialState: ImageState = {
  url1: '',
  url2: ''
};

const imagesSlice = createSlice(
  {
    name: 'images',
    initialState,
    reducers: {
      url1Updated(state, action) {
        state.url1 = action.payload;
      },
      url2Updated(state, action) {
        state.url2 = action.payload;
      }
    }
  }
);

export const { url1Updated, url2Updated } = imagesSlice.actions;

export const selectImageUrl1 = (state: RootState) => state.images.url1;
export const selectImageUrl2 = (state: RootState) => state.images.url2;

export const imagesReducer = imagesSlice.reducer;

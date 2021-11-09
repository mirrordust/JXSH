import React, { useState } from 'react';
import { useParams } from 'react-router';
import MarkdownIt from 'markdown-it';
import MdEditor, { Plugins } from 'react-markdown-editor-lite';
import {
  Button,
  Col,
  Container,
  Form,
  Row
} from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from '../../app/hooks';
import { selectCredential } from '../../app/appSlice';
import {
  createPost,
  selectPostById,
  selectPostsError,
  selectPostsStatus,
  updatePost
} from './postsSlice';
import { StatusBar } from '../../components/StatusBar';

import 'react-markdown-editor-lite/lib/index.css';


// initialize a markdown parser
const mdParser = new MarkdownIt(/* Markdown-it options */);

// Register plugins (globally)
MdEditor.use(Plugins.TabInsert, {
  /**
   * 用户按下 Tab 键时输入的空格的数目
   * 特别地，1 代表输入一个'\t'，而不是一个空格
   * 默认值是 1
   */
  tabMapValue: 1
});
// Plugins.FullScreen.align = 'left';

// Register plugins (editor instance)
const plugins = [
  'font-bold',
  'font-italic',
  'font-underline',
  'font-strikethrough',
  'tab-insert',
  'header',
  'list-unordered',
  'list-ordered',
  'block-quote',
  'block-wrap',
  'block-code-inline',
  'block-code-block',
  'link',
  'image',
  'table',
  'divider',
  'clear',
  'logger',
  'mode-toggle',
  // 'full-screen',
];

export const WEditor = ({ postId }) => {
  const params = useParams();
  let pathParamPostId;
  if (params.postId) {
    pathParamPostId = parseInt(params.postId, 10);
  }

  let leadingPostId;
  if (pathParamPostId) {
    leadingPostId = pathParamPostId;
  }
  if (postId) {
    leadingPostId = postId;
  }
  const isNew = leadingPostId ? false : true;

  const dispatch = useAppDispatch();
  const post = useAppSelector(state => selectPostById(state, leadingPostId));
  const appCredential = useAppSelector(selectCredential);
  const postsStatus = useAppSelector(selectPostsStatus);
  const postsError = useAppSelector(selectPostsError);

  let pid = -1, t = '', c = '',
    pub = false, vn = '', tags = [];
  if (post) {
    pid = post.id;
    t = post.title;
    c = post.body;
    pub = post.published;
    vn = post.view_name;
    tags = post.tags;
  }

  const [isNewPost, setIsNewPost] = useState(isNew);
  const [postId_s, setPostId_s] = useState(pid);

  const [postTitle, setPostTitle] = useState(t);
  const onPostTitleChanged = async (e) => {
    const newValue = e.target.value;
    setPostTitle(newValue);
    await realTimeSave({ title: newValue });
  };

  const [postBody, setPostBody] = useState(c);
  const onPostBodyChanged = async ({ html, text }) => {
    const newValue = text;
    setPostBody(newValue);
    await realTimeSave({ body: newValue });
  };

  const [postViewName, setPostViewName] = useState(vn);
  const onPostViewNameChanged = async (e) => {
    const newValue = e.target.value;
    setPostViewName(newValue);
    await realTimeSave({ view_name: newValue });
  };

  const [postPublished, setPostPublished] = useState(pub);
  const onPostPublishedClicked = async () => {
    const newValue = !postPublished
    setPostPublished(newValue);
    await realTimeSave({ published: newValue }, true);
  };

  const [postTags, setPostTags] = useState(tags);
  const onPostTagsChanged = async () => {
    await realTimeSave(true);
  };

  async function doCreatePost() {
    return await dispatch(createPost({
      credential: appCredential,
      post: {
        id: postId_s,
        title: postTitle,
        body: postBody,
        published: postPublished,
        view_name: postViewName,
        tags: postTags
      }
    })).unwrap();
  }

  /**
   * Update or patch post.
   * 
   * @param {object} changes Post or partial Post object
   * @returns api result
   */
  async function doPutOrPatchPost(changes) {
    let post1;
    if (changes) {
      post1 = { ...changes }
    } else {
      post1 = {
        title: postTitle,
        body: postBody,
        published: postPublished,
        view_name: postViewName,
        tags: postTags
      }
    }
    return await dispatch(updatePost({
      credential: appCredential,
      id: postId_s,
      post: post1
    })).unwrap();
  }

  const canUpdate = !isNewPost && postsStatus !== 'loading';
  const realTimeSave = async (changes, forceUpdate) => {
    if (changes) {
      if (forceUpdate || canUpdate) {
        try {
          await doPutOrPatchPost(changes);
        } catch (err) {
          console.error('Failed to real-time update the post: ', err);
        }
      }
    }
  };

  const onCreateClicked = async () => {
    try {
      const data = await doCreatePost();
      handleCreateOrUpdate(data);
    } catch (err) {
      console.error('Create Post error: ', err);
    }
  };

  const onUpdateClicked = async () => {
    try {
      const data = await doPutOrPatchPost();
      handleCreateOrUpdate(data);
    } catch (err) {
      console.error('Update Post error: ', err);
    }
  };

  function handleCreateOrUpdate(data) {
    const post = data.data;
    setIsNewPost(false);
    setPostId_s(post.id);
    setPostTitle(post.title);
    setPostBody(post.body);
    setPostViewName(post.view_name);
    setPostPublished(post.published);
    setPostTags(post.tags);
  }

  let publishButtonText = postPublished ? 'Revoke publish' : 'publish';

  const status = <StatusBar
    status={postsStatus}
    error={postsError}
  />;

  return (
    <React.Fragment>
      Status{' '}{status}
      <Container fluid >
        <Form>
          <Row>
            <Col>
              <Form.Control
                id="postTitle"
                size="sm"
                type="text"
                placeholder="Title here..."
                value={postTitle}
                onChange={onPostTitleChanged}
              />
            </Col>
            <Col>
              <Form.Control
                id="postViewName"
                size="sm"
                type="text"
                placeholder="View name here..."
                value={postViewName}
                onChange={onPostViewNameChanged}
              />
            </Col>
          </Row>
        </Form>
      </Container>
      <MdEditor
        plugins={plugins}
        style={{ height: '400px' }}
        table={{ maxRow: 10, maxCol: 12 }}
        value={postBody}
        renderHTML={text => mdParser.render(text)}
        onChange={onPostBodyChanged}
      />
      <Button
        size="sm"
        disabled={!isNewPost}
        onClick={onCreateClicked}
      >
        Create
      </Button>{' '}
      <Button
        size="sm"
        disabled={isNewPost}
        onClick={onUpdateClicked}
      >
        Update
      </Button>{' '}
      <Button
        size="sm"
        disabled={isNewPost}
        onClick={onPostPublishedClicked}
      >
        {publishButtonText}
      </Button>
    </React.Fragment>
  );
};

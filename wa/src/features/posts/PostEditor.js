import { useEffect, useState } from 'react';
import { useLocation, useParams } from 'react-router';
import MarkdownIt from 'markdown-it';
import MdEditor, { Plugins } from 'react-markdown-editor-lite';
import {
  Button,
  ButtonGroup,
  Col,
  Container,
  Dropdown,
  DropdownButton,
  Form,
  FormControl,
  InputGroup,
  Row
} from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from '../../app/hooks';
import { selectCredential } from '../../app/appSlice';
// eslint-disable-next-line no-unused-vars
import { Post } from '../../api/model';
import {
  createPost,
  updatePost,
  selectPostsError,
  selectPostsStatus,
  selectPostById,
  updateCurrentEditorUri
} from './postsSlice';
import {
  fetchTags,
  selectAllTags,
  selectTagsStatus
} from '../tags/tagsSlice';
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

let DeletableTag = ({ tag, onDeleteClicked }) => {
  return (
    <ButtonGroup>
      <Button variant="info" size="sm">{tag.name}</Button>
      <Button variant="secondary" size="sm" onClick={onDeleteClicked}>x</Button>
    </ButtonGroup>
  );
};

let NewTagInput = ({ existingPostTags, onAddTag }) => {
  const dispatch = useAppDispatch();

  const tagsStatus = useAppSelector(selectTagsStatus);
  const appCredential = useAppSelector(selectCredential);

  useEffect(() => {
    if (tagsStatus === 'idle') {
      dispatch(fetchTags(appCredential));
    }
  }, [tagsStatus, appCredential, dispatch]);

  const [name, setName] = useState('');

  const existingPostTagNames = existingPostTags.map(t => t.name);
  const allTags = useAppSelector(selectAllTags);

  const selectableTags = allTags.filter(t =>
    existingPostTagNames.indexOf(t.name) === -1
  );
  const dropdowns = selectableTags.map(t =>
    <Dropdown.Item key={t.id} onClick={onAddTag(t)}>
      {t.name}
    </Dropdown.Item>
  );

  return (
    <InputGroup className="mb-3" size="sm">
      <FormControl
        value={name}
        onChange={e => setName(e.target.value)}
      />
      <DropdownButton
        id="input-group-dropdown-2"
        variant="outline-secondary"
        title="选择现有标签"
        align="end"
      >
        {dropdowns}
      </DropdownButton>
      <Button onClick={onAddTag(name)} >新增</Button>
    </InputGroup>
  );
};

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

  const location = useLocation();
  let uri;
  if (location.state) {
    uri = location.state.currentEditorUri;
  }
  dispatch(updateCurrentEditorUri(uri))

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
    // setPostPublished(newValue);
    await realTimeSave({ published: newValue }, true);
  };

  const [postTags, setPostTags] = useState(tags);

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
   * @param {Post} changes Post or partial Post object
   * @returns api result
   */
  async function doPutOrPatchPost(changes) {
    let post1;
    if (changes) {
      let tags1 = postTags;
      if (changes.tags) {
        tags1 = changes.tags;
      }
      post1 = {
        ...changes,
        id: postId_s,
        /** 以 patch 方式进行部分更新时必须带上 tags 属性，
         * 保证 Post.tags的正确，因为后端的 tags 是 many_to_many 关系，
         * 不带 tags 属性的话会丢失标签信息。
         * 具体参考：https://hexdocs.pm/ecto/2.2.11/associations.html#persistence
         * */
        tags: tags1
      }
    } else {
      post1 = {
        id: postId_s,
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
          const data = await doPutOrPatchPost(changes);
          handleCreateOrUpdate(data);
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
      dispatch(updateCurrentEditorUri(`/posts/${data.data.id}`));
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

  const exsitingPostTags = postTags.map(t =>
    <DeletableTag
      key={t.id}
      tag={t}
      onDeleteClicked={
        async () => {
          const newPostTags = postTags.filter(t1 => t1.id !== t.id);
          await realTimeSave({ tags: newPostTags }, true);
        }
      }
    />
  );

  const onAddTag = (value) =>
    async () => {
      if (value) {
        let newPostTags;
        if (value.id) {  // add an existing tag
          newPostTags = [...postTags, value];
        } else {  // a brand new tag
          newPostTags = [...postTags, { name: value }];
        }
        await realTimeSave({ tags: newPostTags }, true);
      }
    };
  const newTaginput = <NewTagInput
    existingPostTags={postTags}
    onAddTag={onAddTag}
  />

  return (
    <>
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

      {exsitingPostTags}

      {newTaginput}

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
    </>
  );
};

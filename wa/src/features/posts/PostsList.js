import { useEffect } from 'react';
import { Link } from 'react-router-dom';
import {
  Badge,
  Button,
  Card,
  Col,
  Container,
  Row
} from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from '../../app/hooks';
import { selectCredential } from '../../app/appSlice';
import {
  fetchPosts,
  deletePost,
  selectPostsError,
  selectPostsStatus,
  selectPostById,
  selectPostsIds
} from './postsSlice';
import { Refresh, StatusBar } from '../../components/StatusBar';
import { correctDate } from '../../utils/utils';


let PostTags = ({ postId }) => {
  const post = useAppSelector((state) => selectPostById(state, postId))
  // must creating a copy of post.tags before sorting
  const orderedTags = [...post.tags].sort((a, b) => a.name.localeCompare(b.name));
  const tags = orderedTags.map(tag =>
    <Badge pill bg="primary" key={tag.id}>
      {tag.name}
    </Badge>
  );
  return tags;
};

let PostExcerpt = ({ postId }) => {
  const dispatch = useAppDispatch();

  const post = useAppSelector((state) => selectPostById(state, postId))
  const appCredential = useAppSelector(selectCredential);

  const onDeletePostClicked = async () => {
    try {
      await dispatch(deletePost({ credential: appCredential, id: postId })).unwrap();
    } catch (err) {
      console.error('Delete Post error: ', err);
    }
  };

  let publishedDate;
  if (post.published_at) {
    if (post.published) {
      publishedDate = correctDate(post.published_at);
    } else {
      publishedDate = '发布已撤回';
    }
  } else {
    publishedDate = '未发布';
  }

  return (
    <Card>
      <Card.Header>{post.title} ({post.view_name})</Card.Header>
      <Card.Body>
        <PostTags postId={post.id} />
        <Card.Text>{post.body.substring(0, 100)}</Card.Text>
        <Card.Text>
          <Badge bg="warning">
            创建：{correctDate(post.inserted_at)}
          </Badge>{' '}
          <Badge bg="secondary">
            更新：{correctDate(post.updated_at)}
          </Badge>{' '}
          <Badge bg="success">
            发布：{publishedDate}
          </Badge>{' '}
          <Badge bg="info">
            浏览：{post.views}
          </Badge>
        </Card.Text>
        <Row>
          <Col className="text-left">
            <Card.Link as={Link} to={`/posts/${post.id}`}>Edit</Card.Link>
          </Col>
          <Col className="text-center">
            <Button
              className="mb-2"
              variant="danger"
              size="sm"
              onClick={onDeletePostClicked}
            >
              Delete
            </Button>
          </Col>
        </Row>
      </Card.Body>
    </Card>
  );
};


export const PostsList = () => {
  const dispatch = useAppDispatch();

  const orderedPostIds = useAppSelector(selectPostsIds);
  const postsStatus = useAppSelector(selectPostsStatus);
  const postsError = useAppSelector(selectPostsError);
  const appCredential = useAppSelector(selectCredential);

  const onRefreshClicked = async () => {
    try {
      await dispatch(fetchPosts(appCredential)).unwrap();
    } catch (err) {
      console.error('Refresh Posts error: ', err);
    }
  };

  useEffect(() => {
    if (postsStatus === 'idle') {
      dispatch(fetchPosts(appCredential))
    }
  }, [postsStatus, appCredential, dispatch]);

  const posts = orderedPostIds.map((postId) => (
    <PostExcerpt key={postId} postId={postId} />
  ));

  const status = <StatusBar
    status={postsStatus}
    error={postsError}
  />;

  return (
    <Container>
      <Refresh onRefreshClicked={onRefreshClicked} />{' '}
      Status{' '}{status}
      <hr />
      <Link to="/posts/new">New Posts</Link>
      <hr />
      Posts
      {posts}
    </Container>
  );
};

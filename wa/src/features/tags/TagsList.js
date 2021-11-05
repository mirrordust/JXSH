import { useState, useEffect } from 'react';
import {
  Alert,
  Badge,
  Button,
  Col,
  Container,
  Form,
  Row,
  Spinner
} from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from '../../app/hooks';
import {
  fetchTags,
  createTag,
  updateTag,
  deleteTag,
  selectAllTags,
  selectTagsError,
  selectTagsStatus,
  selectTagById
} from './tagsSlice';
import { selectCredential } from '../../app/appSlice';


let TagInfo = ({ tagId }) => {
  const dispatch = useAppDispatch();

  const tag = useAppSelector((state) => selectTagById(state, tagId));
  const appCredential = useAppSelector(selectCredential);

  const [tagName, setTagName] = useState(tag.name);
  const onTagNameChanged = (e) => setTagName(e.target.value);

  const [editable, setEditable] = useState(false);
  const onEditableClicked = () => {
    setEditable(!editable);
  };

  const onUpdateTagClicked = async () => {
    try {
      await dispatch(updateTag({
        credential: appCredential,
        id: tagId,
        tag: { ...tag, name: tagName }
      })).unwrap();
      setEditable(false);
    } catch (err) {
      console.error('Update Tag error: ', err);
    }
  };

  const onDeleteTagClicked = async () => {
    try {
      await dispatch(deleteTag({ credential: appCredential, id: tagId })).unwrap();
    } catch (err) {
      console.error('Delete Tag error: ', err);
    }
  };

  return (
    <Row className="align-items-center">
      <Col xs="auto">
        <Form.Control
          className="mb-2"
          id="tagName"
          value={tagName}
          readOnly={!editable}
          onChange={onTagNameChanged}
        />
      </Col>
      <Col xs="auto">
        <Button className="mb-2" variant="info" onClick={onEditableClicked}>
          Edit
        </Button>
      </Col>
      <Col xs="auto">
        <Button className="mb-2" variant="warning" onClick={onUpdateTagClicked}>
          Update
        </Button>
      </Col>
      <Col xs="auto">
        <Button className="mb-2" variant="danger" onClick={onDeleteTagClicked}>
          Delete
        </Button>
      </Col>
    </Row>
  );
};

export const TagsList = () => {
  const dispatch = useAppDispatch();

  const tagList = useAppSelector(selectAllTags);
  const tagsError = useAppSelector(selectTagsError);
  const tagsStatus = useAppSelector(selectTagsStatus);
  const appCredential = useAppSelector(selectCredential);

  const [newTagName, setNewTagName] = useState('');
  const onNewTagNameChanged = (e) => setNewTagName(e.target.value);

  const onCreateTagClicked = async () => {
    try {
      await dispatch(createTag({
        credential: appCredential,
        tag: { name: newTagName }
      })).unwrap();
      setNewTagName('');
    } catch (err) {
      console.error('Update Tag error: ', err);
    }
  };

  useEffect(() => {
    if (tagsStatus === 'idle') {
      dispatch(fetchTags(appCredential));
    }
  }, [tagsStatus, appCredential, dispatch]);

  const tags = tagList.map((tag) => (
    <TagInfo key={tag.id} tagId={tag.id} />
  ));

  let status;
  if (tagsStatus === 'idle') {
    status = <Badge bg="secondary">Idle</Badge>
  } else if (tagsStatus === 'loading') {
    status = <Spinner animation="border" variant="primary" />
  } else if (tagsStatus === 'succeeded') {
    status = <Badge bg="success">Succeeded</Badge>
  } else {
    status = <Badge bg="danger">Failed</Badge>
  }

  let error;
  if (tagsError) {
    error = <Alert variant="danger">{tagsError}</Alert>;
  }

  return (
    <Container>
      Status: {status}
      {error}
      <hr />
      New Tags:
      <Row className="align-items-center">
        <Col xs="auto">
          <Form.Control
            className="mb-2"
            id="newTagName"
            value={newTagName}
            onChange={onNewTagNameChanged}
          />
        </Col>
        <Col xs="auto">
          <Button className="mb-2" variant="primary" onClick={onCreateTagClicked}>
            Create
          </Button>
        </Col>
      </Row>
      <hr />
      Tags:
      {tags}
    </Container>
  );
};

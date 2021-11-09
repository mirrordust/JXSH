import { useState, useEffect } from 'react';
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
  fetchTags,
  createTag,
  updateTag,
  deleteTag,
  selectTagsError,
  selectTagsStatus,
  selectTagById,
  selectTagsIds
} from './tagsSlice';
import { Refresh, StatusBar } from '../../components/StatusBar';


let TagInfo = ({ tagId }) => {
  const dispatch = useAppDispatch();

  const tag = useAppSelector(state => selectTagById(state, tagId));
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
          size="sm"
          value={tagName}
          readOnly={!editable}
          onChange={onTagNameChanged}
        />
      </Col>
      <Col xs="auto">
        <Button
          className="mb-2"
          variant="info"
          size="sm"
          onClick={onEditableClicked}
        >
          Edit
        </Button>
      </Col>
      <Col xs="auto">
        <Button
          className="mb-2"
          variant="warning"
          size="sm"
          onClick={onUpdateTagClicked}
          disabled={!editable}
        >
          Update
        </Button>
      </Col>
      <Col xs="auto">
        <Button
          className="mb-2"
          variant="danger"
          size="sm"
          onClick={onDeleteTagClicked}
        >
          Delete
        </Button>
      </Col>
    </Row>
  );
};

export const TagsList = () => {
  const dispatch = useAppDispatch();

  const tagIds = useAppSelector(selectTagsIds);
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

  const onRefreshClicked = async () => {
    try {
      await dispatch(fetchTags(appCredential)).unwrap();
    } catch (err) {
      console.error('Refresh Tags error: ', err);
    }
  };

  useEffect(() => {
    if (tagsStatus === 'idle') {
      dispatch(fetchTags(appCredential));
    }
  }, [tagsStatus, appCredential, dispatch]);

  const tags = tagIds.map((tagId) => (
    <TagInfo key={tagId} tagId={tagId} />
  ));

  const status = <StatusBar
    status={tagsStatus}
    error={tagsError}
  />;

  return (
    <Container>
      <Refresh onRefreshClicked={onRefreshClicked} />{' '}
      Status{' '}{status}
      <hr />
      New Tags
      <Row className="align-items-center">
        <Col xs="auto">
          <Form.Control
            className="mb-2"
            id="newTagName"
            size="sm"
            value={newTagName}
            onChange={onNewTagNameChanged}
          />
        </Col>
        <Col xs="auto">
          <Button
            className="mb-2"
            variant="primary"
            size="sm"
            onClick={onCreateTagClicked}
          >
            Create
          </Button>
        </Col>
      </Row>
      <hr />
      Tags
      {tags}
    </Container>
  );
};

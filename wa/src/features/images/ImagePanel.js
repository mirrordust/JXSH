import { Button, Container, Form } from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from '../../app/hooks';
import { selectImageUrl1, selectImageUrl2, url1Updated, url2Updated } from './imagesSlice';


export const ImagePanel = () => {
  const dispatch = useAppDispatch();

  const url1 = useAppSelector(selectImageUrl1);
  const url2 = useAppSelector(selectImageUrl2);

  /**
   * Convert url like `https://1drv.ms/u/s!AmM_sfZ5DhQ_cnXDrcVH0NDoZ0Q` to
   * `https://api.onedrive.com/v1.0/shares/s!AmM_sfZ5DhQ_cnXDrcVH0NDoZ0Q/root/content`
   */
  const onGenerateClick = () => {
    const ss = url1.split('/');
    const s = ss[ss.length - 1];
    const url = `https://api.onedrive.com/v1.0/shares/${s}/root/content`;
    dispatch(url2Updated(url));
  };

  const onCopyClick = () => {
    navigator.clipboard.writeText(url2)
  };

  const onClearClick = () => {
    dispatch(url1Updated(''));
    dispatch(url2Updated(''));
  };

  return (
    <Container>
      <Form.Control
        type="text"
        placeholder="Enter url"
        value={url1}
        onChange={e => dispatch(url1Updated(e.target.value))}
      />

      <Button variant="success" onClick={onGenerateClick} >生成</Button>
      {' '}
      <Button variant="warning" onClick={onClearClick} >清除</Button>
      {' '}
      <Button variant="info" onClick={onCopyClick} >复制</Button>

      <Form.Control type="text" disabled value={url2} />

      <img src={url2} alt="the img will show here..." />
    </Container>
  );
};
import { Badge, Button } from 'react-bootstrap';


/**
 * Show api status and error.
 * 
 * @param {object} param0 {status: status, error: error}
 * @returns Status Bar
 */
export const StatusBar = ({ status, error }) => {
  let statusInfo;
  if (status === 'idle') {
    statusInfo =
      <Badge bg="secondary">
        Idle
      </Badge>;
  } else if (status === 'loading') {
    statusInfo =
      <Badge bg="secondary">
        Loading
      </Badge>;
  } else if (status === 'succeeded') {
    statusInfo =
      <Badge bg="success">
        Succeeded
      </Badge>;
  } else {
    statusInfo =
      <Badge bg="danger">
        {error || 'Failed'}
      </Badge>;
  }

  return statusInfo;
};

export const Refresh = ({ onRefreshClicked }) => {
  return (
    <Button
      className="mb-2"
      variant="info"
      size="sm"
      onClick={onRefreshClicked}
    >
      Refresh
    </Button>
  );
}

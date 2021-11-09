import { useState } from 'react';
import {
  Button,
  Container,
  Col,
  Form,
  Row
} from 'react-bootstrap';

import { useAppDispatch, useAppSelector } from './hooks';
import { login, logout, selectCredential, selectAppError, selectAppisLogin, selectAppStatus } from './appSlice';
import { StatusBar } from '../components/StatusBar';


export const LoginPanel = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const dispatch = useAppDispatch();
  const appStatus = useAppSelector(selectAppStatus);
  const appError = useAppSelector(selectAppError);
  const appIsLogin = useAppSelector(selectAppisLogin);
  const appCredential = useAppSelector(selectCredential);

  const onEmailChanged = (e) => setEmail(e.target.value);
  const onPasswordChanged = (e) => setPassword(e.target.value);

  const onLogin = async () => {
    try {
      await dispatch(login({ email, password })).unwrap();
      setEmail('');
      setPassword('');
    } catch (err) {
      console.error('Failed to login: ', err)
    }
  };

  const onLogout = async () => {
    try {
      await dispatch(logout(appCredential)).unwrap();
    } catch (err) {
      console.error('Failed to logout: ', err)
    }
  };

  const status = <StatusBar
    status={appStatus}
    error={appError}
  />

  let sessionInfo;
  if (appIsLogin) {
    sessionInfo =
      <div>
        <h2>Already logged in.</h2>
        <hr />
        <p>access_token: {appCredential.access_token}</p>
        <Button
          variant="warning"
          onClick={onLogout}
        >
          Logout</Button>
      </div>;
  } else {
    sessionInfo =
      <Form className="mx-auto my-20">
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={onEmailChanged}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            value={password}
            onChange={onPasswordChanged}
          />
        </Form.Group>

        <Button
          variant="primary"
          onClick={onLogin}
        >
          Login
        </Button>
      </Form>;
  }

  return (
    <Container>
      <Row>
        <Col sm>
          {status}
          {sessionInfo}
        </Col>
      </Row>
    </Container>
  );
};

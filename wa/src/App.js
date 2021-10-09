import React from 'react';
import { useSelector } from 'react-redux';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom';

import { MyNavbar } from './app/Navbar';
import { PostsList } from './features/posts/PostsList';
import { AddPostForm } from './features/posts/AddPostForm';
import { EditPostForm } from './features/posts/EditPostForm';
import { SinglePostPage } from './features/posts/SinglePostPage';
import { UsersList } from './features/users/UsersList';
import { UserPage } from './features/users/UserPage';
import { NotificationsList } from './features/notifications/NotificationsList';
// import { selectLoginStatus } from './app/appSlice';
import { LoginForm } from './app/LoginForm';


function PrivateRoute({ children, ...rest }) {
  const isLogin = true;//useSelector(selectLoginStatus);
  return (
    <Route
      {...rest}
      render={({ location }) =>
        isLogin ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: "/login",
              state: { from: location }
            }}
          />
        )
      }
    />
  );
}

function App() {
  return (
    <Router>
      {false}
      <MyNavbar />
      <div className="App">
        <Switch>
          <Route exact path="/login" component={LoginForm} />
          <PrivateRoute exact path="/">
            <React.Fragment>
              <AddPostForm />
              <PostsList />
            </React.Fragment>
          </PrivateRoute>
          <PrivateRoute exact path="/posts/:postId" ><SinglePostPage /></PrivateRoute>
          <PrivateRoute exact path="/editPost/:postId"  ><EditPostForm /></PrivateRoute>
          <PrivateRoute exact path="/users" ><UsersList /></PrivateRoute>
          <PrivateRoute exact path="/users/:userId" ><UserPage /></PrivateRoute>
          <PrivateRoute exact path="/notifications"><NotificationsList /></PrivateRoute>
          <Redirect to="/" />
        </Switch>
      </div>
    </Router>
  )
}

export default App;

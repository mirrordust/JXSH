import React from 'react';
import {
  BrowserRouter,
  Link,
  Navigate,
  Outlet,
  Route,
  Routes,
  useLocation
} from 'react-router-dom';
import { Nav, Container, Navbar } from 'react-bootstrap';

import { useAppSelector } from './app/hooks';
import { selectAppisLogin } from './app/appSlice';
import { LoginPanel } from './app/LoginPanel';
import { TagsList } from './features/tags/TagsList';


function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Routes>
          <Route path="/" element={<Layout />} >
            <Route index element={<LoginPanel />} />
            <Route path="about" element={<div>About page</div>} />
            <Route
              path="tags"
              element={
                <RequireAuth>
                  <TagsList />
                </RequireAuth>
              }
            />


            {/* Using path="*"" means "match anything", so this route
                acts like a catch-all for URLs that we don't have explicit
                routes for. */}
            <Route path="*" element={<NoMatch />} />
          </Route>
        </Routes>
      </div>
    </BrowserRouter>
  )
}

function Layout() {
  return (
    <div>
      {/* navigation */}
      <Navbar
        sticky="top"
        collapseOnSelect expand="lg"
        bg="secondary" variant="dark"
      >
        <Container>
          <Navbar.Brand as={Link} to="/">W·Blog Admin</Navbar.Brand>
          <Navbar.Toggle aria-controls="w-navbar-nav" />
          <Navbar.Collapse id="w-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link eventKey="0" as={Link} to="/posts">Posts</Nav.Link>
              <Nav.Link eventKey="1" as={Link} to="/tags">Tags</Nav.Link>
              <Nav.Link eventKey="2" as={Link} to="/collections">Collections</Nav.Link>
              <Nav.Link eventKey="3" as={Link} to="/images">Images</Nav.Link>
              <Nav.Link eventKey="4" as={Link} to="/about">About</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>

      {/* An <Outlet> renders whatever child route is currently active,
      so you can think about this <Outlet> as a placeholder for
      the child routes we defined above. */}
      <Outlet />
    </div>
  );
}

function NoMatch() {
  return (
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/">Go to the home page</Link>
      </p>
    </div>
  );
}

function RequireAuth({ children }) {
  const appIsLogin = useAppSelector(selectAppisLogin);
  const location = useLocation();
  if (!appIsLogin) {
    return <Navigate to="/" state={{ from: location }} />;
  }
  return children;
}

export default App;

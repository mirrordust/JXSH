import { Link } from 'react-router-dom';
import { Nav, Container, Navbar } from 'react-bootstrap';


export const MyNavbar = () => {
  return (
    <Navbar
      sticky="top"
      collapseOnSelect expand="lg"
      bg="secondary" variant="dark"
    >
      <Container>
        <Navbar.Brand>W·Blog Admin</Navbar.Brand>
        <Navbar.Toggle aria-controls="w-navbar-nav" />
        <Navbar.Collapse id="w-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link eventKey="0" as={Link} to="/posts">Posts</Nav.Link>
            <Nav.Link eventKey="1" as={Link} to="/tags">Tags</Nav.Link>
            <Nav.Link eventKey="2" as={Link} to="/collections">Collections</Nav.Link>
            <Nav.Link eventKey="3" as={Link} to="/images">Images</Nav.Link>
            <Nav.Link eventKey="4" as={Link} to="/">PP</Nav.Link>
            <Nav.Link eventKey="5" as={Link} to="/users">UU</Nav.Link>
            <Nav.Link eventKey="6" as={Link} to="/notifications">NN</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

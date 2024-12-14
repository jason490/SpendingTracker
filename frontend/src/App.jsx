import { Container, Nav, Navbar } from 'solid-bootstrap';

function App(props) {
    console.log(props.children)
    return (
        <>
            <Navbar collapseOnSelect expand="lg" >
                <Container>
                    <Navbar.Brand href="/">
                        <img alt="logo" src="/src/assets/svg/logo.svg" width="30" height="30" />
                    </Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse id="responsive-navbar-nav">
                        <Nav class="me-auto" defaultActiveKey="/">
                            <Nav.Link href="/">Dashboard</Nav.Link>
                            <Nav.Link href="/add-remove">Add/Remove</Nav.Link>
                            <Nav.Link href="/export">Export</Nav.Link>
                            <Nav.Link href="/settings">Settings</Nav.Link>
                        </Nav>
                        <Nav>
                            <div class="circle">
                                <img alt="logo" src="/src/assets/svg/user_icon.svg" width="20" height="20" />
                            </div>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            <div>
                {props.children}
            </div>
        </>
    );
}

export default App;

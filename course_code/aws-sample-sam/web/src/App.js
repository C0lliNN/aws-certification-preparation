import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';
import { useEffect, useState } from 'react';
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import Spinner from 'react-bootstrap/Spinner';
import Row from 'react-bootstrap/Row';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import './App.css';

function App() {
  const [images, setImages] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [imageName, setImageName] = useState('');
  const [imageContent, setImageContent] = useState();
  const [show, setShow] = useState(false);

  function handleShowModal() {
    setShow(true);
  }

  function handleCloseModal() {
    setShow(false);
  }

  function handleFetchImages() {
    setIsLoading(true);
    axios
      .get('https://x7wzbbzw30.execute-api.us-east-1.amazonaws.com/Prod/images')
      .then((response) => setImages(response.data))
      .catch((e) => console.log(e))
      .finally(() => setIsLoading(false));
  }

  function handleSaveImage() {
    setIsLoading(true);
    axios
      .post(
        'https://x7wzbbzw30.execute-api.us-east-1.amazonaws.com/Prod/images',
        {
          imageName: imageName,
          imageContent: imageContent
        }
      )
      .then((response) => handleFetchImages())
      .catch((e) => console.log(e))
      .finally(() => handleCloseModal());
  }

  function readFile(e) {
    if (e.target.files.length === 0) {
      return;
    }

    const reader = new FileReader();
    reader.readAsDataURL(e.target.files[0]);
    reader.onload = function () {
      let result = reader.result.split(',')[1];
      setImageContent(result);
    };
    reader.onerror = function (error) {
      console.log('Error: ', error);
    };
  }

  useEffect(() => {
    handleFetchImages();
  }, []);

  return (
    <Container>
      <Row className="align-items-center min-vh-100 my-5">
        <Col>
          <Card>
            <Card.Header>
              <div className="d-flex align-center justify-content-between">
                <h3>Images</h3>
                <Button variant="success" onClick={handleShowModal}>
                  New
                </Button>
              </div>
            </Card.Header>
            <Card.Body>
              <ListGroup variant="flush">
                {images.map((i) => (
                  <ListGroup.Item key={i.ID}>
                    <img
                      src={i.URL}
                      alt={i.Name}
                      style={{
                        maxWidth: '300px',
                        margin: 'auto',
                        display: 'block'
                      }}
                    />
                  </ListGroup.Item>
                ))}
              </ListGroup>
              {isLoading && (
                <div className="my-4 d-flex justify-content-center">
                  <Spinner animation="border" variant="success" />
                </div>
              )}
            </Card.Body>
          </Card>
        </Col>
      </Row>
      <Modal show={show} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>Create Image</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleSaveImage}>
            <Form.Group>
              <Form.Label>Name</Form.Label>
              <Form.Control
                type="text"
                value={imageName}
                onChange={(e) => setImageName(e.target.value)}
              />
            </Form.Group>
            <Form.Group className="my-4">
              <Form.Label>Content</Form.Label>
              <Form.Control
                type="file"
                accept="image/png"
                onClick={(e) => (e.target.value = null)}
                onChange={(e) => readFile(e)}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseModal}>
            Close
          </Button>
          <Button variant="success" onClick={handleSaveImage}>
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </Container>
  );
}

export default App;

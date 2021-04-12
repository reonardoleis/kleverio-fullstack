import logo from "./logo.svg";
import { useState } from "react";
import { validate } from "bitcoin-address-validation";
import "./App.css";
import {
  Container,
  Image,
  Header,
  Input,
  Label,
  Button,
  Divider,
  Card,
  Transition,
  Modal,
  Checkbox,
} from "semantic-ui-react";

import { getBalance } from "./api/api";

import ConfirmedBalance from "./components/ConfirmedBalance";
import UnconfirmedBalance from "./components/UnconfirmedBalance";
function App() {
  const [balance, setBalance] = useState({
    confirmed: 0,
    unconfirmed: 0,
    fetched: false,
  });

  const [loading, setLoading] = useState(false);
  const [errorModal, setErrorModal] = useState(false);
  const [autoUpdateInterval, setAutoUpdateInterval] = useState(undefined);
  const [currentAddress, setCurrentAddress] = useState("");

  const fetchData = async (_address = false) => {
    let address = _address ? _address : document.getElementById("search").value;
    if (!validate(address)) {
      setErrorModal(true);
    } else {
      if (autoUpdateInterval !== undefined) {
        clearInterval(autoUpdateInterval);
        setAutoUpdateInterval(undefined);
      }
      setCurrentAddress(address);
      setLoading(true);
      setBalance({ ...balance, fetched: false });
      let fetchedBalance = await getBalance(address);
      fetchedBalance.fetched = true;
      setBalance(fetchedBalance);
      setLoading(false);
    }
  };

  return (
    <div className="primary-color main-container">
      <Container centered="true">
        <Image src={logo} size="small" centered />
        <Header as="h1" className="text-white" textAlign="center">
          Full-Stack Challenge
        </Header>
      </Container>
      <Container
        centered="true"
        style={{ marginTop: "2rem" }}
        textAlign="center"
      >
        <Label pointing="below" size="large">
          Enter a bitcoin address
        </Label>
        <br></br>
        <Input
          icon="search"
          id="search"
          placeholder="1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX"
          size="huge"
          loading={loading}
        />
        <br></br>
        <Button
          size="huge"
          color="violet"
          style={{ textTransform: "uppercase" }}
          onClick={() => { fetchData() }}
        >
          Get balance
        </Button>
        <br></br>
        <Transition
          duration={{ hide: 150, show: 150 }}
          visible={currentAddress !== ""}
        >
          <Label size="large" style={{ marginTop: "1rem" }}>
            <Checkbox
              checked={autoUpdateInterval !== undefined}
              slider
              label="Auto-update"
              onClick={() => {
                if (autoUpdateInterval !== undefined) {
                  clearInterval(autoUpdateInterval);
                  setAutoUpdateInterval(undefined);
                } else {
                  setAutoUpdateInterval(
                    setInterval(() => {
                      fetchData(currentAddress);
                    }, 1000)
                  );
                }
              }}
            />
          </Label>
        </Transition>
      </Container>
      
      <Transition duration={{ hide: 150, show: 150 }} visible={balance.fetched || autoUpdateInterval !== undefined}>
        <Container>
        <Divider />
          <Card.Group centered>
            <ConfirmedBalance value={balance.confirmed} />
            <UnconfirmedBalance value={balance.unconfirmed} />
          </Card.Group>
        </Container>
      </Transition>
      <Modal
        size="mini"
        open={errorModal}
        onClose={() => {
          setErrorModal(false);
        }}
      >
        <Modal.Content>
          You should specify a valid bitcoin address.
        </Modal.Content>
        <Modal.Actions>
          <Button
            color="violet"
            onClick={() => {
              setErrorModal(false);
            }}
          >
            Ok
          </Button>
        </Modal.Actions>
      </Modal>
    </div>
  );
}

export default App;

import { Header, Icon, Card, CardDescription, Divider } from "semantic-ui-react";

const ConfirmedBalance = (props) => {
  return (
    <Card>
      <Card.Content>
        <Header>
          <Icon name="check circle outline"/>
          <Header.Content>Confirmed</Header.Content>
        </Header>
        <Divider/>
        <CardDescription>
          <Header as="h1" textAlign="center">
            { props.value }
          </Header>
        </CardDescription>
      </Card.Content>
    </Card>
  );
};

export default ConfirmedBalance;

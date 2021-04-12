import { Header, Icon, Card, CardDescription, Divider } from "semantic-ui-react";

const UnconfirmedBalance = (props) => {
  return (
    <Card>
      <Card.Content>
        <Header>
          <Icon name="clock outline"/>
          <Header.Content>Unconfirmed</Header.Content>
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

export default UnconfirmedBalance;

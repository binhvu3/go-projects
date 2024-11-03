import { Button, Container, Stack } from '@chakra-ui/react'
import NavBar from './components/NavBar'
import TodoForm from './components/TodoForm'

function App() {

  return (
    <Stack h="100vh">
      <NavBar />
      <Container>
        <TodoForm />
        {/* <TodoList />  */}
      </Container>
    </Stack>
  )
}

export default App

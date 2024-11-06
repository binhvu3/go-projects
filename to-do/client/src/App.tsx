import { Container, Stack } from '@chakra-ui/react'
import NavBar from './components/NavBar'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'

const PORT = 5000;
export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:" + PORT +"/api" : "/api";
function App() {

  return (
    <Stack h="100vh">
      <NavBar />
      <Container>
        <TodoForm />
        <TodoList /> 
      </Container>
    </Stack>
  )
}

export default App

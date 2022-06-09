import './App.css'
import Candidate from './components/Candidate';

function App() {

  return (
    <div className="App">
      <h1>Election Application</h1>
      <div className='candidateList'>
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
        <Candidate />
      </div>
    </div>
  )
}

export default App
